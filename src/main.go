package main

import (
    "fmt"
    "os"
    "nativecoin-os/p2p"
)

func main() {
    fmt.Println("Booting NativeCoin OS with P2P...")

    h := p2p.StartNode()

    if len(os.Args) > 1 {
        // Connect to another peer via cmd argument
        peerAddr := os.Args[1]
        p2p.ConnectToPeer(h, peerAddr)
    }

    select {} // Keep node running
}
