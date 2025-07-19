# Database Configuration for MedusaXD API

MedusaXD API supports multiple database backends with PostgreSQL being the recommended choice for production deployments, especially on Render.com.

## Supported Databases

### PostgreSQL (Recommended)
- **Production Ready**: Excellent performance and reliability
- **Render.com Native**: First-class support on Render platform
- **Features**: Full SQL support, ACID compliance, excellent concurrency
- **Versions**: 12, 13, 14, 15 (recommended)

### MySQL
- **Compatibility**: Version 5.7.8 or higher
- **Use Case**: Legacy systems or specific requirements
- **Performance**: Good for read-heavy workloads

### SQLite
- **Development Only**: Not recommended for production
- **Use Case**: Local development and testing
- **Limitations**: No concurrent writes, single file storage

## PostgreSQL on Render.com

### Database Plans (Updated 2025)

Render.com now offers **flexible plans** with independent storage and compute scaling.

#### Free Tier
- **Storage**: 1GB (fixed)
- **RAM**: 256MB
- **CPU**: 0.1
- **Connections**: 100 concurrent
- **Retention**: 30 days (database expires)
- **Backups**: None
- **Cost**: $0/month

#### Basic Tier (New)
- **Basic-256mb**: $6/month (0.1 CPU, 256MB RAM, 100 connections)
- **Basic-1gb**: $19/month (0.5 CPU, 1GB RAM, 100 connections)
- **Basic-4gb**: $75/month (2 CPU, 4GB RAM, 100 connections)
- **Storage**: $0.30/GB/month (independent scaling)
- **Features**: Point-in-time recovery, logical backups

#### Pro Tier (Production Ready)
- **Pro-4gb**: $55/month (1 CPU, 4GB RAM, 100 connections)
- **Pro-8gb**: $100/month (2 CPU, 8GB RAM, 200 connections)
- **Pro-16gb**: $200/month (4 CPU, 16GB RAM, 400 connections)
- **Pro-32gb**: $400/month (8 CPU, 32GB RAM, 500 connections)
- **Storage**: $0.30/GB/month (independent scaling)
- **Features**: High availability, read replicas, point-in-time recovery

#### Accelerated Tier (Memory-Intensive)
- **Accelerated-16gb**: $160/month (2 CPU, 16GB RAM, 400 connections)
- **Accelerated-32gb**: $350/month (4 CPU, 32GB RAM, 500 connections)
- **Accelerated-64gb**: $750/month (8 CPU, 64GB RAM, 500 connections)
- **Storage**: $0.30/GB/month (independent scaling)
- **Features**: 1:8 CPU-to-RAM ratio, optimized for memory workloads

### Connection Configuration

#### Environment Variables

```bash
# Primary connection string (auto-provided by Render)
DATABASE_URL=postgresql://username:password@hostname:port/database

# Alternative format
SQL_DSN=username:password@tcp(hostname:port)/database?sslmode=require
```

#### Connection Parameters

```bash
# SSL Configuration (required on Render)
POSTGRES_SSL_MODE=require

# Connection Pool Settings
POSTGRES_MAX_CONNECTIONS=20
POSTGRES_IDLE_TIMEOUT=300
POSTGRES_MAX_LIFETIME=3600
```

### Database Initialization

MedusaXD API automatically handles database initialization:

1. **Schema Creation**: Tables are created automatically
2. **Migrations**: Database schema updates are applied on startup
3. **Indexes**: Performance indexes are created automatically
4. **Constraints**: Foreign keys and constraints are established

### Performance Optimization

#### Connection Pooling

```go
// Recommended connection pool settings
maxOpenConns := 20
maxIdleConns := 5
connMaxLifetime := time.Hour
```

#### Query Optimization

- Use prepared statements for repeated queries
- Implement proper indexing strategies
- Monitor slow query logs
- Use connection pooling

#### Monitoring

```bash
# Enable query logging
POSTGRES_LOG_QUERIES=true

# Log slow queries (>1 second)
POSTGRES_SLOW_QUERY_THRESHOLD=1000
```

## Database Schema

### Core Tables

