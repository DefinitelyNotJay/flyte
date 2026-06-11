# Flyte Monorepo

This project is structured as a monorepo containing both the backend and frontend services.

## Structure

- `backend/`: Go (Gin) REST API with PostgreSQL and PASETO authentication.
- `frontend/`: NuxtJS (Vue 3) frontend application using Bun.

## Getting Started

### Prerequisites

- [Go](https://golang.org/) 1.25+
- [Bun](https://bun.sh/)
- [Node.js](https://nodejs.org/) v22+ (for Nuxt)
- [Docker](https://www.docker.com/) (for PostgreSQL)
- [Task](https://taskfile.dev/) (optional, but recommended)

### Installation

Install both backend and frontend dependencies:

```bash
task install
```

### Development

Run both backend and frontend in development mode:

```bash
task
```

Or run them separately:

```bash
task backend:server
task frontend:dev
```

### Testing

Run backend tests:

```bash
task backend:test
```

## Backend Commands

Refer to `backend/Makefile` for low-level backend tasks (migrations, sqlc, mocks).
