import { useCallback, useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";

import Notification from "@reearth-cms/components/atoms/Notification";
import { User } from "@reearth-cms/components/molecules/Workspace/types";
import { useCreateWorkspaceMutation, useGetMeQuery } from "@reearth-cms/gql/graphql-client-api";
import { useT } from "@reearth-cms/i18n";
import { useWorkspace, useProject } from "@reearth-cms/state";

export default ({ projectId, workspaceId }: { projectId?: string; workspaceId?: string }) => {
  const [currentWorkspace, setCurrentWorkspace] = useWorkspace();
  const [currentProject, setCurrentProject] = useProject();
  const [workspaceModalShown, setWorkspaceModalShown] = useState(false);
  const { data, refetch } = useGetMeQuery();
  const t = useT();

  const navigate = useNavigate();

  const user: User = {
    name: data?.me?.name || "",
  };

  const workspaces = data?.me?.workspaces;
  const workspace = workspaces?.find(workspace => workspace.id === workspaceId);
  const personalWorkspace = workspaces?.find(
    workspace => workspace.id === data?.me?.myWorkspace.id,
  );
  const personal = workspaceId === data?.me?.myWorkspace.id;

  useEffect(() => {
    if (currentWorkspace || workspaceId || !data) return;
    setCurrentWorkspace(data.me?.myWorkspace);
    navigate(`/workspace/${data.me?.myWorkspace?.id}`);
  }, [data, navigate, setCurrentWorkspace, currentWorkspace, workspaceId]);

  useEffect(() => {
    if (workspace?.id && workspace.id !== currentWorkspace?.id) {
      setCurrentWorkspace({
        personal,
        ...workspace,
      });
    }
  }, [currentWorkspace, workspace, setCurrentWorkspace, personal]);

  useEffect(() => {
    if (projectId && projectId !== currentProject?.id) {
      // TO DO: UPDATE HERE OR REMOVE IT. LOGIC ISNT GOOD...
      setCurrentProject({ id: projectId });
    }
  }, [projectId, currentProject?.id, setCurrentProject]);

  const handleWorkspaceChange = useCallback(
    (workspaceId: string) => {
      const workspace = workspaces?.find(workspace => workspace.id === workspaceId);
      if (workspace) {
        setCurrentWorkspace(workspace);
        navigate(`/workspace/${workspaceId}`);
      }
    },
    [workspaces, setCurrentWorkspace, navigate],
  );

  const [createWorkspaceMutation] = useCreateWorkspaceMutation();
  const handleWorkspaceCreate = useCallback(
    async (data: { name: string }) => {
      const results = await createWorkspaceMutation({
        variables: { name: data.name },
        refetchQueries: ["GetWorkspaces"],
      });
      if (results.data?.createWorkspace) {
        Notification.success({ message: t("Successfully created workspace!") });
        setCurrentWorkspace(results.data.createWorkspace.workspace);
        navigate(`/workspace/${results.data.createWorkspace.workspace.id}`);
      }
      refetch();
    },
    [createWorkspaceMutation, setCurrentWorkspace, refetch, navigate, t],
  );

  const handleWorkspaceModalClose = useCallback(() => {
    setWorkspaceModalShown(false);
  }, []);

  const handleWorkspaceModalOpen = useCallback(() => setWorkspaceModalShown(true), []);

  const handleNavigateToSettings = useCallback(() => {
    navigate(`/workspace/${personalWorkspace?.id}/account`);
  }, [personalWorkspace?.id, navigate]);

  return {
    user,
    personalWorkspace,
    workspaces,
    currentWorkspace,
    workspaceModalShown,
    currentProject,
    handleWorkspaceModalClose,
    handleWorkspaceModalOpen,
    handleWorkspaceCreate,
    handleWorkspaceChange,
    handleNavigateToSettings,
  };
};