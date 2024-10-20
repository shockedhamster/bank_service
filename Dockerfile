FROM golang:1.23.1

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем все файлы проекта в контейнер
COPY ./ ./ 

# Устанавливаем утилиту wait-for-it
COPY wait-for-it.sh /usr/local/bin/wait-for-it

# Делаем wait-for-it.sh исполняемым
RUN chmod +x /usr/local/bin/wait-for-it
RUN chmod -R 777 ./.database/data

# Устанавливаем psql
RUN apt-get update && apt-get -y install postgresql-client

# Строим Go приложение
RUN go mod download
RUN go build -o bank_service ./cmd/main.go

# Устанавливаем net-tools
RUN apt-get update && apt-get install -y net-tools

# Открываем порт 8080
EXPOSE 8080

# Команда для запуска приложения с wait-for-it
CMD ["/usr/local/bin/wait-for-it", "db:5432", "--", "./bank_service"]