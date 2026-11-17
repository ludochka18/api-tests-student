# API-тесты для проекта https://testboard.avito.com/

## Полезные ссылки

1. [Swagger-спецификация](https://testboard.avito.com/swagger/index.html)
2. [Техническое задание](https://docs.google.com/document/d/1IZ9SxEcePoQT6lyEmukycp1qNok4HZtok73UQnPwKSA/edit?usp=sharing)
3. [Библиотека для написания API-тестов](https://github.com/steinfletcher/apitest)
4. [Мощная библиотека анализа JSON для Go](https://github.com/tidwall/gjson)

## Запуск тестов

1. Установите любое IDE для разработки: [Jetbrains GoLand](https://www.jetbrains.com/ru-ru/go/) (Платное или по студенческой лицензии), [VS Code](https://code.visualstudio.com/) (бесплатное).
2. Установите [Git](https://git-scm.com/install/) на компьютер
3. [Установите](https://go.dev/doc/install) Go версии 1.25.1 или выше
4. Я использую клонирование по HTTPS, поэтому SSH-ключ для этого не требуется.
5. Склонируйте себе репозиторий: `git clone <путь к вашему репозиторию>`
6. Подтяните необходимые зависимости через команду в корне проекта:

```bash
   go mod tidy
```

7. В файле `.env.override` (создайте его в корне, если нет) выставьте правильные значения для переменных окружения:

```
TEST_LOGIN=логин_первого_аккаунта
TEST_PASSWORD=пароль_первого_аккаунта
TEST_SECOND_LOGIN=логин_второго_аккаунта
TEST_SECOND_PASSWORD=пароль_второго_аккаунта
```

Для тестов избранного я использую два разных аккаунта TestBoard: первый создаёт
объявления, а второй добавляет их в избранное. Для запуска укажите данные обоих
аккаунтов в локальном файле `.env.override`. Реальные логины и пароли я не
добавляю в README и репозиторий.

8. Для запуска всех тестов достаточно выполнить команду в корне проекта:

```bash
  go test -v ./...
```

9. Отдельный пакет с тестами можно запустить командой:

```bash
  go test -v ./tests/scenarios/favorites/
```

## Логирование

Для включения логирования request/response в консоль: выставьте переменную окружения `DEBUG=true` в
`.env.override`

## Запуск линтера

1. Установить приложение [golangci-lint](https://golangci-lint.run/docs/welcome/install/local/)
2. Запуск линтера (из корня репозитория):

```bash
  golangci-lint run
```

Для проверки повторяемости я запускаю `go test ./... -count=3`. Тесты обращаются к TestBoard и требуют доступности сервиса: ответ HTTP 503 означает, что проверка не прошла. Если очистка завершилась с ошибкой, я проверяю оставшиеся тестовые объявления и удаляю их вручную.

