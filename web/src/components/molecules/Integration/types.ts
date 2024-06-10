export interface IntegrationMember {
  id: string;
  integration?: Integration;
  integrationRole: Role;
  invitedById: string;
  active: boolean;
}

export type Role = "WRITER" | "READER" | "OWNER" | "MAINTAINER";

export interface Integration {
  id: string;
  name: string;
  description?: string | null;
  logoUrl: string;
  developerId: string;
  developer: Developer;
  iType: IntegrationType;
  config: {
    token?: string;
    webhooks?: Webhook[];
  };
}

export interface Developer {
  id: string;
  name: string;
  email: string;
}

export type IntegrationType = "Private" | "Public";

export interface Webhook {
  id: string;
  name: string;
  url: string;
  active: boolean;
  trigger: WebhookTrigger;
}

export interface WebhookTrigger {
  onItemCreate?: boolean | null;
  onItemUpdate?: boolean | null;
  onItemDelete?: boolean | null;
  onAssetUpload?: boolean | null;
  onAssetDelete?: boolean | null;
  onItemPublish?: boolean | null;
  onItemUnPublish?: boolean | null;
}
