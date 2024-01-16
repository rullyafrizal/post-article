# Blog Post App
## How To Run Locally

1. Make sure you already setup mysql server in your local
2. Create database "blogs" in your local mysql
3. Copy config.env.example to config.env
    ```
    cp config/config.env.example config/config.env
    ```
4. Fill env file with your database configuration, do not change APP_PORT
5. Install golang migrate, you can read the steps in here https://github.com/golang-migrate/migrate
6. Run DB Migration, please replace all value with curly brackets with your DB config
    ```
    migrate -database "mysql://{username}:{password}@tcp({host}:{port})/{db_name}" -path infra/database/migration up
    ```
7. Run application with command `go run cmd/main.go`
8. Open `index.html` with your browser, make sure you have internet connection
9. Ping me if you have any trouble when running locally using above steps

## Tech Stacks
- Golang 1.20
- Go Fiber
- MySQL 8
- Viper Env
- GORM