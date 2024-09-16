# Используем официальный образ Go для сборки приложения
FROM golang:1.22-alpine AS builder

# Устанавливаем рабочую директорию внутри контейнера
WORKDIR /app

# Устанавливаем необходимые пакеты
RUN apk add --no-cache gcc musl-dev

# Копируем go.mod и go.sum файлы и загружаем зависимости
COPY go.mod go.sum ./
RUN go mod download

# Копируем исходный код приложения в контейнер
COPY . .

# Сборка бинарного файла приложения
RUN go build -o classifier_service main.go

# Используем минимальный alpine-образ для запуска приложения
FROM alpine:latest

# Устанавливаем необходимые библиотеки
RUN apk add --no-cache libc6-compat

# Копируем собранное приложение из предыдущего контейнера
COPY --from=builder /app/classifier_service /classifier_service


# Запускаем приложение
CMD ["/classifier_service"]
