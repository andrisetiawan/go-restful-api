Go RESTful API
===

A simple sample of RESTful API using Golang.

## Non standard lib dependency

- Gorilla Mux for API routing


## Build & Run

```
go build && ./go-restful-api
```

## API

#### Create Single Item

```
[POST] /items
{
  "id": "1",
  "name": "samlple item name"
}
```

#### Get All Items

```
[GET] /items
```

#### Get Single Item by ID

```
[GET] /items/:id
```

#### Delete Single Item by ID

```
[DELETE] /items/:id
```

