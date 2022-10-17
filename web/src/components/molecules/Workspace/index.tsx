import styled from "@emotion/styled";

import ProjectCreationModal, {
  FormValues,
} from "@reearth-cms/components/molecules/Common/ProjectCreationModal";
import WorkspaceCreationModal from "@reearth-cms/components/molecules/Common/WorkspaceCreationModal";
import ProjectList from "@reearth-cms/components/molecules/ProjectList/ProjectList";
import Greeting from "@reearth-cms/components/molecules/Workspace/Greeting";
import { Project as ProjectType } from "@reearth-cms/components/molecules/Workspace/types";
import WorkspaceHeader from "@reearth-cms/components/molecules/Workspace/WorkspaceHeader";

export type Props = {
  projects?: ProjectType[];
  projectModal?: boolean;
  workspaceModal?: boolean;
  onProjectSearch: (value: string) => void;
  onProjectModalOpen: () => void;
  onProjectNavigation: (project: ProjectType, tab?: string) => void;
  onWorkspaceModalClose: () => void;
  onWorkspaceModalOpen: () => void;
  onWorkspaceCreate: (data: { name: string }) => Promise<void>;
  onClose?: (refetch?: boolean) => void;
  onSubmit?: (values: FormValues) => Promise<void> | void;
};

const WorkspaceWrapper: React.FC<Props> = ({
  projects,
  projectModal,
  workspaceModal,
  onProjectSearch,
  onProjectModalOpen,
  onProjectNavigation,
  onWorkspaceModalClose,
  onWorkspaceModalOpen,
  onWorkspaceCreate,
  onClose,
  onSubmit,
}) => {
  return (
    <Wrapper>
      <Greeting />
      <Content>
        <WorkspaceHeader
          onProjectSearch={onProjectSearch}
          onProjectModalOpen={onProjectModalOpen}
          onWorkspaceModalOpen={onWorkspaceModalOpen}
        />
        <ProjectList
          projects={projects}
          onProjectModalOpen={onProjectModalOpen}
          onProjectNavigation={onProjectNavigation}
        />
        <ProjectCreationModal open={projectModal} onClose={onClose} onSubmit={onSubmit} />
      </Content>
      <WorkspaceCreationModal
        open={workspaceModal}
        onClose={onWorkspaceModalClose}
        onSubmit={onWorkspaceCreate}
      />
    </Wrapper>
  );
};

const Wrapper = styled.div`
  background-color: #fff;
  margin: 16px;
  min-height: 100%;
`;

const Content = styled.div`
  padding: 16px;
`;

export default WorkspaceWrapper;