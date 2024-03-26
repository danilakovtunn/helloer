# Сервис "приветствий"

### Порядок установки
Сборка контейнера приложения
```bash
docker build --tag helloer .
```

Запуск приложения и бд
```bash
docker compose up -d
```

Накатка миграций
```bash
bash deploy/migrate.sh up
```

### Ручки сервиса

Всего 2 ручки:
1. Положить в сервис кастомное приветствие
```bash
curl http://127.0.0.1:8080/my_hello -d 'Привет всем, кого не видел'
```
2. Вывести все приветствия
```bash
curl http://127.0.0.1:8080/hellos
```

Выполнил студент 528 группы, Ковтун Данила