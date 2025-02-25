package main

import (
	"Qubecinema/challenge/distributer"
	"Qubecinema/challenge/utils"
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
)

func init() {
	err := utils.LoadCSV("cities.csv")
	if err != nil {
		fmt.Printf("Error loading csv %v", err)
		return
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	color.HiGreen("...............Distributor System...............")

	for {
		color.HiCyan("\n1. Add Distributor")
		color.HiCyan("2. Check Permission")
		color.HiCyan("3. Exit")
		color.HiYellow("Enter your choice: ")

		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("❌ Error reading input:", err)
			continue
		}

		input = strings.TrimSpace(input)

		switch input {
		case "1":
			fmt.Print("Enter Distributor Name: ")
			name, _ := reader.ReadString('\n')
			name = strings.TrimSpace(name)
			fmt.Print("Enter Parent Distributor (leave empty if none): ")
			parent, _ := reader.ReadString('\n')
			parent = strings.TrimSpace(parent)

			fmt.Print("Enter regions to INCLUDE (comma-separated, e.g., IN,US, KA-IN): ")
			includeInput, _ := reader.ReadString('\n')
			include := strings.Split(strings.TrimSpace(includeInput), ",")

			fmt.Print("Enter regions to EXCLUDE (comma-separated, e.g., KA-IN,CENAI-TN-INDIA): ")
			excludeInput, _ := reader.ReadString('\n')
			exclude := strings.Split(strings.TrimSpace(excludeInput), ",")
			err := distributer.AddDistributor(name, parent, include, exclude)
			if err != nil {
				color.HiRed("Ditributer Not Added %v", err)
			}
		case "2":
			fmt.Print("Enter Distributor Name: ")
			name, _ := reader.ReadString('\n')
			name = strings.TrimSpace(name)
			fmt.Print("Enter regions to CHECK (KA-IN): ")
			includeInput, _ := reader.ReadString('\n')
			includeInput = strings.TrimSpace(includeInput)
			perm, err := distributer.CheckPermission(name, includeInput)
			if err != nil {
				color.HiRed("Distributer does not exist")
			} else {
				color.HiGreen("IS ALLOWED: %t", perm)
			}
		case "3":
			color.HiRed("Exiting...")
			return

		default:
			color.HiRed("❌ Invalid choice, please try again.")
		}
	}
}
