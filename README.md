# Secrets Manager

## Purpose

A open source secrets management platform that you can deploy/use our cloud hosted solution to manage secrets for your applications. You are able to upload scripts that will run to rotate these secrets.

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