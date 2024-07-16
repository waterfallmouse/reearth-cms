import styled from "@emotion/styled";
import MonacoEditor, { OnMount, BeforeMount } from "@monaco-editor/react";
import { Ion, Cartesian3, Viewer as CesiumViewer, Cartographic, Math, SceneMode } from "cesium";
import { editor } from "monaco-editor";
import { useCallback, useEffect, useMemo, useState, useRef } from "react";
import {
  CesiumComponentRef,
  CesiumMovementEvent,
  Viewer,
  ScreenSpaceCameraController,
  RootEventTarget,
} from "resium";

import Button from "@reearth-cms/components/atoms/Button";
import Icon from "@reearth-cms/components/atoms/Icon";
import Marker from "@reearth-cms/components/atoms/Icon/Icons/mapPinFilled.svg";
import { Field } from "@reearth-cms/components/molecules/Schema/types";
import { config } from "@reearth-cms/config";

type GeoType = "point" | "lineString" | "polygon";

interface Props {
  field: Field;
}

const GeometryItem: React.FC<Props> = ({ field }) => {
  const editorSupportedTypes = useMemo(
    () => field.typeProperty?.editorSupportedTypes?.[0],
    [field.typeProperty?.editorSupportedTypes],
  );
  const editorRef = useRef<editor.IStandaloneCodeEditor>();

  const copyButtonClick = useCallback(() => {
    const value = editorRef.current?.getValue();
    if (value) navigator.clipboard.writeText(value);
  }, []);
  const editorDeleteButtonClick = useCallback(() => {
    editorRef.current?.setValue("");
  }, []);

  const options = useMemo(
    () => ({
      bracketPairColorization: {
        enabled: true,
      },
      minimap: {
        enabled: false,
      },
      readOnly: false,
      wordWrap: "on" as const,
      scrollBeyondLastLine: false,
    }),
    [],
  );

  const [isEmpty, setIsEmpty] = useState(false);

  const handleEditorOnChange = (value?: string) => {
    value ? setIsEmpty(false) : setIsEmpty(true);
  };

  const handleEditorDidMount: OnMount = (editor, monaco) => {
    editorRef.current = editor;
  };
  const handleEditorWillMount: BeforeMount = monaco => {
    monaco.languages.json.jsonDefaults.setDiagnosticsOptions({
      schemaValidation: "error",
      schemas: [
        {
          uri: "",
          fileMatch: ["*"],
          schema: {
            type: "object",
            properties: {
              type: {
                enum: [
                  "Point",
                  "LineString",
                  "Polygon",
                  "MultiPoint",
                  "MultiLineString",
                  "MultiPolygon",
                  "GeometryCollection",
                ],
              },
              coodinates: {},
            },
          },
        },
      ],
    });
  };

  const [isReady, setIsReady] = useState(false);
  useEffect(() => {
    Ion.defaultAccessToken = config()?.cesiumIonAccessToken ?? Ion.defaultAccessToken;
    setIsReady(true);
  }, []);

  const [geoType, setGeoType] = useState<GeoType>();
  const [isDrawing, setIsDrawing] = useState(false);

  const geoTypeSet = useCallback((type: GeoType) => {
    setGeoType(type);
    setIsDrawing(prev => !prev);
  }, []);

  const pinButtonClick = useCallback(() => {
    geoTypeSet("point");
  }, [geoTypeSet]);

  const lineStringButtonClick = useCallback(() => {
    geoTypeSet("lineString");
  }, [geoTypeSet]);

  const polygonButtonClick = useCallback(() => {
    geoTypeSet("polygon");
  }, [geoTypeSet]);

  const [geoValues, setGeoValues] = useState<Map<string, number[]>>(new Map());
  const [enableTranslate, setEnableTranslate] = useState(true);

  const viewer = useRef<CesiumComponentRef<CesiumViewer>>(null);

  const handleZoom = useCallback((isZoomIn: boolean) => {
    if (viewer.current?.cesiumElement) {
      const ellipsoid = viewer.current.cesiumElement.scene.globe.ellipsoid;
      const camera = viewer.current.cesiumElement.camera;
      const cameraHeight = ellipsoid.cartesianToCartographic(camera.position).height;
      const moveRate = cameraHeight / 10;
      if (isZoomIn) {
        camera.moveForward(moveRate);
      } else {
        camera.moveBackward(moveRate);
      }
    }
  }, []);

  const positionsRef = useRef<number[][]>([]);

  const handleClick = useCallback(
    (_movement: CesiumMovementEvent) => {
      if (!isDrawing) return;
      if (_movement.position && viewer.current?.cesiumElement) {
        const ellipsoid = viewer.current.cesiumElement.scene.globe.ellipsoid;
        const cartesian = viewer.current.cesiumElement.camera.pickEllipsoid(
          _movement.position,
          ellipsoid,
        );
        if (cartesian) {
          const cartographic = Cartographic.fromCartesian(cartesian);
          const lon = Math.toDegrees(cartographic.longitude);
          const lat = Math.toDegrees(cartographic.latitude);
          if (geoType === "point") {
            const entity = viewer.current.cesiumElement.entities.add({
              position: cartesian,
              billboard: {
                image: Marker,
                width: 30,
                height: 30,
              },
            });
            setIsDrawing(false);
            setGeoValues(prev => {
              prev.set(entity.id, [lon, lat]);
              return new Map(prev);
            });
            editorRef.current?.setValue(
              JSON.stringify(
                {
                  type: "Point",
                  coodinates: [lon, lat],
                },
                null,
                2,
              ),
            );
          } else {
            positionsRef.current?.push([lon, lat]);
            if (geoType === "lineString") {
              viewer.current.cesiumElement.entities.add({
                position: cartesian,
                polyline: {
                  positions: Cartesian3.fromDegreesArray(positionsRef.current.flat()),
                },
              });
              editorRef.current?.setValue(
                JSON.stringify(
                  {
                    type: "LineString",
                    coodinates: positionsRef.current,
                  },
                  null,
                  2,
                ),
              );
            } else {
              viewer.current.cesiumElement.entities.add({
                position: cartesian,
                polygon: {
                  hierarchy: Cartesian3.fromDegreesArray(positionsRef.current.flat()),
                  extrudedHeight: 50000,
                },
              });
              editorRef.current?.setValue(
                JSON.stringify(
                  {
                    type: "Polygon",
                    coodinates: [positionsRef.current],
                  },
                  null,
                  2,
                ),
              );
            }
          }
        }
      }
    },
    [geoType, isDrawing],
  );

  const timeout = useRef<NodeJS.Timeout>();
  const singleClick = useCallback(
    (movement: CesiumMovementEvent) => {
      if (timeout.current) {
        clearTimeout(timeout.current);
        timeout.current = undefined;
      } else {
        timeout.current = setTimeout(() => {
          handleClick(movement);
          timeout.current = undefined;
        }, 250);
      }
    },
    [handleClick],
  );

  const doubleClick = useCallback(() => {
    setIsDrawing(false);
  }, []);

  const [isGrabbing, setIsGrabbing] = useState(false);
  const [entityId, setEntityId] = useState("");

  const onMouseDown = useCallback(() => {
    setIsGrabbing(true);
  }, []);

  const onMouseMove = useCallback(
    (movement: CesiumMovementEvent, _: RootEventTarget) => {
      if (entityId && viewer.current?.cesiumElement && movement.endPosition) {
        const cartesian = viewer.current.cesiumElement.camera.pickEllipsoid(
          movement.endPosition,
          viewer.current.cesiumElement.scene.globe.ellipsoid,
        );
        const point = viewer.current.cesiumElement.entities.getById(entityId);
        if (point && cartesian) {
          point.position = cartesian as any;
          const cartographic = Cartographic.fromCartesian(cartesian);
          const lon = Math.toDegrees(cartographic.longitude);
          const lat = Math.toDegrees(cartographic.latitude);
          setGeoValues(prev => {
            prev.set(entityId, [lon, lat]);
            return new Map(prev);
          });
        }
      }
    },
    [entityId],
  );

  const onMouseUp = useCallback(() => {
    setEntityId("");
    setIsGrabbing(false);
  }, []);

  useEffect(() => {
    const handleEnter = (event: KeyboardEvent) => {
      if (event.key === "Enter") {
        setIsDrawing(false);
      }
    };
    document.addEventListener("keydown", handleEnter);
    return () => {
      document.removeEventListener("keydown", handleEnter);
    };
  }, []);

  return (
    <GeometryField>
      <EditorWrapper>
        <EditorButtons>
          <EditorButton
            icon={<Icon icon="editorCopy" size={12} />}
            size="small"
            onClick={copyButtonClick}
          />
          <EditorButton
            icon={<Icon icon="trash" size={12} />}
            size="small"
            onClick={editorDeleteButtonClick}
          />
        </EditorButtons>
        <MonacoEditor
          height="100%"
          language={"json"}
          options={options}
          value={JSON.stringify(
            {
              type: "Point",
              coodinates: [],
            },
            null,
            2,
          )}
          onChange={handleEditorOnChange}
          onMount={handleEditorDidMount}
          beforeMount={handleEditorWillMount}
        />
        <Placeholder isEmpty={isEmpty}>
          {JSON.stringify(
            {
              type: "Point",
              coodinates: [],
            },
            null,
            2,
          )}
        </Placeholder>
      </EditorWrapper>
      {isReady && (
        <ViewerWrapper>
          <ViewerButtons>
            {field.type === "GeometryEditor" && (
              <GeoButtons>
                {(editorSupportedTypes === "POINT" || editorSupportedTypes === "ANY") && (
                  <GeoButton
                    icon={<Icon icon="mapPin" size={22} />}
                    onClick={pinButtonClick}
                    selected={isDrawing && geoType === "point"}
                  />
                )}
                {(editorSupportedTypes === "LINESTRING" || editorSupportedTypes === "ANY") && (
                  <GeoButton
                    icon={<Icon icon="lineString" size={22} />}
                    onClick={lineStringButtonClick}
                    selected={isDrawing && geoType === "lineString"}
                  />
                )}
                {(editorSupportedTypes === "POLYGON" || editorSupportedTypes === "ANY") && (
                  <GeoButton
                    icon={<Icon icon="polygon" size={22} />}
                    onClick={polygonButtonClick}
                    selected={isDrawing && geoType === "polygon"}
                  />
                )}
              </GeoButtons>
            )}
            <ZoomButtons>
              <Button
                icon={<Icon icon="plus" />}
                onClick={() => {
                  handleZoom(true);
                }}
              />
              <Button
                icon={<Icon icon="minus" />}
                onClick={() => {
                  handleZoom(false);
                }}
              />
            </ZoomButtons>
          </ViewerButtons>
          <StyledViewer
            infoBox={false}
            navigationHelpButton={false}
            homeButton={false}
            projectionPicker={false}
            sceneModePicker={false}
            sceneMode={SceneMode.SCENE2D}
            baseLayerPicker={false}
            fullscreenButton={false}
            vrButton={false}
            selectionIndicator={false}
            timeline={false}
            animation={false}
            geocoder={false}
            onClick={singleClick}
            onDoubleClick={doubleClick}
            onMouseDown={onMouseDown}
            onMouseMove={onMouseMove}
            onMouseUp={onMouseUp}
            ref={viewer}
            isGrabbing={isGrabbing}
            isDrawing={isDrawing}>
            <ScreenSpaceCameraController enableTranslate={enableTranslate} />
          </StyledViewer>
        </ViewerWrapper>
      )}
    </GeometryField>
  );
};

