# Сервис банковских операций

# Usage

make build - билд приложения

make compose-up - запуск приложения

make migrate-up - применение миграций

# Примеры запросов

Регистрация

POST   /auth/sign-up
```curl --location --request POST 'http://localhost:8080/auth/sign-up' \
--header 'Content-Type: application/json' \
--data-raw '{
    "username":"TestUser123456",
    "password":"123456"
}
```
Пример ответа:
```
{
    "id": 1
}
```
Аутентификация

POST   /auth/sign-in
```
curl --location --request POST 'http://localhost:8080/auth/sign-in' \
--header 'Content-Type: application/json' \
--data-raw '{
    "username":"TestUser123456",
    "password":"123456"
}
```
Пример ответа:
```
{
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MjkzMTEwMjUsImlhdCI6MTcyOTI2NzgyNSwidXNlcl9pZCI6MX0.MxUvy0xQzPqtCUNQR2B7zBE3KS5JZrWd-nNSFExpbVM"
}
```
Сделать депозит

POST   /account/deposit
```
curl --location 'localhost:8080/account/deposit' --header 'Content-Type: application/json' --header 'Authorization: ••••••' --data '{
    "account_id": 1,
    "amount": 100
}'
```
Пример ответа:
```
{
    "message": "deposit succesful"
}
```
Вывод средств

POST   /account/withdraw
```
curl --location 'localhost:8080/account/withdraw' --header 'Content-Type: application/json' --header 'Authorization: ••••••' --data '{
    "account_id": 1,
    "amount": 100
}'
```
Пример ответа:
```
{
    "message": "withdraw succesful"
}
```
Перевод средств

POST   /account/transfer
```
curl --location 'localhost:8080/account/transfer' --header 'Content-Type: application/json' --data '{
    "id_from": 2,
    "id_to": 1,
    "amount": 100
}'
```
Пример ответа:
```
{
    "message": "transfer succesful"
}
```
Баланс пользователя

GET    /operations/user-balance/:currency
```
curl --location 'localhost:8080/operations/user-balance/USD' --header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MjkzMTEwMjUsImlhdCI6MTcyOTI2NzgyNSwidXNlcl9pZCI6MX0.MxUvy0xQzPqtCUNQR2B7zBE3KS5JZrWd-nNSFExpbVM' --data ''
```
Пример ответа:
```
{
    "balance": 36.202299700000005,
    "currency": "USD"
}
```
История транзакций

GET    /operations/transaction-history
```
curl --location --request GET 'localhost:8080/operations/transaction-history' --header 'Content-Type: application/json' --header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MjkzMTEwMjUsImlhdCI6MTcyOTI2NzgyNSwidXNlcl9pZCI6MX0.MxUvy0xQzPqtCUNQR2B7zBE3KS5JZrWd-nNSFExpbVM' --data '{
    "id": 1,
    "sort_type": "amount_high_to_low",
    "limit": 5,
    "offset": 5
}'
```
Пример ответа:
```
{
    "transaction list": [
        {
            "id": 16,
            "account_id": 1,
            "amount": 1000,
            "operation_type": "transfer_to",
            "created": "2024-10-15T22:46:25.47706Z"
        },
        {
            "id": 17,
            "account_id": 1,
            "amount": 1000,
            "operation_type": "transfer_from",
            "created": "2024-10-15T22:46:32.062831Z"
        },
        {
            "id": 22,
            "account_id": 1,
            "amount": 1000,
            "operation_type": "deposit",
            "created": "2024-10-15T22:48:09.510633Z"
        },
        {
            "id": 23,
            "account_id": 1,
            "amount": 1000,
            "operation_type": "deposit",
            "created": "2024-10-15T22:48:10.13647Z"
        },
        {
            "id": 24,
            "account_id": 1,
            "amount": 1000,
            "operation_type": "deposit",
            "created": "2024-10-15T22:48:10.831553Z"
        }
    ]
}
```
