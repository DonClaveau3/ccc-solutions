package main

import "fmt"
import "rsc.io/quote"
import "github.com/spf13/cobra"

func main() {
	fmt.Println("hello world")
	fmt.Println(quote.Go())
}