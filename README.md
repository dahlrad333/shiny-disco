# shiny-disco
XX

## Graph Database 
erDiagram
    DOMAIN {
      string name
      date created
    }

    SUBDOMAIN {
      string name
      string alias
    }

    IP {
      string address
      string type
    }

    SERVICE {
      string name
      string port
      string protocol
    }

    USER {
      string username
    }

    EMAIL {
      string address
    }

    VULNERABILITY {
      string id
      string description
      string severity
    }

    DOMAIN ||--o| SUBDOMAIN: "has"
    SUBDOMAIN ||--o| IP: "resolves_to"
    IP ||--o| SERVICE: "hosts"
    SERVICE ||--o| VULNERABILITY: "has_vulnerability"
    IP ||--o| USER: "associated_with"
    USER ||--o| EMAIL: "uses"
    USER ||--o| VULNERABILITY: "exposed_to"
    VULNERABILITY ||--|{ EMAIL: "reported_to"
