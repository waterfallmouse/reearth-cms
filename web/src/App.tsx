import { Provider as Auth0Provider } from "@reearth-cms/auth";
import NotFound from "@reearth-cms/components/atoms/NotFound";
import Account from "@reearth-cms/components/organisms/Settings/Account";
import DashboardPage from "@reearth-cms/components/pages/Dashboard";
import RootPage from "@reearth-cms/components/pages/RootPage";
import WorkspacePage from "@reearth-cms/components/pages/Workspace";
import { Provider as GqlProvider } from "@reearth-cms/gql";
import { Provider as I18nProvider } from "@reearth-cms/i18n";
import { BrowserRouter as Router, useRoutes } from "react-router-dom";

import "./App.css";
import "antd/dist/antd.css";
import WorkspaceListPage from "./components/pages/WorkspaceList";

function AppRoutes() {
  const routes = useRoutes([
    { path: "/", element: <RootPage /> },
    { path: "/dashboard", element: <DashboardPage /> },
    { path: "/dashboard/:workspaceId", element: <DashboardPage /> },
    { path: "/workspace/:workspaceId", element: <WorkspacePage /> },
    { path: "/workspacelist", element: <WorkspaceListPage /> },
    { path: "/account", element: <Account /> },
    { path: "*", element: <NotFound /> },
  ]);
  return routes;
}

function App() {
  return (
    <Auth0Provider>
      <I18nProvider>
        <GqlProvider>
          <Router>
            <AppRoutes />
          </Router>
        </GqlProvider>
      </I18nProvider>
    </Auth0Provider>
  );
}

export default App;
