// Command demo is the smallest thing hanzo.app can build that is not a
// container image: a static Go binary, published by name+digest.
package main

import "fmt"

// Version is stamped by the build (-ldflags "-X main.Version=…").
var Version = "dev"

func main() { fmt.Printf("hanzo build demo %s\n", Version) }
