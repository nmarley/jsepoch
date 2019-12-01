package main

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

var jsepochCmd = &cobra.Command{
	Use:   "jsepoch",
	Short: "",
	Long:  "",
	Run:   printEpochMilliSecondPrecision,
}

func printEpochMilliSecondPrecision(c *cobra.Command, inp []string) {
    t := time.Now().UTC()

	fmt.Println(t.UnixNano() / 1000000)
}

func main() {
	jsepochCmd.Execute()
}
