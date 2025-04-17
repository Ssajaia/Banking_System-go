package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Bank struct {
	Name    string  `json:"name"`
	Lname   string  `json:"lname"`
	Balance float64 `json:"balance"`
	ID      string  `json:"id"`
}

func (b *Bank) UpdateBalance(newBalance float64) {
	b.Balance = newBalance
}

func (b *Bank) GetBalance() float64 {
	return b.Balance
}

func (b *Bank) Deposit(amount float64) {
	b.Balance += amount
}

func (b *Bank) Withdraw(amount float64) {
	if amount > b.Balance {
		fmt.Println("Insufficient funds for withdrawal.")
		return
	}
	b.Balance -= amount
}

func (b Bank) String() string {
	return fmt.Sprintf("[Name: %s, Last name: %s, Balance: %.2f, ID: %s]", b.Name, b.Lname, b.Balance, b.ID)
}

// Save accounts to a JSON file
func SaveAccounts(accounts []Bank, filename string) error {
	data, err := json.MarshalIndent(accounts, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(filename, data, 0644)
}

// Load accounts from a JSON file
func LoadAccounts(filename string) ([]Bank, error) {
	var accounts []Bank
	file, err := os.Open(filename)
	if err != nil {
		if os.IsNotExist(err) {
			return accounts, nil // Return empty slice if file does not exist
		}
		return nil, err
	}
	defer file.Close()

	if err := json.NewDecoder(file).Decode(&accounts); err != nil {
		return nil, err
	}
	return accounts, nil
}
