# Run hot reload for coding

1. Install reflex 
    go get github.com/cespare/reflex
2. Run this code inside the folder 
    reflex -r '\.go' -s -- sh -c "go run main.go"
