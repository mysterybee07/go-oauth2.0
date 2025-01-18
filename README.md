# OAuth 2.0

## What is OAuth?

OAuth (Open Authorization) is an open standard for token-based authorization, enabling third-party applications to access user resources securely without exposing user credentials. It provides a mechanism to grant limited access to resources while maintaining security and user control.

---

## Principles of OAuth 2.0

### 1. Authorization, Not Authentication
OAuth 2.0 is an **authorization protocol**, not an **authentication protocol**.  
Its primary goal is to allow applications to gain access to a set of resources, such as APIs or user data, with the user's consent, without handling their credentials directly.

### 2. Access Tokens
OAuth 2.0 uses **Access Tokens** to grant permission to resources.

- **Access Token** is a data object that represents the user's authorization.
- The protocol does not enforce a specific token format, but **JSON Web Tokens (JWT)** are commonly used.
- Tokens may include additional information, such as:
  - User ID
  - Scopes
  - Expiration times

### 3. Token Expiry
- Access Tokens typically have an **expiration date** for security purposes.
- Applications need to refresh expired tokens using a **Refresh Token**, if available, to ensure secure, limited-time access.

---

## OAuth 2.0 Roles

### 1. Resource Owner
The Resource Owner is the entity that owns the protected resources.  
**Example**: A user who has photos stored in a cloud service (e.g., Google Photos).

### 2. Client
The Client is the application that requests access to the Resource Owner's resources.  
**Example**: A third-party photo printing app that needs access to the user’s Google Photos to create a photo book.

### 3. Authorization Server
The Authorization Server handles requests for Access Tokens and issues them upon successful authentication and consent from the Resource Owner.  
- It provides the **Authorization Endpoint** for user interaction and **Token Endpoint** for machine-to-machine token exchanges.  
**Example**: Google’s OAuth 2.0 server.

### 4. Resource Server
The Resource Server hosts and protects the Resource Owner's data. It validates Access Tokens received from the Client and provides the requested resources.  
**Example**: Google Photos API, which returns the user’s photos when provided with a valid Access Token.

---

## OAuth 2.0 Scopes
Scopes define the level of access a Client has to resources.  
**Example**: A scope for Google Photos could limit access to "read-only" or grant permissions for "upload and delete."

---

## OAuth 2.0 Access Tokens and Authorization Code

- Access Tokens may be issued after an Authorization Code is exchanged to enhance security.
- A Refresh Token may also be issued for long-term access without re-prompting the user for consent.

---

## Grant Types in OAuth 2.0

1. **Authorization Code Grant**: Used by web applications for secure, server-side token exchange.
2. **Implicit Grant** (deprecated): Simplified flow for SPAs but less secure.
3. **Authorization Code with PKCE**: Secures token exchanges for SPAs and mobile apps.
4. **Resource Owner Credentials Grant**: Direct access using the Resource Owner's credentials (for trusted clients only).
5. **Client Credentials Grant**: Used by backend systems and non-user-facing applications.
6. **Device Authorization Flow**: Ideal for apps on input-limited devices like smart TVs.
7. **Refresh Token Grant**: Used to obtain a new Access Token when the current one expires.

---

## How OAuth 2.0 Works: A Simplified Flow

1. **Client Registration**: The Client registers with the Authorization Server to get a client ID and secret.
2. **Authorization Request**: The Client requests user consent and provides scopes, redirect URIs, etc.
3. **User Interaction**: The Resource Owner authenticates and consents to grant access.
4. **Token Exchange**: The Client exchanges an Authorization Code for an Access Token.
5. **Access Resource**: The Client uses the Access Token to request data from the Resource Server.

---

