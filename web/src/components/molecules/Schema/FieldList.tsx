import styled from "@emotion/styled";
import React from "react";

import Icon from "@reearth-cms/components/atoms/Icon";
import List from "@reearth-cms/components/atoms/List";
import { fieldTypes } from "@reearth-cms/components/organisms/Project/Schema/fieldTypes";

import { FieldType } from "../Dashboard/types";

export interface Props {
  className?: string;
  addField: (fieldType: FieldType) => void;
}

const data = [
  {
    title: "Text",
    color: "#FF7875",
    fields: ["Text", "TextArea", "MarkdownText"],
  },
  {
    title: "Asset",
    color: "#FF9C6E",
    fields: ["Asset"],
  },
  {
    title: "Select",
    color: "#7CB305",
    fields: ["Select"],
  },
  {
    title: "Number",
    color: "#36CFC9",
    fields: ["Integer"],
  },
  {
    title: "URL",
    color: "#9254DE",
    fields: ["URL"],
  },
];

const FieldList: React.FC<Props> = ({ addField }) => {
  return (
    <>
      <h1>Add Field</h1>
      <FieldStyledList
        itemLayout="horizontal"
        dataSource={data}
        renderItem={item => (
          <>
            <FieldCategoryTitle>{item.title}</FieldCategoryTitle>
            {item.fields?.map((field: string) => (
              <List.Item key={field} onClick={() => addField(field as FieldType)}>
                <List.Item.Meta
                  avatar={<Icon icon={fieldTypes[field].icon} color={item.color} />}
                  title={fieldTypes[field].title}
                  description={fieldTypes[field].description}
                />
              </List.Item>
            ))}
          </>
        )}
      />
    </>
  );
};

const FieldCategoryTitle = styled.h2`
  font-weight: 400;
  font-size: 12px;
  line-height: 20px;
  margin-bottom: 12px;
  margin-top: 12px;
  color: rgba(0, 0, 0, 0.45);
`;

const FieldStyledList = styled(List)`
  .ant-list-item {
    background-color: #fff;
    cursor: pointer;
    + .ant-list-item {
      margin-top: 12px;
    }
    padding: 4px;
    box-shadow: 0px 2px 8px rgba(0, 0, 0, 0.15);
    .ant-list-item-meta {
      .ant-list-item-meta-title {
        margin: 0;
      }
      align-items: center;
      .ant-list-item-meta-avatar {
        border: 1px solid #f0f0f0;
        width: 28px;
        height: 28px;
        display: flex;
        justify-content: center;
        align-items: center;
      }
    }
  }
`;

export default FieldList;
