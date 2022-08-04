import { gql } from "@apollo/client";

export const GET_MODELS = gql`
  query GetModels($projectId: ID!, $first: Int, $last: Int, $after: Cursor, $before: Cursor) {
    projects(projectId: $projectId, first: $first, last: $last, after: $after, before: $before) {
      nodes {
        id
        name
        description
        key
      }
    }
  }
`;

export const CREATE_MODEL = gql`
  mutation CreateModel($projectId: ID!, $name: String, $description: String, $key: String) {
    createModel(
      input: { projectId: $projectId, name: $name, description: $description, key: $key }
    ) {
      model {
        id
        name
        description
        key
      }
    }
  }
`;

export const DELETE_MODEL = gql`
  mutation DeleteModel($modelId: ID!) {
    deleteModel(input: { modelId: $modelId }) {
      modelId
    }
  }
`;

export const UPDATE_MODEL = gql`
  mutation UpdateModel($modelId: ID!, $name: String, $description: String, key: String) {
    updateModel(input: { modelId: $modelId, name: $name, description: $description, key: $key }) {
      project {
        id
        name
        description
        key
      }
    }
  }
`;
