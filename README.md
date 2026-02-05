# Matrix RSS Bot

Matrix RSS Bot is a lightweight service that monitors one or more RSS/Atom feeds and posts new entries to a specified [Matrix](https://matrix.org/) room. It is designed for easy deployment, including Docker support, and is ideal for communities or individuals who want to receive feed updates directly in their Matrix chat rooms.

## Features

- Monitors multiple RSS/Atom feeds
- Posts new entries to a Matrix room as formatted messages
- Simple configuration via environment variables
- Docker and docker-compose support for easy deployment

## How It Works

The bot periodically checks the configured RSS/Atom feeds. When a new entry is detected, it sends a message to the specified Matrix room using the Matrix Client-Server API.

## Quick Start

### 1. Clone the Repository

```sh
git clone https://github.com/Fingo2409/matrix-rss.git
cd matrix-rss
```

### 2. Configuration

Copy the example environment file and edit it to your needs:

```sh
cp env-example .env
# Edit .env with your feed URLs and Matrix credentials
```

**Environment variables:**

| Variable         | Description                                 | Example                        |
|------------------|---------------------------------------------|--------------------------------|
| FEED_URLS        | Comma-separated list of RSS/Atom feed URLs  | https://example.com/feed.xml   |
| MATRIX_SERVER    | Matrix homeserver URL                       | https://matrix.org             |
| MATRIX_ROOM_ID   | Matrix room ID (e.g. !roomid:matrix.org)    | !abcdef:matrix.org             |
| MATRIX_TOKEN     | Matrix access token                         | syt_xxx...                     |
| CHECK_INTERVAL   | Feed check interval in minutes (integer)    | 60                             |

### 3. Run with Docker

```sh
docker-compose up -d
```

Or build and run manually:

```sh
docker build -t matrix-rss .
docker run --env-file .env --network host ghcr.io/fingo2409/matrix-rss:latest
```

### 4. Run Locally (Go)

Ensure you have Go 1.23+ installed:

```sh
cd src
go build -o matrix-rss .
./matrix-rss
```

## Example docker-compose.yml

```yaml
services:
	matrix-rss:
		image: ghcr.io/fingo2409/matrix-rss:latest
		container_name: matrix-rss
		restart: unless-stopped
		network_mode: host
		environment:
			FEED_URLS: ${FEED_URLS}
			MATRIX_SERVER: ${MATRIX_SERVER}
			MATRIX_ROOM_ID: ${MATRIX_ROOM_ID}
			MATRIX_TOKEN: ${MATRIX_TOKEN}
			CHECK_INTERVAL: ${CHECK_INTERVAL}
```

## License

This project is licensed under the MIT License. See [LICENSE.md](LICENSE.md) for details.

---

*Made with ❤️ by Fingo2409*
