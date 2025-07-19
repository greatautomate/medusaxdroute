<p align="right">
   <a href="./README.zh.md">中文</a> | <strong>English</strong>
</p>
<div align="center">

![MedusaXD API](/web/public/logo.png)

# MedusaXD API

🚀 Next-Generation AI Gateway & Asset Management System

<p align="center">
  <a href="https://raw.githubusercontent.com/MedusaXD/medusaxd-api/main/LICENSE">
    <img src="https://img.shields.io/github/license/MedusaXD/medusaxd-api?color=brightgreen" alt="license">
  </a>
  <a href="https://github.com/MedusaXD/medusaxd-api/releases/latest">
    <img src="https://img.shields.io/github/v/release/MedusaXD/medusaxd-api?color=brightgreen&include_prereleases" alt="release">
  </a>
  <a href="https://hub.docker.com/r/medusaxd/medusaxd-api">
    <img src="https://img.shields.io/badge/docker-dockerHub-blue" alt="docker">
  </a>
  <a href="https://goreportcard.com/report/github.com/MedusaXD/medusaxd-api">
    <img src="https://goreportcard.com/badge/github.com/MedusaXD/medusaxd-api" alt="GoReportCard">
  </a>
  <a href="https://render.com">
    <img src="https://img.shields.io/badge/deploy-render-purple" alt="Deploy on Render">
  </a>
</p>
</div>

## 📝 Project Description

