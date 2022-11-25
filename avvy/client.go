package avvy


import (
    "binary"
   "fmt"
   "strings"
    // "github.com/ethereum/go-ethereum/ethclient"
)

type Client struct {
    Name string
}

/* Executes a single iteration of hashing */
func nameHashIteration(prevHash uint64, label string) uint64 {
    var l1 [31]byte
    var l2 [31]byte
    var n1 uint64
    var n2 uint64
    for i, char := range label {
        if i < 31 {
            l1[i] = byte(char)
        } else {
            l2[i] = byte(char)
        }
    }
    buf := bytes.NewReader(l1)
    err := binary.Read(buf, binary.LittleEndian, &n1)
    buf := bytes.NewReader(l2)
    err := binary.Read(buf, binary.LittleEndian, &n2)
    fmt.Println(n1)
    fmt.Println(n2)
    return prevHash
}

/* Hashes the provided name */
func (c *Client) NameHash(domain string) uint64 {
    labels := strings.Split(domain, ".")

    // reverse labels
    for i, j := 0, len(labels)-1; i < j; i, j = i+1, j-1 {
        labels[i], labels[j] = labels[j], labels[i]
    }

    var hash uint64 = 0
    for _, label := range labels {
        hash = nameHashIteration(hash, label)
    }

    return hash
}