#### Users
```sql
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(255) UNIQUE NOT NULL,
    email VARCHAR(255) UNIQUE,
    password_hash VARCHAR(255),
    role INTEGER DEFAULT 1,
    status INTEGER DEFAULT 1,
    quota BIGINT DEFAULT 0,
    used_quota BIGINT DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

#### Channels
```sql
CREATE TABLE channels (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    type INTEGER NOT NULL,
    key TEXT,
    base_url TEXT,
    other TEXT,
    models TEXT,
    status INTEGER DEFAULT 1,
    priority INTEGER DEFAULT 0,
    weight INTEGER DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

#### Tokens
```sql
CREATE TABLE tokens (
    id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES users(id),
    name VARCHAR(255),
    key VARCHAR(255) UNIQUE NOT NULL,
    status INTEGER DEFAULT 1,
    quota BIGINT DEFAULT 0,
    used_quota BIGINT DEFAULT 0,
    models TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

### Indexes

```sql
-- Performance indexes
CREATE INDEX idx_users_username ON users(username);
CREATE INDEX idx_users_email ON users(email);
CREATE INDEX idx_tokens_key ON tokens(key);
CREATE INDEX idx_tokens_user_id ON tokens(user_id);
CREATE INDEX idx_channels_type ON channels(type);
CREATE INDEX idx_channels_status ON channels(status);
```

## Backup and Recovery

### Render.com Backups

- **Automatic**: Daily backups included
- **Retention**: Varies by plan (7-30 days)
- **Recovery**: Point-in-time recovery available
- **Download**: Backup files can be downloaded

### Manual Backup

```bash
# Create backup
pg_dump $DATABASE_URL > backup.sql

# Restore backup
psql $DATABASE_URL < backup.sql
```

### Backup Strategy

1. **Regular Backups**: Automated daily backups
2. **Pre-deployment**: Backup before major updates
3. **Testing**: Regular restore testing
4. **Monitoring**: Backup success monitoring

## Migration Management

### Automatic Migrations

MedusaXD API handles migrations automatically:

```go
// Migration on startup
func init() {
    db.AutoMigrate(&User{}, &Channel{}, &Token{}, &Log{})
}
```

### Manual Migration

```bash
# Run specific migration
./medusaxd-api --migrate

# Check migration status
./medusaxd-api --migrate-status
```

## Security

### Access Control

```bash
# Database user permissions
GRANT CONNECT ON DATABASE medusaxd_api TO medusaxd_user;
GRANT USAGE ON SCHEMA public TO medusaxd_user;
GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA public TO medusaxd_user;
```

### SSL/TLS

```bash
# Force SSL connections
POSTGRES_SSL_MODE=require

# Verify SSL certificate
POSTGRES_SSL_VERIFY=true
```

### Data Encryption

- **At Rest**: Render provides encryption at rest
- **In Transit**: SSL/TLS encryption required
- **Application**: Sensitive data encrypted in application

## Monitoring and Maintenance

### Health Checks

```sql
-- Database health check
SELECT 1;

-- Connection count
SELECT count(*) FROM pg_stat_activity;

-- Database size
SELECT pg_size_pretty(pg_database_size('medusaxd_api'));
```

### Performance Monitoring

```sql
-- Slow queries
SELECT query, mean_time, calls 
FROM pg_stat_statements 
ORDER BY mean_time DESC 
LIMIT 10;

-- Table sizes
SELECT schemaname, tablename, 
       pg_size_pretty(pg_total_relation_size(schemaname||'.'||tablename)) as size
FROM pg_tables 
ORDER BY pg_total_relation_size(schemaname||'.'||tablename) DESC;
```

### Maintenance Tasks

```sql
-- Update statistics
ANALYZE;

-- Vacuum tables
VACUUM ANALYZE;

-- Reindex
REINDEX DATABASE medusaxd_api;
```

## Troubleshooting

### Common Issues

1. **Connection Timeouts**
   - Check connection pool settings
   - Verify network connectivity
   - Review SSL configuration

2. **Performance Issues**
   - Analyze slow queries
   - Check index usage
   - Monitor connection pool

3. **Migration Failures**
   - Check database permissions
   - Verify schema compatibility
   - Review migration logs

### Debugging

```bash
# Enable debug logging
POSTGRES_DEBUG=true

# Connection testing
psql $DATABASE_URL -c "SELECT version();"

# Performance analysis
EXPLAIN ANALYZE SELECT * FROM users WHERE username = 'test';
```
