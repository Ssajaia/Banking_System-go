package main

import "fmt"

func localCreate() (string, string, float64, string) {
	var name, lname, ID string
	var initialBalance float64

	fmt.Println("Enter your name: ")
	fmt.Scan(&name)
	fmt.Println("Enter LastName: ")
	fmt.Scan(&lname)
	fmt.Println("Enter ID: ")
	fmt.Scan(&ID)
	fmt.Println("Enter initial balance: ")
	fmt.Scan(&initialBalance)

	fmt.Print("\nAccount Created successfully\n")
	return name, lname, initialBalance, ID
}

func Answer(answer int, acc *Bank) {
	var newBal float64
	switch answer {
	case 1:
		fmt.Println("Enter updated balance:")
		fmt.Scan(&newBal)
		acc.UpdateBalance(newBal)
	case 2:
		fmt.Printf("Current Balance: %.2f\n", acc.GetBalance())
	case 3:
		fmt.Println("Enter Deposit amount:")
		fmt.Scan(&newBal)
		acc.Deposit(newBal)
	case 4:
		fmt.Println("Enter Withdraw amount:")
		fmt.Scan(&newBal)
		acc.Withdraw(newBal)
	case 5:
		name, lname, balance, ID := localCreate()
		newAccount := Bank{name, lname, balance, ID}
		fmt.Printf("New account created for %s %s with ID %s\n", newAccount.Name, newAccount.Lname, newAccount.ID)

		accounts, err := LoadAccounts("accounts.json")
		if err != nil {
			fmt.Println("Error loading accounts:", err)
			return
		}

		
		accounts = append(accounts, newAccount)
		if err := SaveAccounts(accounts, "accounts.json"); err != nil {
			fmt.Println("Error saving accounts:", err)
		}
	default:
		fmt.Println("Invalid option. Please try again.")
	}
}

func choices() {
	fmt.Println("1. Update your balance")
	fmt.Println("2. See your balance")
	fmt.Println("3. Deposit")
	fmt.Println("4. Withdraw")
	fmt.Println("5. Create new Account")
	fmt.Println("0. Exit")
}

func main() {
	
	accounts, err := LoadAccounts("accounts.json")
	if err != nil {
		fmt.Println("Error loading accounts:", err)
		return
	}

	var myAccount *Bank 

	
	if len(accounts) == 0 {
		name, lname, balance, ID := localCreate()
		myAccount = &Bank{name, lname, balance, ID} 
		accounts = append(accounts, *myAccount)    
		if err := SaveAccounts(accounts, "accounts.json"); err != nil {
			fmt.Println("Error saving accounts:", err)
		}
	} else {
		myAccount = &accounts[0] 
		fmt.Printf("Using existing account: %s %s with ID %s\n", myAccount.Name, myAccount.Lname, myAccount.ID)
	}

	for {
		choices()
		ans := 0
		fmt.Scan(&ans)

		if ans == 0 {
			fmt.Println("Exiting the program.")
			break 
		}

		Answer(ans, myAccount)
	}
}