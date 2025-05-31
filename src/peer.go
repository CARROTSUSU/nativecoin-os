package p2p

import (
    "context"
    "fmt"
    "log"

    libp2p "github.com/libp2p/go-libp2p"
    peerstore "github.com/libp2p/go-libp2p-core/peer"
    network "github.com/libp2p/go-libp2p-core/network"
    host "github.com/libp2p/go-libp2p-core/host"
    ma "github.com/multiformats/go-multiaddr"
)

func StartNode() host.Host {
    h, err := libp2p.New()
    if err != nil {
        log.Fatalf("Failed to create host: %v", err)
    }

    fmt.Println("Node ID:", h.ID().Pretty())
    fmt.Println("Listening on:")
    for _, addr := range h.Addrs() {
        fmt.Println(addr.String() + "/p2p/" + h.ID().Pretty())
    }

    h.SetStreamHandler("/nativecoin/1.0.0", func(s network.Stream) {
        fmt.Println("New stream opened")
        buf := make([]byte, 256)
        n, err := s.Read(buf)
        if err == nil {
            fmt.Println("Received:", string(buf[:n]))
        }
        s.Close()
    })

    return h
}

func ConnectToPeer(h host.Host, addr string) {
    maddr, err := ma.NewMultiaddr(addr)
    if err != nil {
        log.Fatal(err)
    }

    info, err := peerstore.AddrInfoFromP2pAddr(maddr)
    if err != nil {
        log.Fatal(err)
    }

    if err := h.Connect(context.Background(), *info); err != nil {
        log.Fatalf("Failed to connect: %v", err)
    }

    fmt.Println("Connected to", info.ID)
    s, err := h.NewStream(context.Background(), info.ID, "/nativecoin/1.0.0")
    if err != nil {
        log.Fatal(err)
    }

    msg := "Hello from NativeCoin OS node!"
    s.Write([]byte(msg))
    s.Close()
}
