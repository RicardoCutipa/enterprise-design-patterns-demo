# Enterprise Design Patterns in Go

This repository is a companion demo for a Medium article about enterprise design patterns inspired by Martin Fowler's *Patterns of Enterprise Application Architecture*.

The goal is simple: show how a small business workflow can stay readable even when the code is split into layers with distinct responsibilities.

## What the demo represents

Imagine an order placement flow in a real business system:

1. A customer submits an order.
2. The domain calculates the total.
3. The service validates the request and coordinates the workflow.
4. The unit of work batches the save operation.
5. The repository persists the order through a data mapper.

That is exactly the kind of separation enterprise design patterns are meant to support.

## Patterns in the code

- `Service Layer`: `orders.Service` coordinates the use case.
- `Repository`: `orders.Repository` hides how orders are stored.
- `Data Mapper`: `orders.Mapper` converts between `Order` and `Record`.
- `Unit of Work`: `orders.UnitOfWork` groups write operations before commit.

## Project structure

- `main.go`: runnable demo that prints the workflow
- `orders/order.go`: domain model and total calculation
- `orders/service.go`: application logic
- `orders/repository.go`: storage abstraction and in-memory implementation
- `orders/mapper.go`: translation between domain objects and records
- `orders/unit_of_work.go`: commit boundary for writes
- `orders/service_test.go`: tests for the business flow

## Expected output

When you run the demo, you should see the order flow explained step by step, followed by the stored result.

## Run locally

```bash
go test ./...
go run ./main.go
```

## Why this demo exists

The code is intentionally small, but the structure is the important part. In a larger system, this same shape lets you replace the storage layer, add more validation, or introduce a real database without rewriting the business rules.
