# Go GIN API

## Built with

- [GIN](https://github.com/gin-gonic/gin) for routing engine.
- [NGINX](https://www.nginx.com/) as a request-response proxy-router / load balancer.
- [PostgreSQL](https://www.postgresql.org/) as the data storage layer.
- [PGAdmin](https://www.pgadmin.org/) for PostgreSQL management.
- [Docker](https://www.docker.com/)

## Table of Contents

* [Directory](#directory-structure)
* [Installation](#installation)
* [Todo](#todo)

## Directory Structure

* [App Directory](https://github.com/DLzer/go-gin-api-boilerplate/tree/main/app)
    * [App Config](https://github.com/DLzer/go-gin-api-boilerplate/tree/main/app/configs)
    * [API Config](https://github.com/DLzer/go-gin-api-boilerplate/blob/main/app/api/api.go)
    * [Middleware](https://github.com/DLzer/go-gin-api-boilerplate/tree/main/app/middleware)
* [NGINX Config](https://github.com/DLzer/go-gin-api-boilerplate/blob/main/nginx/nginx.conf)
* [Docker Config](https://github.com/DLzer/go-gin-api-boilerplate/blob/main/docker-compose.yml)

## Installation

1. Clone the repository

```bash
$ git clone https://github.com/DLzer/go-gin-api-boilerplate.git
```

2. Run docker build

```bash
$ cd go-gin-api-boilerplate
$ docker build
```

3. Run docker compose

```bash
$ docker compose up -d
```

4. Test the endpoints

```bash
$ curl http://localhost:80/events
{status: 200, message: "Events", data: []}
```

## Todo

- Testing
- Seeding
- Deployment workflow
- Linting setup