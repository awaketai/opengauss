# GORM OpenGauss Driver

> just adaptation opengauss 3.0.0

## Quick Start

```go
import (
  "github.com/awaketai/opengauss"
  "gorm.io/gorm"
)

// https://github.com/jackc/pgx
dsn := "host=localhost user=gaussdb dbname=postgres sslmode=disable password=Enmo@123 port=5432"
db, err := gorm.Open(opengauss.Open(dsn), &gorm.Config{})
```

## Configuration

```go
import (
  "github.com/awaketai/opengauss"
  "gorm.io/gorm"
)

db, err := gorm.Open(opengauss.New(postgres.Config{
  DSN: "host=localhost user=gaussdb dbname=postgres sslmode=disable password=Enmo@123 port=5432 TimeZone=Asia/Shanghai", // data source name, refer https://github.com/jackc/pgx
  PreferSimpleProtocol: true, // disables implicit prepared statement usage. By default pgx automatically uses the extended protocol
}), &gorm.Config{})
```


Checkout [https://gorm.io](https://gorm.io) for details.