export default GeometryItem;

const GeometryField = styled.div`
  display: flex;
  max-height: 432px;
  aspect-ratio: 1.86 / 1;
  box-shadow: 0px 2px 8px 0px #00000026;
`;

const EditorWrapper = styled.div`
  width: 40%;
  position: relative;
`;

const EditorButtons = styled.div`
  position: absolute;
  right: 8px;
  color: #8c8c8c;
  z-index: 1;
  display: flex;
  flex-direction: column;
  gap: 4px;
  padding-top: 12px;
`;

const EditorButton = styled(Button)`
  color: #8c8c8c;
`;

const Placeholder = styled.div<{ isEmpty: boolean }>`
  display: ${({ isEmpty }) => (isEmpty ? "block" : "none")};
  position: absolute;
  white-space: pre-wrap;
  top: 0px;
  left: 65px;
  font-size: 14px;
  color: #bfbfbf;
  font-family: Consolas, "Courier New", monospace;
  pointer-events: none;
  user-select: none;
  line-height: 1.4;
`;

const ViewerWrapper = styled.div`
  position: relative;
  flex: 1;
`;

const StyledViewer = styled(Viewer)<{ isDrawing: boolean; isGrabbing: boolean }>`
  height: 100%;
  cursor: ${({ isDrawing, isGrabbing }) =>
    isDrawing ? "crosshair" : isGrabbing ? "grabbing" : "grab"};
`;

const ViewerButtons = styled.div`
  z-index: 1;
  position: absolute;
  right: 8px;
  height: 100%;
  display: flex;
  flex-direction: column;
  justify-content: space-between;
`;

const GeoButtons = styled.div`
  padding: 15px 0;
`;

const GeoButton = styled(Button)<{ selected: boolean }>`
  color: ${({ selected }) => (selected ? "#1677ff" : "#434343")};
`;

const ZoomButtons = styled.div`
  display: flex;
  flex-direction: column;
  padding: 10px 0;
  button:first-child {
    border-end-start-radius: 0;
    border-end-end-radius: 0;
  }
  button:last-child {
    border-start-start-radius: 0;
    border-start-end-radius: 0;
  }
`;
