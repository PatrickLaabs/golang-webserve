# Golang Web DMS Proejct
The goal of this project is to build a simply dms based on web with golang.
This will function as a web based view of your files / documents.

## Run hot reload for coding

1. Install reflex 
    go get github.com/cespare/reflex
2. Run this code inside the folder 
    reflex -r '\.go' -s -- sh -c "go run main.go"

## Strcture
- main.go (Main programm)
- README.md
- assets/ (assets for the bootstrap template)
- files/ (file location)
- filestwo/ (file location)
- templates/ (bootstrap template and gohtml files)
