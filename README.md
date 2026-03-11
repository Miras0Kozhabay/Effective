

````markdown


REST API для управления онлайн-подписками пользователей.  

---

## 🚀 Стек

- Go (Gin framework)  
- PostgreSQL  
- Docker / Docker Compose  
- Swagger (OpenAPI 3.0)

---

## 📦 Запуск проекта

1. Клонируй репозиторий:
```bash
git clone https://github.com/Miras0Kozhabay/Effective.git
cd Effective
````

2. Создай `.env` с настройками БД (пример):

```env
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=subscriptions
DB_HOST=db
DB_PORT=5432
PORT=8080
```

3. Запусти Docker Compose:

```bash
docker compose up --build
```

4. Проверь сервис:

* API: `http://localhost:8080`
* Swagger UI: `http://localhost:8081`

---

## 📝 Основные эндпоинты

### Пинг

```
GET /ping
```

Проверка работы сервиса.
**Ответ:** `{ "message": "pong" }`

### Подписки

```
GET /subscriptions
POST /subscriptions
GET /subscriptions/{id}
PUT /subscriptions/{id}
DELETE /subscriptions/{id}
GET /subscriptions/total
```

### Фильтры и сортировка

`GET /subscriptions` поддерживает query-параметры:

* `user_id` — фильтр по пользователю
* `service_name` — фильтр по сервису
* `min_price`, `max_price` — диапазон цен
* `sort` — сортировка: `price_asc`, `price_desc`, `start_date_asc`, `start_date_desc`

Пример:

```
GET /subscriptions?user_id=60601fee-2bf1-4721-ae6f-7636e79a0cba&sort=price_desc
```

---

## ⚙️ Возможности проекта

* CRUD для подписок
* Суммарная стоимость подписок с фильтрами по пользователю, сервису и дате
* Фильтры и сортировка подписок
* Документация API через Swagger UI

---

## 🔧 Структура проекта

```
cmd/             # main.go
internal/
  config/        # загрузка .env
  handler/       # обработчики запросов
  repository/    # работа с БД
  service/       # бизнес-логика (пока что пустая для данного проекта)
  model/         # структуры данных
migrations/      # SQL миграции
swagger.yaml     # OpenAPI спецификация
Dockerfile
docker-compose.yml
```

