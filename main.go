package main

import (
	"fmt"

	"github.com/Mentor-Paris/CleanMyWorkspace"

	"CleanMyWorkspace/Clean"
)

func main() {
	workspace := CleanMyWorkspace.GenererateWorkSpace()

	// Afficher le souk
	fmt.Println("=== LE SOUK AVANT NETTOYAGE ===")
	displayWorkspace(workspace)

	fmt.Println("\n🧹 Nettoyage en cours...\n")

	cleanedWorkspace := Clean.CleanWorkSpace(workspace)

	fmt.Println("=== LE SOUK APRES NETTOYAGE ===")
	displayWorkspace(cleanedWorkspace)

	fmt.Println("\n✨ Bravo ! Notre espace de travail est maintenant propre ! ✨")
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
