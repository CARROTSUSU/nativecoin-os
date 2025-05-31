package nativecoin

import "fmt"

type NativeCoin struct {
    balances map[string]uint64
}

func New() *NativeCoin {
    return &NativeCoin{balances: make(map[string]uint64)}
}

func (nc *NativeCoin) GetBalance(addr string) uint64 {
    return nc.balances[addr]
}

func (nc *NativeCoin) Mint(addr string, amount uint64) {
    nc.balances[addr] += amount
    fmt.Printf("Minted %d coins to %s\n", amount, addr)
}

func (nc *NativeCoin) Transfer(from, to string, amount uint64) bool {
    if nc.balances[from] < amount {
        return false
    }
    nc.balances[from] -= amount
    nc.balances[to] += amount
    return true
}

