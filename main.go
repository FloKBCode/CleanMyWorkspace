package main

import (
	"fmt"

	"github.com/Mentor-Paris/CleanMyWorkspace"
)

func main() {
	workspace := CleanMyWorkspace.GenererateWorkSpace()

	// Afficher le souk
	fmt.Println("=== LE SOUK AVANT NETTOYAGE ===")
	displayWorkspace(workspace)
}

func displayWorkspace(workspace *[][]*string) {
	if workspace == nil {
		fmt.Println("Le souk est vide")
		return
	}

	for _, row := range *workspace {
		fmt.Print("|")
		for _, cell := range row {
			if cell == nil {
				fmt.Print("nil|")
			} else {
				fmt.Print(*cell + "|")
			}
		}
		fmt.Println()
	}
}
