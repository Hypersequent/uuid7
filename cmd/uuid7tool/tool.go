package main

import (
	"fmt"
	"os"

	"github.com/hypersequent/uuid7"
)

func main() {
	if len(os.Args) < 2 {
		printUsage()
		return
	}

	command := os.Args[1]

	switch command {
	case "new", "generate":
		generateUUID()
	case "help", "-h", "--help":
		printUsage()
	default:
		fmt.Printf("Unknown command: %s\n\n", command)
		printUsage()
		os.Exit(1)
	}
}

func generateUUID() {
	// Generate a new UUID7
	uuid, err := uuid7.UUID7()
	if err != nil {
		fmt.Printf("Error generating UUID: %v\n", err)
		os.Exit(1)
	}

	// Print only canonical HyperUUID7 base58 encoded string
	base58String := uuid7.EncodeBase58(uuid)
	fmt.Printf("%s\n", base58String)
}

func printUsage() {
	fmt.Println("Hypersequent HyperUUID7 Tool")
	fmt.Println("")
	fmt.Println("Usage:")
	fmt.Println("  go run cmd/uuid7tool/tool.go <command>")
	fmt.Println("  # or build and run:")
	fmt.Println("  go build -o uuid7tool cmd/uuid7tool/tool.go")
	fmt.Println("  ./uuid7tool <command>")
	fmt.Println("")
	fmt.Println("Commands:")
	fmt.Println("  new, generate    Generate and print a new HyperUUID7")
	fmt.Println("  help, -h, --help Show this help message")
	fmt.Println("")
	fmt.Println("Examples:")
	fmt.Println("  go run cmd/uuid7tool/tool.go new")
	fmt.Println("  go run cmd/uuid7tool/tool.go generate")
}
