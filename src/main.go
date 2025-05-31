
package main

import (
    "fmt"
    "log"
    "nativecoin-os/alchemy"
    "nativecoin-os/nativecoin"
    "os"
)

func main() {
    apiKey := os.Getenv("ALCHEMY_API_KEY")
    if apiKey == "" {
        log.Fatal("Please set ALCHEMY_API_KEY env variable")
    }

    client, err := alchemy.ConnectAlchemy(apiKey)
    if err != nil {
        log.Fatal(err)
    }
    defer client.Close()

    fmt.Println("Starting NativeCoin OS kernel...")

    coin := nativecoin.New()
    coin.Mint("address1", 1000)
    coin.Mint("address2", 500)

    fmt.Printf("Balance address1: %d\n", coin.GetBalance("address1"))
    fmt.Printf("Balance address2: %d\n", coin.GetBalance("address2"))

    success := coin.Transfer("address1", "address2", 300)
    if success {
        fmt.Println("Transfer success")
    } else {
        fmt.Println("Transfer failed")
    }

    fmt.Printf("Balance address1: %d\n", coin.GetBalance("address1"))
    fmt.Printf("Balance address2: %d\n", coin.GetBalance("address2"))
}
