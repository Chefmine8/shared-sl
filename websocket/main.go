package main

import (
	"fmt"
	"os"
	"strconv"

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
			fmt.Printf("Serveur démarré sur le port %d (Fly: %t, Accident: %t, Little: %t)\n", port, fly, accident, little)
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
		RunE: func(cmd *cobra.Command, args []string) error {
			addr := args[0]
			order, err := strconv.Atoi(args[1])
			if err != nil {
				return fmt.Errorf("order parm gotta be a number : %w", err)
			}

			fmt.Printf("Connexion à %s en position #%d\n", addr, order)
			return nil
		},
	}

	rootCmd.AddCommand(serverCmd, joinCmd)
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