> [!NOTE]
> MedusaXD API is an open-source project based on [One API](https://github.com/songquanpeng/one-api) with extensive enhancements and improvements.

> [!IMPORTANT]
> - This project is designed for production use with enterprise-grade features and reliability.
> - Users must comply with OpenAI's [Terms of Use](https://openai.com/policies/terms-of-use) and **applicable laws and regulations**. Do not use for illegal purposes.
> - Ensure compliance with local AI service regulations in your jurisdiction.

## 📚 Documentation

For detailed documentation, visit our official documentation site: [https://docs.medusaxd.com/](https://docs.medusaxd.com/)

Deploy instantly on Render.com with our one-click deployment guide.

## ✨ Key Features

MedusaXD API provides comprehensive functionality for AI gateway management. For detailed features, see our [Features Documentation](https://docs.medusaxd.com/features):

1. 🎨 Modern, intuitive UI interface
2. 🌍 Full internationalization support (English primary)
3. 💰 Flexible billing and payment integration
4. 🔍 Advanced usage analytics and monitoring
5. 🔄 Full backward compatibility with One API
6. 💵 Flexible pricing models (per-request, per-token)
7. ⚖️ Intelligent load balancing and channel weighting
8. 📈 Comprehensive analytics dashboard
9. 🔒 Advanced token management and model restrictions
10. 🤖 Multiple authentication providers (GitHub, OIDC, Telegram)
11. 🔄 Rerank model support (Cohere, Jina), [API Docs](https://docs.medusaxd.com/api/rerank)
12. ⚡ OpenAI Realtime API support (including Azure), [API Docs](https://docs.medusaxd.com/api/realtime)
13. ⚡ Claude Messages format support, [API Docs](https://docs.medusaxd.com/api/claude)
14. 🌐 Built-in chat interface with routing support
15. 🧠 Advanced reasoning effort control via model name suffixes:
    1. OpenAI o-series models
        - Add `-high` suffix for high reasoning effort (e.g., `o3-mini-high`)
        - Add `-medium` suffix for medium reasoning effort (e.g., `o3-mini-medium`)
        - Add `-low` suffix for low reasoning effort (e.g., `o3-mini-low`)
    2. Claude thinking models
        - Add `-thinking` suffix to enable thinking mode (e.g., `claude-3-7-sonnet-20250219-thinking`)
16. 🔄 Thinking-to-content conversion functionality
17. 🔄 Per-user model rate limiting
18. 💰 Advanced cache billing support with configurable ratios:
    1. Configure cache billing ratios in System Settings
    2. Set channel-specific cache ratios (0-1 range)
    3. Supported providers:
        - [x] OpenAI
        - [x] Azure
        - [x] DeepSeek
        - [x] Claude

## 🤖 Model Support

MedusaXD API supports a comprehensive range of AI models and services. For detailed API documentation, see [API Reference](https://docs.medusaxd.com/api):

1. **GPT Models** - Full OpenAI GPT model family including custom GPTs (gpt-4-gizmo-*)
2. **Image Generation** - [Midjourney-Proxy(Plus)](https://github.com/novicezk/midjourney-proxy) integration, [API Docs](https://docs.medusaxd.com/api/midjourney)
3. **Music Generation** - [Suno API](https://github.com/Suno-API/Suno-API) integration, [API Docs](https://docs.medusaxd.com/api/suno)
4. **Custom Channels** - Support for custom API endpoints with full URL configuration
5. **Rerank Models** - [Cohere](https://cohere.ai/) and [Jina](https://jina.ai/) reranking, [API Docs](https://docs.medusaxd.com/api/rerank)
6. **Claude Integration** - Full Claude Messages format support, [API Docs](https://docs.medusaxd.com/api/claude)
7. **Dify Integration** - Chatflow and agent support
8. **Azure OpenAI** - Complete Azure OpenAI service integration
9. **Anthropic** - Direct Anthropic API support
10. **Google AI** - Gemini and other Google AI models

## ⚙️ Environment Configuration

For detailed configuration instructions, see [Environment Variables Guide](https://docs.medusaxd.com/deployment/environment-variables):

### Core Settings
- `GENERATE_DEFAULT_TOKEN`: Generate initial tokens for new users (default: `false`)
- `STREAMING_TIMEOUT`: Streaming response timeout in seconds (default: `120`)
- `ERROR_LOG_ENABLED`: Enable error logging and display (default: `false`)
- `CRYPTO_SECRET`: Encryption key for database content security

### Database Configuration
- `DATABASE_URL`: PostgreSQL connection string for Render deployment
- `SQL_DSN`: Alternative database connection format

### Provider Settings
- `DIFY_DEBUG`: Output Dify workflow and node information (default: `true`)
- `COHERE_SAFETY_SETTING`: Cohere model safety level (`NONE`, `CONTEXTUAL`, `STRICT`, default: `NONE`)
- `GEMINI_VISION_MAX_IMAGE_NUM`: Maximum images for Gemini models (default: `16`)
- `AZURE_DEFAULT_API_VERSION`: Azure API version (default: `2025-04-01-preview`)

### Performance & Limits
- `FORCE_STREAM_OPTION`: Override client stream_options parameter (default: `true`)
- `GET_MEDIA_TOKEN`: Count image tokens (default: `true`)
- `GET_MEDIA_TOKEN_NOT_STREAM`: Count image tokens in non-streaming mode (default: `true`)
- `UPDATE_TASK`: Update async tasks (Midjourney, Suno) (default: `true`)
- `MAX_FILE_DOWNLOAD_MB`: Maximum file download size in MB (default: `20`)
- `NOTIFICATION_LIMIT_DURATION_MINUTE`: Notification rate limit duration (default: `10`)
- `NOTIFY_LIMIT_COUNT`: Maximum notifications per duration (default: `2`)

## 🚀 Deployment

MedusaXD API is optimized for cloud deployment with special focus on Render.com. For detailed deployment guides, see [Deployment Documentation](https://docs.medusaxd.com/deployment):

> [!TIP]
> **Recommended**: Deploy on Render.com with one-click setup!
> Latest Docker image: `medusaxd/medusaxd-api:latest`

### 🌟 Deploy on Render.com (Recommended)

[![Deploy to Render](https://render.com/images/deploy-to-render-button.svg)](https://render.com/deploy)

**Why Render.com? (Updated 2025)**
- ✅ **New Flexible PostgreSQL Plans**: Independent storage/compute scaling starting at $6/month
- ✅ **Point-in-Time Recovery**: Now included on ALL paid database plans
- ✅ Free PostgreSQL database (1GB, 30-day limit) for testing
- ✅ Automatic HTTPS and custom domains
- ✅ Zero-downtime deployments
- ✅ Built-in monitoring and logging
- ✅ Environment variable management

#### Quick Render Deployment
1. Fork this repository
2. Connect your GitHub account to Render
3. Create a new Web Service from your fork
4. Add a PostgreSQL database
5. Configure environment variables (see below)
6. Deploy!

### 🔧 Database Support
- **PostgreSQL** (Recommended for Render): Version >= 9.6
- **MySQL**: Version >= 5.7.8
- **SQLite**: For local development only

### 🐳 Alternative Deployment Methods

#### Docker Compose (Local Development)
```shell
# Clone the repository
git clone https://github.com/MedusaXD/medusaxd-api.git
cd medusaxd-api
# Edit docker-compose.yml as needed
# Start services
docker-compose up -d
```

#### Direct Docker Deployment
```shell
# With PostgreSQL (Render-compatible)
docker run --name medusaxd-api -d --restart always \
  -p 3000:3000 \
  -e DATABASE_URL="postgresql://user:password@host:5432/database" \
  -e TZ=UTC \
  medusaxd/medusaxd-api:latest

# With SQLite (local only)
docker run --name medusaxd-api -d --restart always \
  -p 3000:3000 \
  -v /path/to/data:/data \
  medusaxd/medusaxd-api:latest
```

## 🔄 Channel Retry & Caching

Advanced retry and caching mechanisms ensure high availability and performance. Configure retry settings in `Settings -> Operations -> General Settings`. **Caching is highly recommended** for production deployments.

### Cache Configuration
1. **Redis Cache**: Set `REDIS_CONN_STRING` for distributed caching
2. **Memory Cache**: Enable with `MEMORY_CACHE_ENABLED` (auto-enabled with Redis)
3. **Render Redis**: Use Render's Redis add-on for seamless integration

## 📖 API Documentation

Comprehensive API documentation available at [API Reference](https://docs.medusaxd.com/api):

### Core APIs
- [**Chat Completions**](https://docs.medusaxd.com/api/chat) - OpenAI-compatible chat interface
- [**Image Generation**](https://docs.medusaxd.com/api/image) - DALL-E and Midjourney support
- [**Embeddings**](https://docs.medusaxd.com/api/embeddings) - Text embedding generation
- [**Rerank**](https://docs.medusaxd.com/api/rerank) - Document reranking with Cohere/Jina
- [**Realtime**](https://docs.medusaxd.com/api/realtime) - OpenAI Realtime API support
- [**Claude Messages**](https://docs.medusaxd.com/api/claude) - Anthropic Claude integration

### Extended APIs
- [**Audio Processing**](https://docs.medusaxd.com/api/audio) - Speech-to-text and text-to-speech
- [**Music Generation**](https://docs.medusaxd.com/api/music) - Suno API integration
- [**Custom Models**](https://docs.medusaxd.com/api/custom) - Custom endpoint support

## 🔗 Related Projects

### Core Dependencies
- [**One API**](https://github.com/songquanpeng/one-api) - Original foundation project
- [**Midjourney-Proxy**](https://github.com/novicezk/midjourney-proxy) - Midjourney API integration
- [**Suno API**](https://github.com/Suno-API/Suno-API) - Music generation support

### Ecosystem
- [**ChatNIO**](https://github.com/Deeptrain-Community/chatnio) - Next-gen AI solution
- [**API Key Tools**](https://github.com/Calcium-Ion/neko-api-key-tool) - Usage monitoring utilities

### MedusaXD Extensions
- [**MedusaXD Dashboard**](https://github.com/MedusaXD/medusaxd-dashboard) - Advanced analytics dashboard
- [**MedusaXD CLI**](https://github.com/MedusaXD/medusaxd-cli) - Command-line management tools

## 🆘 Support & Community

Get help and connect with the community:

### 📚 Documentation & Guides
- [**Complete Documentation**](https://docs.medusaxd.com/) - Comprehensive guides and API docs
- [**Deployment Guide**](https://docs.medusaxd.com/deployment) - Step-by-step deployment instructions
- [**FAQ**](https://docs.medusaxd.com/faq) - Frequently asked questions

### 💬 Community Support
- [**GitHub Discussions**](https://github.com/MedusaXD/medusaxd-api/discussions) - Community Q&A
- [**GitHub Issues**](https://github.com/MedusaXD/medusaxd-api/issues) - Bug reports and feature requests
- [**Discord Server**](https://discord.gg/medusaxd) - Real-time community chat

### 🚀 Professional Support
- [**Enterprise Support**](mailto:enterprise@medusaxd.com) - Dedicated enterprise assistance
- [**Custom Development**](mailto:dev@medusaxd.com) - Custom feature development
- [**Consulting Services**](mailto:consulting@medusaxd.com) - AI integration consulting

## 🌟 Star History

[![Star History Chart](https://api.star-history.com/svg?repos=MedusaXD/medusaxd-api&type=Date)](https://star-history.com/#MedusaXD/medusaxd-api&Date)

---

<div align="center">

**Made with ❤️ by the MedusaXD Team**

[🌐 Website](https://medusaxd.com) • [📖 Docs](https://docs.medusaxd.com) • [💬 Discord](https://discord.gg/medusaxd) • [🐦 Twitter](https://twitter.com/medusaxd)

</div>
