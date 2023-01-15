# URL-SHORTERN

## Overview
- This is a web application that utilizes Fiber (web framework), Gorm (ORM for Go), Controller and Service pattern (architecture), PostgreSQL (database), and Redis (cache).

## How it works
- When a URL is fetched, the application will save the data in Redis as a cache to improve performance.

## Requirements
- Go (version 1.15 or higher)
- PostgreSQL (version 12 or higher)
- Redis (version 6 or higher)

## Installation
1. Clone the repository
2. Run `go get` to install dependencies
3. Update the configuration in the `config.yml` file to match your local setup
4. Make Docker and Redis `make docker` and `make redis`
5. Run `go run main.go` to start the application

## Contributing
- We welcome contributions to this project. Please submit a pull request or open an issue to start a discussion.