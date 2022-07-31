# Secrets Manager

## Data Diagram
```mermaid
erDiagram
  User {
    string uuid
    string openid_id
    string openid_source
  }
  Secrets {
    string encrypted
    Time rotationTime
    blob rotationScript
  }
```