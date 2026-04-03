```mermaid
graph TD
    %% Define the Cloud/Internet
    subgraph CLOUD ["Public Cloud / Internet"]
        direction TB
        internet_icon( )
    end

    %% --- PUBLIC SUBNET (192.168.0.0/24) ---
    subgraph PUBLIC_NET ["PUBLIC SUBNET (192.168.0.0/24)"]
        direction TB
        
        NGINX_RP[Nginx Reverse Proxy]
        FRONTEND_PUB[Public Frontend]
        BACKEND_PUB[Public Backend]
        
        %% Connect items within the subnet
        NGINX_RP --- FRONTEND_PUB
        NGINX_RP --- BACKEND_PUB
    end

    %% --- PRIVATE SUBNET (192.168.1.0/24) ---
    subgraph PRIVATE_NET ["PRIVATE SUBNET (192.168.1.0/24)"]
        direction TB
        
        NGINX_INT[Nginx Internal]
        BACKEND_PRI[Private Backend]
        CMS_PORTAL[Private CMS Portal]
        PREVIEW_PORTAL[Private Preview Portal]
        POSTGRES_DB[(Postgres DB)]
        MINIO_S3[MinIO S3 Storage]
        
        %% Connect items within the subnet
        NGINX_INT --- BACKEND_PRI
        NGINX_INT --- CMS_PORTAL
        NGINX_INT --- PREVIEW_PORTAL
        NGINX_INT --- POSTGRES_DB
        NGINX_INT --- MINIO_S3
    end

    %% --- VPN SUBNET (192.168.2.0/24) ---
    subgraph VPN_NET ["VPN SUBNET (192.168.2.0/24)"]
        direction LR
        
        VPN_GATEWAY[VPN Gateway Server]
        NGINX_VPN_IF[Nginx VPN Interface]
        
        %% Connect items within the subnet
        NGINX_VPN_IF --- VPN_GATEWAY
    end

    %% ==========================================
    %% Inter-Subnet and Cloud Connections (The Arrows)
    %% ==========================================

    %% Cloud connection to Public Nginx (RP)
    internet_icon == "External Traffic" ==> NGINX_RP:::public_if

    %% Cloud connection from VPN Gateway
    VPN_GATEWAY:::public_if == "Secure Tunnel" ==> internet_icon

    %% Connection between Public and Private subnets (Bridge)
    BACKEND_PUB:::public_net_if === BACKEND_PUB:::private_net_if

    %% Connect Private Subnet Nginx to other private servers
    NGINX_INT --- POSTGRES_DB
    NGINX_INT --- MINIO_S3
    NGINX_INT --- BACKEND_PRI
    NGINX_INT --- CMS_PORTAL
    NGINX_INT --- PREVIEW_PORTAL

    %% Connection of Nginx Reverse Proxy to VPN Subnet
    NGINX_RP:::public_net_if === NGINX_VPN_IF:::vpn_net_if

    %% ==========================================
    %% Labels and IP Addresses
    %% ==========================================

    %% Public Subnet IPs
    NGINX_RP:::public_if:::label: "1.2.3.4 (Pub IF)<br>192.168.0.2 (Pri IF)"
    FRONTEND_PUB:::label: "192.168.0.4"
    BACKEND_PUB:::public_net_if:::label: "192.168.0.3"

    %% Private Subnet IPs
    BACKEND_PUB:::private_net_if:::label: "192.168.1.2"
    POSTGRES_DB:::label: "192.168.1.3"
    MINIO_S3:::label: "192.168.1.4 (S3)"
    BACKEND_PRI:::label: "192.168.1.4 (BE)"
    CMS_PORTAL:::label: "192.168.1.5"
    PREVIEW_PORTAL:::label: "192.168.1.6"
    NGINX_INT:::label: "192.168.1.7"

    %% VPN Subnet IPs
    NGINX_VPN_IF:::vpn_net_if:::label: "192.168.2.2"
    VPN_GATEWAY:::public_if:::label: "1.2.3.5 (Pub IF)<br>192.168.2.3 (Pri IF)"

    %% ==========================================
    %% Styling
    %% ==========================================
    %% Subnet Box Styling
    style PUBLIC_NET fill:#e6f7ff,stroke:#1890ff,stroke-width:2px,stroke-dasharray: 5 5;
    style PRIVATE_NET fill:#fff1f0,stroke:#f5222d,stroke-width:2px,stroke-dasharray: 5 5;
    style VPN_NET fill:#f6ffed,stroke:#52c41a,stroke-width:2px,stroke-dasharray: 5 5;
    style CLOUD fill:#fafafa,stroke:#d9d9d9,stroke-width:2px;

    %% Icon and Component Styling
    classDef public_if fill:#1890ff,stroke:#fff,color:#fff,stroke-width:2px;
    classDef private_net_if fill:#f5222d,stroke:#fff,color:#fff,stroke-width:2px;
    classDef public_net_if fill:#1890ff,stroke:#fff,color:#fff,stroke-width:2px;
    classDef vpn_net_if fill:#52c41a,stroke:#fff,color:#fff,stroke-width:2px;
    
    classDef label fill:none,stroke:none,color:#333,font-size:12px;

    %% Assigning styles to nodes
    %% NGINX_RP has a public IF (1.2.3.4) and a private IF (192.168.0.2)
    %% In this visualization, NGINX_RP is the main box.
```
