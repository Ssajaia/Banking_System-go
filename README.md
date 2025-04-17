# Go Bank Account Management

A simple command-line application written in Go for managing bank accounts. This application allows users to create accounts, update balances, deposit and withdraw funds, and save account information to a JSON file.

## Features

- Create a new bank account with an initial balance.
- Update the account balance.
- Deposit money into the account.
- Withdraw money from the account (with insufficient funds check).
- View the current balance.
- Save and load account information from a JSON file.

## Requirements

- Go 1.16 or later

## Installation

1. Clone the repository:

   ```bash
   [git clone https://github.com/yourusername/go-bank-account-management.git](https://github.com/Ssajaia/Banking_System-go.git)
   cd Banking_System-go
Build the application:
bash

Copy Code
go build -o bank
Run the application:
bash

Copy Code
./bank
Usage
When you run the application, you will be prompted to create a new account if no accounts exist. Enter your name, last name, ID, and initial balance.
After creating an account, you will see a menu with the following options:
Update your balance
See your balance
Deposit
Withdraw
Create a new account
Exit
Follow the prompts to perform the desired actions.
JSON File
The application saves all account information in a file named accounts.json. This file is created in the same directory as the application. If the file does not exist, it will be created automatically when you create a new account.

Contributing
Contributions are welcome! If you have suggestions for improvements or new features, feel free to open an issue or submit a pull request.

License
This project is licensed under the MIT License - see the LICENSE file for details.

Acknowledgments
Go programming language
JSON for data storage
