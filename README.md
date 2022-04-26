# Development
### Set up
```bash
docker-compose -f docker-compose.development.yml build --no-cache
docker-compose -f docker-compose.development.yml up -d
```

### Directory tree
```
├─ docker/
    └─ api
        └─ environments
            ├─ production/
            ├─ development/
            └─ test/
├─ documents/
    ├─ design/
    ├─ go.mod
    └─ go.sum
├─ server/
    ├─ adapters/
        ├─ controllers/
        ├─ gateways/
        └─ presenter/
    ├─ infrastructure/
        ├─ routes/
        └─ database/
    ├─ usecases/
    ├─ entities/
    ├─ .air.toml
    ├─ go.mod
    ├─ go.sum
    ├─ main.go
    ├─ start.sh
    └─ staticcheck.conf
└─ docker-compose.development.yml
```

## Using Clean Architecture
- adapters directory
    - controllers
        - This directory role is input request
    - gateways
        - This directory role is connect to infrastructures
    - presenter
        - This directory role is format output data
- infrastructures directory
    - routes
        - Set routing for each endpoints
    - database
        - Connect to database
- usecases
    - Access to entities and gateway
- entities
    - Domain logics

# Swagger
### About
- Use goa tool.(https://github.com/goadesign/goa) and swagger ui
### Usage
- When you want to Add routing or update routing on swagger
    - Update `documents/design/design.go`
    - Execute following commands
    ```
    docker-compose -f docker-compose.development.yml exec swagger-editor bash
    goa gen documents/design
    ```
