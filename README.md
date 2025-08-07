# RSS Aggregator API

> **A modular, production-ready REST API for aggregating and managing RSS feed articles**

This Go-based RSS aggregator provides a clean, scalable architecture for collecting articles from multiple RSS feeds, storing them in Firebase, and serving them through a comprehensive REST API. Built with modularity and maintainability in mind.

## 🚀 Key Features

- **Smart Aggregation**: Automatically collects articles from configured RSS feeds with duplicate detection
- **Flexible Retrieval**: Multiple endpoints for accessing articles (all, latest, random, search)
- **Source Management**: Dynamic RSS feed source management via API
- **Real-time Storage**: Firebase Realtime Database integration with robust error handling
- **Clean Architecture**: Modular codebase with separated concerns (handlers, shared utilities)
- **Comprehensive API**: OpenAPI 3.0.3 specification with detailed documentation

## 🏗️ Architecture Overview

### Project Structure
```
├── main.go                 # Application entry point & Firebase initialization
├── handlers/               # Modular route handlers
│   ├── aggregator.go      # RSS feed aggregation logic
│   ├── articles.go        # Article retrieval endpoints
│   ├── sources.go         # RSS feed source management
│   └── shared.go          # Shared types, variables, and utilities
├── openapi.json           # Complete API documentation
├── credentials.json       # Firebase service account credentials
└── go.mod                 # Go module dependencies
```

### Technical Stack
- **Runtime**: Go 1.19+
- **Database**: Firebase Realtime Database
- **RSS Parser**: [gofeed](https://github.com/mmcdole/gofeed) library
- **Firebase SDK**: Firebase Admin SDK for Go
- **Documentation**: OpenAPI 3.0.3 specification

## 🚀 Getting Started

### Prerequisites
- **Go 1.19 or higher**
- **Firebase project** with Realtime Database enabled
- **Service account credentials** (credentials.json file)

### Quick Setup

1. **Clone and navigate to the project**:
   ```bash
   cd ./apip-learners/apis/eli/GoAggregateRSS
   ```

2. **Install dependencies**:
   ```bash
   go mod tidy
   ```

3. **Configure Firebase**:
   - Place your `credentials.json` file in the project root
   - Update Firebase config in `main.go` if needed (project ID, database URL)

4. **Run the application**:
   ```bash
   go run main.go
   ```

5. **Verify installation**:
   ```bash
   curl http://localhost:8080/
   # Expected: "Welcome to the RSS Aggregator!"
   ```

### First Steps After Setup

1. **Test connectivity**: `GET /` - Should return welcome message
2. **Aggregate feeds**: `POST /agg` - Collects articles from all configured sources
3. **View articles**: `GET /articles` - See all collected articles
4. **Check sources**: `GET /sources` - View configured RSS feeds

## 📚 API Documentation

### Core Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| `GET` | `/` | Welcome message |
| `POST` | `/agg` | Aggregate articles from RSS feeds |
| `GET` | `/articles` | Retrieve all articles |
| `GET` | `/articles/latest?amount=N` | Get latest N articles |
| `GET` | `/articles/rand?amount=N` | Get N random articles |
| `GET` | `/articles/search?keyword=X` | Search articles by keyword |
| `GET` | `/sources` | List RSS feed sources |
| `PUT` | `/sources` | Add new RSS feed source |
| `DELETE` | `/sources` | Remove RSS feed source |

### Example Usage

**Aggregate new articles**:
```bash
curl -X POST http://localhost:8080/agg
```

**Get latest 5 articles**:
```bash
curl "http://localhost:8080/articles/latest?amount=5"
```

**Search for AI-related articles**:
```bash
curl "http://localhost:8080/articles/search?keyword=AI"
```

**Add new RSS source**:
```bash
curl -X PUT http://localhost:8080/sources \
  -H "Content-Type: application/json" \
  -d '{"source": "https://example.com/feed.xml"}'
```

### Response Format

All endpoints return consistent JSON responses:
```json
{
  "status": 200,
  "message": "Operation description",
  "count": 5,
  "articles": {
    "article_id": {
      "title": "Article Title",
      "link": "https://example.com/article",
      "content": "Article content...",
      "publishedDate": 1672531200000,
      "receivedDate": 1672531300000,
      "source": "https://source-feed.com/rss"
    }
  }
}
```

## 🔗 Pre-configured RSS Sources

The application comes with 13 curated RSS feeds covering:
- **Tech News**: TechCrunch, Wired AI, Google Cloud Blog
- **Development**: Web Developer Juice, TechRepublic
- **Business**: Startup news, SaaS industry updates
- **Cloud Computing**: Cloud computing news, cloud security
- **Video Content**: YouTube channel feeds
- **Newsletter**: Campaign Archive feeds

## 📋 OpenAPI Specification

Complete API documentation is available in `openapi.json`, including:
- Detailed endpoint descriptions
- Parameter specifications
- Response schemas
- Example requests and responses
- Error handling documentation

## 🛠️ Development Notes

### Adding New Endpoints
1. Create handler function in appropriate file under `handlers/`
2. Register route in `main.go`
3. Update OpenAPI specification
4. Add tests and documentation

### Extending RSS Sources
Modify the `MyBlogs` slice in `main.go` or use the `/sources` API endpoints for dynamic management.

### Error Monitoring
All operations include comprehensive logging. Check application logs for detailed error information and debugging.

---

**Built with ❤️ using Go and Firebase**