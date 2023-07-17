// changes
// clear the whole screen each time a command is executed
// add the ability to run each command again
package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

func main() {
	loadModule()
	reader := bufio.NewReader(os.Stdin)
	for {
		displayMenu()
		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		switch choice {
		case "1":
			addMembers(reader)
		case "2":
			removeMembers(reader)
		case "3":
			bulkRequests(reader)
		case "q":
			os.Exit(0) //exit the program when the user selects q
		default:
			fmt.Println("Invalid choice, please choose from the listed options")
		}
	}
}

// function to display the main menu
func displayMenu() {
	clearScreen()

	fmt.Println()
	fmt.Println()
	fmt.Println("       Service Desk Automator       ")
	fmt.Println()
	fmt.Println("Please select an option:")
	fmt.Println()
	fmt.Println("1. Add Members to Groups")
	fmt.Println("2. Remove Members from Groups")
	fmt.Println("3. Bulk Requests")
	fmt.Println()
	fmt.Println("q. Quit")
	fmt.Println()
	fmt.Print("Select an Option: ")
}

func clearScreen() {
	switch runtime.GOOS {
	case "windows":
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	default:
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

// function to display the sub menu for option 1 Add Members in the main menu
func addMembers(reader *bufio.Reader) {
	for {
		fmt.Println("You selected Add Members to Groups")

		employeeID, groupName := readInput(reader)

		cmd := exec.Command("powershell", "-Command", "Import-Module", "./ServiceDeskAutomator.psm1;", fmt.Sprintf("Add-GroupMemberships %s %s", employeeID, groupName))

		output, err := cmd.Output()
		if err != nil {
			fmt.Printf("Command execution failed: %v\n", err)
			if exitErr, ok := err.(*exec.ExitError); ok {
				fmt.Println(string(exitErr.Stderr))
			}
		} else {
			fmt.Println(string(output))
		}

		fmt.Print("Run again? (y/n): ")
		choice, _ := reader.ReadString('\n')
		choice = strings.ToLower(strings.TrimSpace(choice))
		if choice != "y" {
			clearScreen()
			break
		}
	}
}

// function to display the sub menu for option 2 Remove Members in the main menu
func removeMembers(reader *bufio.Reader) {
	for {
		fmt.Println("You selected Remove Members from Groups")

		employeeID, groupName := readInput(reader)

		cmd := exec.Command("powershell", "-Command", "Import-Module", "./ServiceDeskAutomator.psm1;", fmt.Sprintf("Remove-GroupMemberships %s %s", employeeID, groupName))

		output, err := cmd.Output()
		if err != nil {
			fmt.Printf("Command execution failed: %v\n", err)
			if exitErr, ok := err.(*exec.ExitError); ok {
				fmt.Println(string(exitErr.Stderr))
			}
		} else {
			fmt.Println(string(output))
		}

		fmt.Print("Run again? (y/n): ")
		choice, _ := reader.ReadString('\n')
		choice = strings.ToLower(strings.TrimSpace(choice))
		if choice != "y" {
			clearScreen()
			break
		}
	}
}

// function to display the sub menu for option 3 Bulk Requests in the main menu
func bulkRequests(reader *bufio.Reader) {
	for {
		fmt.Println("You selected Bulk Requests")

		CSVPath := readFilePath(reader)

		cmd := exec.Command("powershell", "-Command", "Import-Module", "./ServiceDeskAutomator.psm1;", fmt.Sprintf("Invoke-ModifyGroupsFromCsv %s", CSVPath))

		output, err := cmd.Output()
		if err != nil {
			fmt.Printf("Command execution failed: %v\n", err)
			if exitErr, ok := err.(*exec.ExitError); ok {
				fmt.Println(string(exitErr.Stderr))
			}
		} else {
			fmt.Println(string(output))
		}

		fmt.Print("Run again? (y/n): ")
		choice, _ := reader.ReadString('\n')
		choice = strings.ToLower(strings.TrimSpace(choice))
		if choice != "y" {
			clearScreen()
			break
		}
	}
}

func readInput(reader *bufio.Reader) (string, string) {
	fmt.Print("Enter the employee ID: ")
	employeeID, _ := reader.ReadString('\n')
	employeeID = strings.TrimSpace(employeeID)

	fmt.Print("Enter the group name: ")
	groupName, _ := reader.ReadString('\n')
	groupName = strings.TrimSpace(groupName)

	return employeeID, groupName
}

func readFilePath(reader *bufio.Reader) string {
	fmt.Print("Enter the path to your CSV: ")
	CSVPath, _ := reader.ReadString('\n')
	CSVPath = strings.TrimSpace(CSVPath)

	return CSVPath
}

func loadModule() error {
	currentDir, err := os.Getwd()
	if err != nil {
		return err
	}

	modulePath := filepath.Join(currentDir, "ServiceDeskAutomator.psm1")

	cmd := exec.Command("powershell", "-Command", "Import-Module", modulePath)
	err = cmd.Run()
	if err != nil {
		return fmt.Errorf("failed to load module: %v", err)
	}

	return nil
}
