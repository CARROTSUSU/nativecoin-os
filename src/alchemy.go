package alchemy

import (
    "context"
    "fmt"
    "github.com/ethereum/go-ethereum/ethclient"
)

func ConnectAlchemy(apiKey string) (*ethclient.Client, error) {
    url := "https://eth-goerli.alchemyapi.io/v2/" + apiKey
    client, err := ethclient.DialContext(context.Background(), url)
    if err != nil {
        return nil, err
    }
    fmt.Println("Connected to Alchemy Ethereum node")
    return client, nil
}
