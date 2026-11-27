package main

import (
	"fmt"
	"os"
	"sl-net/client"
	"sl-net/server"

	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "sl-net",
		Short: "Routes an ASCII stream train across multiple networked terminals",
	}

	var port int
	var fly bool
	var accident bool
	var little bool

	serverCmd := &cobra.Command{
		Use:   "server",
		Short: "Run the server",
		Run: func(cmd *cobra.Command, args []string) {
			server.StartWebSocket(port)
		},
	}

	serverCmd.Flags().IntVarP(&port, "port", "p", 42010, "Port to listen on")
	serverCmd.Flags().BoolVarP(&fly, "fly", "f", false, "Flying train")
	serverCmd.Flags().BoolVarP(&accident, "accident", "a", false, "Accidented train")
	serverCmd.Flags().BoolVarP(&little, "little", "l", false, "Little train")

	joinCmd := &cobra.Command{
		Use:   "join <address:port> <order>",
		Short: "Connect a client to the train chain",
		Args:  cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			addr := args[0]
			order := args[1]

			fmt.Printf("Connexion à %s en position #%s\n", addr, order)
			client.StartWebSocket(addr, order)
		},
	}

	rootCmd.AddCommand(serverCmd, joinCmd)
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
