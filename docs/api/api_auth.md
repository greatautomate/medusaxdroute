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

**注意：**

*   **必须同时提供 `Authorization` 和 `MedusaXD-API-User` 两个请求头才能通过鉴权。**
*   如果只提供其中一个请求头，或者两个请求头都未提供，则会返回 `401 Unauthorized` 错误。
*   如果 `Authorization` 中的 Access Token 无效，则会返回 `401 Unauthorized` 错误，并提示“无权进行此操作，access token 无效”。
*   如果 `MedusaXD-API-User` 中的用户 ID 与 Access Token 不匹配，则会返回 `401 Unauthorized` 错误，并提示“无权进行此操作，与登录用户不匹配，请重新登录”。
*   如果没有提供 `New-Api-User` 请求头，则会返回 `401 Unauthorized` 错误，并提示“无权进行此操作，未提供 New-Api-User”。
*   如果 `New-Api-User` 请求头格式错误，则会返回 `401 Unauthorized` 错误，并提示“无权进行此操作，New-Api-User 格式错误”。
*   如果用户已被禁用，则会返回 `403 Forbidden` 错误，并提示“用户已被封禁”。
*   如果用户权限不足，则会返回 `403 Forbidden` 错误，并提示“无权进行此操作，权限不足”。
*   如果用户信息无效，则会返回 `403 Forbidden` 错误，并提示“无权进行此操作，用户信息无效”。

## Curl 示例

假设您的 Access Token 为 `access_token`，用户 ID 为 `123`，要访问的 API 接口为 `/api/user/self`，则可以使用以下 curl 命令：

```bash
curl -X GET \
  -H "Authorization: access_token" \
  -H "MedusaXD-API-User: 123" \
  https://your-domain.com/api/user/self
```

请将 `access_token`、`123` 和 `https://your-domain.com` 替换为实际的值。

