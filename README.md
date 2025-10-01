# Log Crawler

A Go-based HTTP client application designed to test and monitor API endpoints while capturing detailed performance logs and request/response data.

## Description

Log Crawler is a specialized tool that:

- Makes HTTP requests to configured API endpoints
- Captures detailed timing information for performance analysis
- Logs client-side and server-side performance metrics
- Processes multiple services from JSON configuration files
- Saves all output to organized log files for analysis
- Extracts and formats performance tuning headers from responses

The application is particularly useful for:
- API performance monitoring
- Load testing scenarios
- Debugging API response times
- Collecting performance metrics across multiple services
- Automated API health checks

## Features

- **Configurable Services**: Define multiple API endpoints with different methods, paths, and parameters
- **Performance Logging**: Captures start/end times for both client and server-side operations
- **Response Analysis**: Processes performance tuning headers from API responses
- **Organized Output**: Saves logs to separate files based on module names
- **Flexible Configuration**: JSON-based configuration for easy service management
- **Error Handling**: Comprehensive error reporting and status code handling

## Project Structure

```
log-crawler/
├── main.go                    # Entry point - handles log file creation and command execution
├── cmd/
│   └── crawler/
│       └── main.go           # Core crawler logic and HTTP client functionality
├── internal/
│   └── common/
│       └── config.go         # Configuration structures and types
├── data/                     # JSON configuration files for different modules
│   ├── api-list-1.json
│   └── api-list-2.json
└── result/                   # Generated log files
    ├── api-list-1.log
    └── api-list-2.log
```

## Configuration Format

Each JSON configuration file in the `data/` directory should follow this structure:

```json
{
    "urlPrefix": "https://api.example.com/v1",
    "token": "Bearer your-auth-token-here",
    "services": [
        {
            "enable": true,
            "method": "POST",
            "path": "/endpoint/path",
            "param": "{\"key\":\"value\"}"
        }
    ]
}
```

### Configuration Fields

- **urlPrefix**: Base URL for all API endpoints
- **token**: Authorization token (typically Bearer token)
- **services**: Array of service configurations
  - **enable**: Boolean to enable/disable the service
  - **method**: HTTP method (GET, POST, PUT, DELETE, etc.)
  - **path**: API endpoint path (appended to urlPrefix)
  - **param**: JSON string containing request body parameters


## Installation

1. **Clone the repository:**
   ```powershell
   git clone <repository-url>
   cd log-crawler
   ```

2. **Install dependencies:**
   ```powershell
   go mod tidy
   ```

3. **Build the application:**
   ```powershell
   go build -o log-crawler.exe
   ```

## Usage

### Basic Usage

```powershell
# Run with a specific module configuration
go run main.go <moduleName>

# Or using the built executable
.\log-crawler.exe <moduleName>
```

### Examples

```powershell
# Test api-list-1 endpoints
go run main.go api-list-1

# Test api-list-2 services
go run main.go api-list-2
```

### Creating New Configurations

1. Create a new JSON file in the `data/` directory (e.g., `data/new-module.json`)
2. Configure your API endpoints following the JSON structure above
3. Run the crawler with your new module name:
   ```powershell
   go run main.go new-module
   ```

### Output

- Log files are automatically created in the `result/` directory
- Each module generates its own log file: `result/<moduleName>.log`


## Requirements

- Go 1.23.3 or later
- Network access to target API endpoints
- Valid authentication tokens for secured endpoints