# URL Shortener

Сервис сокращения ссылок на Go.

## Стек

- Go 1.22, chi (роутер)
- PostgreSQL 16
- Docker, docker-compose

## Запуск

```bash
docker-compose up --build
```

Сервис будет доступен на `http://localhost:8080`.

## API

### Создать короткую ссылку

```bash
curl -X POST http://localhost:8080/api/shorten \
  -H "Content-Type: application/json" \
  -d '{"url": "https://github.com"}'
```

Ответ:

```json
{
  "result": "http://localhost:8080/aB3kZ1"
}
```

### Перейти по короткой ссылке

```bash
curl -L http://localhost:8080/aB3kZ1
```

## Структура проекта

```
├── cmd/shortener/main.go          # точка входа
├── internal/
│   ├── handler/handler.go         # HTTP-хендлеры
│   ├── storage/postgres.go        # слой работы с БД
│   └── shortid/shortid.go         # генерация коротких кодов
├── migrations/001_create_urls.sql # SQL-миграция
├── docker-compose.yml
├── Dockerfile
└── README.md
```
