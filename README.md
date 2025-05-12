# 📦 Go XML to MySQL Sample Project

This project demonstrates how to build a Go application that reads structured data from an XML file and inserts it into a MySQL database.

The MySQL server is containerized using Docker and runs on a Raspberry Pi, making this project suitable for lightweight, portable deployments.

---

## 🚀 Features

- Parses XML data using Go's standard `encoding/xml` package.
- Connects to MySQL using the `database/sql` and `go-sql-driver/mysql` packages.
- Configurable via environment variables.
- Lightweight and suitable for running on devices like Raspberry Pi.

---

## 🧱 Project Architecture

/yourapp
├── main.go # Entry point
├── handlers/ # HTTP API handlers (optional)
├── services/ # Business logic
├── repositories/ # Database access
├── models/ # Structs for XML and DB mapping
├── config/ # Environment and configuration loading
├── database/ # MySQL connection logic
└── xml/ # XML parsing logic

Yaml

---

## 🐳 Docker & MySQL Setup

The MySQL database is hosted in a Docker container on a Raspberry Pi.

### Sample `docker-compose.yml` for MySQL:

```yaml
version: '3.8'

services:
  db:
    image: mysql:8.0
    container_name: mysql-db
    restart: always
    environment:
      MYSQL_ROOT_PASSWORD: root
      MYSQL_DATABASE: myapp
      MYSQL_USER: user
      MYSQL_PASSWORD: password
    ports:
      - "3306:3306"
    volumes:
      - db_data:/var/lib/mysql

volumes:
  db_data:
