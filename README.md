# test-netzme
Technical Test Netzme
- HOST : http://localhost:3000
- SWAGGER : http://localhost:3000/swagger/index.html

## Pre Requisite
- Go version 1.20

## How To Run
``` bash

# setup swagger 
$ go install github.com/swaggo/swag/cmd/swag@latest
$ swag init --propertyStrategy snakecase

# mod
$ go mod tidy

# run 
$ go run main.go
```

## Architecture 
This project built in clean architecture that contains some layer :
1. Driver   
2. Factory 
3. Delivery
4. Repository
5. Usecase
6. Model

# Author
mcholismalik.official@gmail.com
