
package main

import (
    "fmt"
    "math/rand"
    "time"
)

type Bank struct {
    Name       string
    BankMoney  float64
    Deposit    float64
    Credit     float64
    Clients    []Client
}

type Client struct {
    Name          string
    Surname       string
    AccountNumber string
    CDeposit      float64
    CCredit       float64
}

func NewBank(name string, money float64) Bank {
    return Bank{Name: name, BankMoney: money}
}

func NewClient(name, surname, accountNumber string) Client {
    return Client{Name: name, Surname: surname, AccountNumber: accountNumber}
}

func (c *Client) StartActivity(bank *Bank, actions chan string) {
    for {
        action := rand.Intn(2)
        amount := float64(rand.Intn(1000))

        if action == 0 { // Deposit
            c.CDeposit += amount
            bank.Deposit += amount
            actions <- fmt.Sprintf("Client %s %s deposited %.2f", c.Name, c.Surname, amount)
        } else { // Credit
            if bank.BankMoney >= amount {
                c.CCredit += amount
                bank.Credit += amount
                bank.BankMoney -= amount
                actions <- fmt.Sprintf("Client %s %s took credit of %.2f", c.Name, c.Surname, amount)
            } else {
                actions <- fmt.Sprintf("Client %s %s couldn't take credit of %.2f (insufficient funds)", c.Name, c.Surname, amount)
            }
        }

        time.Sleep(1 * time.Second)
    }
}

func Menu(bank *Bank, clients *[]Client, actions chan string) {
    for {
        fmt.Println("Menu:")
        fmt.Println("1. Create bank")
        fmt.Println("2. Add client")
        fmt.Println("3. View client by surname")
        fmt.Println("4. View client by account number")
        fmt.Println("5. Exit")

        var choice int
        fmt.Scan(&choice)

        switch choice {
        case 1:
            var name string
            var money float64
            fmt.Print("Enter bank name: ")
            fmt.Scan(&name)
            fmt.Print("Enter initial money: ")
            fmt.Scan(&money)
            *bank = NewBank(name, money)
            fmt.Println("Bank created.")
        case 2:
            var name, surname, accountNumber string
            fmt.Print("Enter client name: ")
            fmt.Scan(&name)
            fmt.Print("Enter client surname: ")
            fmt.Scan(&surname)
            fmt.Print("Enter account number: ")
            fmt.Scan(&accountNumber)
            client := NewClient(name, surname, accountNumber)
            *clients = append(*clients, client)
            go client.StartActivity(bank, actions)
            fmt.Println("Client added.")
        case 3:
            var surname string
            fmt.Print("Enter client surname: ")
            fmt.Scan(&surname)
            for _, c := range *clients {
                if c.Surname == surname {
                    fmt.Printf("Client: %+v\n", c)
                }
            }
        case 4:
            var accountNumber string
            fmt.Print("Enter account number: ")
            fmt.Scan(&accountNumber)
            for _, c := range *clients {
                if c.AccountNumber == accountNumber {
                    fmt.Printf("Client: %+v\n", c)
                }
            }
        case 5:
            close(actions)
            return
        }
    }
}

func main() {
    var bank Bank
    var clients []Client
    actions := make(chan string)

    go func() {
        for action := range actions {
            fmt.Println(action)
        }
    }()

    Menu(&bank, &clients, actions)
}
