# API Authentication Documentation

## Authentication Methods

### Access Token

For API endpoints that require authentication, you must provide both of the following request headers for Access Token authentication:

1. **`Authorization` field in request headers**

    Place the Access Token in the `Authorization` field of the HTTP request headers, formatted as follows:

    ```
    Authorization: <your_access_token>
    ```

    Where `<your_access_token>` should be replaced with the actual Access Token value.

2. **`MedusaXD-Api-User` field in request headers**

    Place the user ID in the `MedusaXD-Api-User` field of the HTTP request headers, formatted as follows:

    ```
    MedusaXD-Api-User: <your_user_id>
    ```

    Where `<your_user_id>` should be replaced with the actual user ID.

**Note:**

*   **Both `Authorization` and `MedusaXD-Api-User` request headers must be provided to pass authentication.**
*   If only one of the request headers is provided, or if neither is provided, a `401 Unauthorized` error will be returned.
*   If the Access Token in `Authorization` is invalid, a `401 Unauthorized` error will be returned with the message "No permission for this operation: invalid access token".
*   If the user ID in `MedusaXD-Api-User` does not match the Access Token, a `401 Unauthorized` error will be returned with the message "No permission for this operation: does not match logged-in user, please log in again".
*   If the `MedusaXD-Api-User` request header is not provided, a `401 Unauthorized` error will be returned with the message "No permission for this operation: MedusaXD-Api-User not provided".
*   If the `MedusaXD-Api-User` request header format is incorrect, a `401 Unauthorized` error will be returned with the message "No permission for this operation: MedusaXD-Api-User format error".
*   If the user has been disabled, a `403 Forbidden` error will be returned with the message "User has been banned".
*   If the user has insufficient permissions, a `403 Forbidden` error will be returned with the message "No permission for this operation: insufficient permissions".
*   If the user information is invalid, a `403 Forbidden` error will be returned with the message "No permission for this operation: invalid user information".

## Curl Example

Assuming your Access Token is `access_token`, user ID is `123`, and the API endpoint you want to access is `/api/user/self`, you can use the following curl command:

```bash
curl -X GET \
  -H "Authorization: access_token" \
  -H "MedusaXD-Api-User: 123" \
  https://your-domain.com/api/user/self
```

Please replace `access_token`, `123`, and `https://your-domain.com` with actual values.
