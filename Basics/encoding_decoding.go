package main

import (
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
)

func main() {
	base64Example()
	jsonExample()
	hexExample()
}

func base64Example() {
	// Base64 Encoding
	data := "Hello, World!"
	encoded := base64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println("Encoded:", encoded)

	// Base64 Decoding
	decoded, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		fmt.Println("Error decoding:", err)
		return
	}
	fmt.Println("Decoded:", string(decoded))
}

// go run Basics/encoding_decoding.go
// Output:
// Encoded: SGVsbG8sIFdvcmxkIQ==
// Decoded: Hello, World!

// JSON Encoding and Decoding
func jsonExample() {
	// Import "encoding/json" package
	// Define a struct for JSON encoding
	type Person struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	// Create an instance of Person
	p := Person{Name: "Alice", Age: 30}

	// JSON Encoding
	jsonData, err := json.Marshal(p)
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		return
	}
	fmt.Println("JSON Encoded:", string(jsonData))

	// JSON Decoding
	var pDecoded Person
	err = json.Unmarshal(jsonData, &pDecoded)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}
	fmt.Println("JSON Decoded:", pDecoded)
}

// Output:
// JSON Encoded: {"name":"Alice","age":30}
// JSON Decoded: {Alice 30}

// Hex Encoding and Decoding

func hexExample() {
	// Import "encoding/hex" package
	// Hex Encoding
	data := "Hello, World!"
	encoded := make([]byte, hex.EncodedLen(len(data)))
	hex.Encode(encoded, []byte(data))
	fmt.Println("Hex Encoded:", string(encoded))

	// Hex Decoding
	decoded := make([]byte, hex.DecodedLen(len(encoded)))
	n, err := hex.Decode(decoded, encoded)
	if err != nil {
		fmt.Println("Error decoding hex:", err)
		return
	}
	fmt.Println("Hex Decoded:", string(decoded[:n]))
}

// Output:
// Hex Encoded: 48656c6c6f2c20576f726c6421
// Hex Decoded: Hello, World!
