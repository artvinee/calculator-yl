# yl-calc

yl-calc — это сервис-калькулятор на Go, который предоставляет HTTP API для вычисления математических выражений. Сервис парсит выражения, преобразует их в обратную польскую нотацию (RPN) и вычисляет результат.

## Быстрый запуск

```cmd
go run ./main.go
```

## Использование API

Сервис предоставляет следующую конечную точку:

### POST /api/v1/calculate

## Возможности

- Поддержка базовых арифметических операций: сложение (`+`), вычитание (`-`), умножение (`*`), деление (`/`).
- Обработка скобок для определения приоритета операций.
- Поддержка унарных операторов.
- Возвращает результаты в формате JSON.

## Структура проекта
```
yl-calc/
├── testing/
│   ├── calc_test.go
├── core/
│   ├── calc.go
│   ├── errors.go
│   ├── server.go
│   └── structs.go
├── go.mod
└── main.go
```
#### core/calc.go:
- Содержит основную логику для разбора и вычисления математических выражений.
- Реализует функции tokenize, toRPN и evaluateRPN, которые отвечают за разбиение выражения на токены, преобразование в обратную польскую нотацию и вычисление результата соответственно.
#### core/errors.go:
- Содержит определения пользовательских ошибок.
#### core/server.go:
- Настраивает HTTP сервер и определяет конечную точку API /api/v1/calculate.
- Обрабатывает входящие запросы, декодирует JSON данные и возвращает результат вычисления или сообщение об ошибке.
#### core/structs.go:
- Определяет структуры данных, используемые для тел запросов и ответов.
#### testing/calc_test:
- Cодержит набор тестов для функции CalculateExpression, проверяющих корректность вычисления как валидных, так и невалидных математических выражений.
- Тесты включают сценарии с ожидаемыми результатами и ошибками, чтобы убедиться в правильной обработке различных случаев.
#### main.go:
- Точка входа в приложение.
- Вызывает StartServer из server.go для инициализации и запуска HTTP сервера.

#### Формат запроса

**Content-Type:** `application/json`

**Тело:**
```json
{
  "expression": "ваше_выражение_здесь"
}
```

#### Ответ

**Успех (200)**
```json
{
  "result": вычисленное_значение
}
```

**Ошибка (422 Unprocessable Entity)**
```json
{
  "error": "сообщение_об_ошибке"
}
```

#### Примеры

**Успешное вычисление**
```curl
curl --location 'localhost:8080/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{
  "expression": "2+2*2"
}'
```

**Ошибка 422**
```curl
curl --location 'localhost:8080/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{
  "expression": "2 / 0"
}'
```
