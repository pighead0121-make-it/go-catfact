# ABOUT
A microservice built in Go that fetches catfacts from an exteral API, decorates the business logic with logging and stores request in postgreSQL.

# TECH STACK
Go (1.26), PostgreQL (17), Docker

# QUICK START

Start all services (PostgreSQL + API):

```bash
docker compose up -d --build
```

Fact a fact:
```bash
curl http://localhost:3000/fact
```

Expected Response:
```json
{
  "fact": "Cats sleep 70% of their lives.",
}
```

Check logged requests in APIserver:
```bash
docker logs catfact-api
```

Inspect stored requests in PostgreSQL:
```bash
docker compose exec -it postgres psql -U postgres -d catfactdb -c "SELECT * FROM request_logs ORDER BY created_at DESC LIMIT 5;"
```