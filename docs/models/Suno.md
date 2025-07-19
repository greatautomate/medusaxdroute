# Suno API Documentation

**Introduction**: Suno API Documentation

## API Endpoints
Supported endpoints:
+ [x] /suno/submit/music
+ [x] /suno/submit/lyrics
+ [x] /suno/fetch
+ [x] /suno/fetch/:id

## Model List

### Suno API Support

- suno_music (Custom mode, Inspiration mode, Continue)
- suno_lyrics (Generate lyrics)


## Model Pricing Settings (Set in Settings-Operations-Model Fixed Price Settings)
```json
{
  "suno_music": 0.3,
  "suno_lyrics": 0.01
}
```

## Channel Configuration

### Connecting to Suno API

1. Deploy Suno API and configure Suno account (strongly recommend setting API key), [Project Repository](https://github.com/Suno-API/Suno-API)

2. Add channel in channel management, select channel type **Suno API**, refer to the model list above for models
3. Fill in **Proxy** with the Suno API deployment address, e.g.: http://localhost:8080
4. Fill in the API key for Suno API, if no key is set, you can fill in anything

### Connecting to Upstream MedusaXD API

1. Add channel in channel management, select channel type **Suno API**, or any type, just need models to include the models from the list above
2. Fill in **Proxy** with the upstream MedusaXD API address, e.g.: http://localhost:3000
3. Fill in the API key for upstream MedusaXD API