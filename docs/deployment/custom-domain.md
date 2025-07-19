# Custom Domain Configuration for MedusaXD API

This guide explains how to configure a custom domain for your MedusaXD API deployment on Render.com.

## 🌐 Overview

By default, your MedusaXD API will be accessible at `https://your-app-name.onrender.com`. However, you can configure a custom domain like `https://api.yourdomain.com` for a more professional setup.

## 📋 Prerequisites

- ✅ MedusaXD API deployed on Render.com
- ✅ Custom domain name (e.g., `yourdomain.com`)
- ✅ Access to your domain's DNS settings

## 🚀 Step-by-Step Setup

### Step 1: Add Custom Domain in Render

1. **Go to your Render service dashboard**
2. **Navigate to "Settings" tab**
3. **Scroll down to "Custom Domains" section**
4. **Click "Add Custom Domain"**
5. **Enter your domain** (e.g., `api.yourdomain.com`)
6. **Click "Save"**

### Step 2: Configure DNS Records

Render will provide you with DNS instructions. Typically:

#### For Subdomain (Recommended)
```
Type: CNAME
Name: api
Value: your-app-name.onrender.com
```

#### For Root Domain
```
Type: A
Name: @
Value: [IP addresses provided by Render]
```

### Step 3: Update SERVER_ADDRESS Environment Variable

1. **Go to your Render service dashboard**
2. **Navigate to "Environment" tab**
3. **Find the `SERVER_ADDRESS` variable**
4. **Update the value to your custom domain**:
   ```
   SERVER_ADDRESS=https://api.yourdomain.com
   ```
5. **Click "Save Changes"**
6. **Redeploy your service**

### Step 4: Verify Configuration

1. **Wait for DNS propagation** (can take up to 48 hours)
2. **Visit your custom domain** to ensure it works
3. **Check the homepage** - it should now display your custom domain in the API base URL field

## 🔧 Configuration Examples

### Example 1: Subdomain Setup
```yaml
# In render.yaml or via Render dashboard
envVars:
  - key: SERVER_ADDRESS
    value: https://api.yourdomain.com
```

### Example 2: Root Domain Setup
```yaml
# In render.yaml or via Render dashboard
envVars:
  - key: SERVER_ADDRESS
    value: https://yourdomain.com
```

### Example 3: Development vs Production
```yaml
# Development
envVars:
  - key: SERVER_ADDRESS
    value: https://dev-api.yourdomain.com

# Production
envVars:
  - key: SERVER_ADDRESS
    value: https://api.yourdomain.com
```

## 🛡️ SSL/TLS Certificates

- ✅ **Automatic SSL**: Render automatically provides SSL certificates for custom domains
- ✅ **HTTPS Enforced**: All traffic is automatically redirected to HTTPS
- ✅ **Certificate Renewal**: Render handles automatic certificate renewal

## 🔍 Troubleshooting

### Domain Not Working
1. **Check DNS propagation**: Use tools like `dig` or online DNS checkers
2. **Verify CNAME/A records**: Ensure they point to the correct Render values
3. **Wait for propagation**: DNS changes can take up to 48 hours

### Homepage Still Shows Old URL
1. **Clear browser cache**: Force refresh the page
2. **Check environment variable**: Ensure `SERVER_ADDRESS` is updated
3. **Redeploy service**: Trigger a new deployment in Render

### SSL Certificate Issues
1. **Wait for provisioning**: SSL certificates can take a few minutes to provision
2. **Check domain validation**: Ensure DNS records are correct
3. **Contact Render support**: If issues persist

## 📱 User Experience After Custom Domain

Once configured, users will see:

- ✅ **Professional URL**: `https://api.yourdomain.com` instead of `.onrender.com`
- ✅ **Branded Experience**: Your domain in the API base URL field
- ✅ **Easy Integration**: Clients can use your custom domain for API calls
- ✅ **SSL Security**: Automatic HTTPS encryption

## 🎯 Best Practices

### Domain Naming
- ✅ Use subdomains: `api.yourdomain.com`
- ✅ Keep it short and memorable
- ✅ Use HTTPS-only domains

### DNS Configuration
- ✅ Use CNAME for subdomains (preferred)
- ✅ Set appropriate TTL values (300-3600 seconds)
- ✅ Monitor DNS propagation

### Environment Management
- ✅ Use different domains for different environments
- ✅ Update `SERVER_ADDRESS` immediately after domain setup
- ✅ Test thoroughly after changes

## 🔄 Updating Existing Deployments

If you already have MedusaXD API deployed and want to add a custom domain:

1. **Add custom domain in Render** (Steps 1-2 above)
2. **Update environment variable** (Step 3 above)
3. **Notify existing users** about the new API base URL
4. **Update documentation** with the new domain
5. **Consider keeping old domain** for backward compatibility during transition

## 📞 Support

If you encounter issues:

- 📖 **Render Documentation**: [Custom Domains Guide](https://render.com/docs/custom-domains)
- 💬 **Render Support**: Contact through Render dashboard
- 🐛 **MedusaXD Issues**: [GitHub Issues](https://github.com/MedusaXD/medusaxd-api/issues)

---

**🎉 Congratulations!** Your MedusaXD API now has a professional custom domain setup!
