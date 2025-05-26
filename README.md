# Quotes API Service

Мини-сервис "Цитатник" - это REST API для хранения и управления цитатами, реализованный на Go.

## Функциональные возможности

- Добавление новых цитат
- Просмотр всех цитат
- Получение случайной цитаты
- Фильтрация цитат по автору
- Удаление цитат по ID

## Требования

- Go 1.21+
- Git (для клонирования репозитория)

## Установка и запуск

1. Клонируйте репозиторий:
```bash
git clone https://github.com/your-username/quotesAPI.git
cd quotesAPI
```

2. (Опционально) Создайте файл `.env` для настройки порта:
```bash
echo "PORT=8080" > .env
```

3. Соберите и запустите сервер:
```bash
go run cmd/main.go
```

Сервер запустится на порту, указанном в переменной окружения `PORT` (по умолчанию 8080).

## Конфигурация

Сервис поддерживает настройку через:
1. Файл `.env` в корне проекта:
   ```
   PORT=8080
   ```
2. Переменные окружения:
   ```bash
   export PORT=8080
   ```

## Использование API

### Добавление цитаты
```bash
curl -X POST http://localhost:8080/quotes \
  -H "Content-Type: application/json" \
  -d '{"author":"Confucius", "quote":"Life is simple, but we insist on making it complicated."}'
```

### Получение всех цитат
```bash
curl http://localhost:8080/quotes
```

### Получение случайной цитаты
```bash
curl http://localhost:8080/quotes/random
```

### Фильтрация по автору
```bash
curl http://localhost:8080/quotes?author=Confucius
```

### Удаление цитаты
```bash
curl -X DELETE http://localhost:8080/quotes/1
```


## Запуск тестов

Для запуска unit-тестов выполните:
```bash
go test -v ./...
```

Для проверки покрытия тестами:
```bash
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```

## Особенности реализации

- Конфигурация через .env файл или переменные окружения
- Graceful shutdown при получении сигналов завершения
- Таймауты для HTTP сервера
- Хранение данных в памяти
- Чистая архитектура с разделением на слои
- Подробное логирование запросов

