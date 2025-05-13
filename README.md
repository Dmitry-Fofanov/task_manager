# Менеджер задач

## Запуск проекта

### 1. Клонирование репозитория
```bash
git clone https://github.com/Dmitry-Fofanov/task_manager.git
cd task_manager
```

### 2. Настройка окружения

Создать файл `.env`. Можно использовать пример из `.env.example`:

```bash
cp .env.example .env
```

### 3. Запуск контейнеров

```bash
docker compose up
```


## Проверка работы API

### Создание задачи

```bash
curl --request POST \
     --url 'http://localhost/tasks/' \
     --header 'Content-Type: application/json' \
     --data '{"title": "Название задачи"}'
```

После выполнения запроса будет получен ответ
```json
{
    "id": 1,
    "title": "Название задачи",
    "status": "new",
    "created_at": "*время создания задачи*",
    "updated_at": "*время создания задачи*"
}
```


### Получение списка задач

```bash
curl 'http://localhost/tasks/'
```

После выполнения запроса будет получен ответ
```json
{
    "tasks": [
        {
            "id": 1,
            "title": "Название задачи",
            "status": "new",
            "created_at": "*время создания задачи*",
            "updated_at": "*время создания задачи*"
        }
    ]
}
```

### Обновление данных задачи

```bash
curl --request PUT \
     --url 'http://localhost/tasks/1/' \
     --header 'Content-Type: application/json' \
     --data '{
           "title": "Название задачи",
           "description": "Описание задачи",
           "status": "in_progress"
         }'
```

После выполнения запроса будет получен ответ
```
{
    "id": 1,
    "title": "Название задачи",
    "description": "Описание задачи",
    "status": "in_progress",
    "created_at": "*время создания задачи*",
    "updated_at": "*время изменения задачи*"
}
```

### Удаление задачи

```bash
curl --request DELETE -i \
     --url 'http://localhost/tasks/1/'
```

После выполнения запроса будет получен пустой ответ со статусом
```plaintext
HTTP/1.1 204 No Content
```
