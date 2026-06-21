# Enterprise Design Patterns in Go

Small demo for a Medium article about enterprise design patterns from Martin Fowler's *Patterns of Enterprise Application Architecture*.

## What it shows

- `Service Layer` coordinates the use case
- `Repository` hides storage details
- `Data Mapper` converts between domain objects and records
- `Unit of Work` batches the write

## Run

```bash
go test ./...
go run ./main.go
```
