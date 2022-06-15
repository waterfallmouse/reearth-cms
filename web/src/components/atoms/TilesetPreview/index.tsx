import { Viewer, Cesium3DTileset } from "cesium";
import {
  Viewer as ResiumViewer,
  Cesium3DTileset as Resium3DTileset,
} from "resium";
import { Cesium3DTilesetProps } from "resium/dist/Cesium3DTileset/Cesium3DTileset";
import { ViewerProps } from "resium/dist/Viewer/Viewer";

type TilesetPreviewProps = {
  viewerProps?: ViewerProps;
  tilesetProps: Cesium3DTilesetProps;
};

const TilesetPreview: React.FC<TilesetPreviewProps> = ({
  viewerProps,
  tilesetProps,
}) => {
  let viewer: Viewer | undefined;

  const handleReady = async (tileset: Cesium3DTileset) => {
    try {
      await viewer?.zoomTo(tileset.root.tileset);
    } catch (error) {
      console.log(error);
    }
  };

  return (
    <ResiumViewer
      {...viewerProps}
      ref={(e) => {
        viewer = e?.cesiumElement;
      }}
    >
      <Resium3DTileset {...tilesetProps} onReady={handleReady} />
    </ResiumViewer>
  );
};

export default TilesetPreview;
