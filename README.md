# entp (enhanced NTP)

Библиотека и утилита для получения точного времени через NTP (Network Time Protocol).

Основана на [beevik/ntp](https://github.com/beevik/ntp), добавляет удобный клиент с поддержкой списка серверов, таймаутов и контекста.

---

## Установка

Модуль импортируется напрямую:

```bash
go get github.com/DenisBochko/entp
```

### В коде:

```go
import "github.com/DenisBochko/entp"
```

## Использование в коде

### Простейший пример:

```go
package main

import (
	"context"
	"fmt"
	"log"

	"github.com/DenisBochko/entp"
)

func main() {
	client := entp.NewClient()

	t, err := client.Now(context.Background())
	if err != nil {
		log.Fatalf("failed to get NTP time: %v", err)
	}

	fmt.Println("Exact time:", t)
}
```

### С опциями:

```go
client := entp.NewClient(
    entp.WithTimeout(2 * time.Second),
    entp.WithAddServers("time.cloudflare.com"),
    entp.WithReplaceDefaultServers("xk.pool.ntp.org")
)
```

## Запуск линтера

Проект использует [golangci-lint](https://golangci-lint.run)

### Установка:

```bash
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
```

### Запуск:

```bash
golangci-lint run ./... --config=./.golangci-lint.yml
```

## Запуск тестов

### Обычные тесты:

```bash
go test ./...
```

### С выводом покрытия:

```bash
go test -cover ./...
```

**P.S. Чтобы тесты правильно работали, необходимо выключить vpn**

## CLI пример:

```bash
go run ./cmd/main.go
```
