package core

type TransferInterfae interface {
	MakeTransfer(Account, Account) (Account, Transaction, error)
	MakeRefund(Transaction) (Transaction, error)
}

type UserInterface interface {
	CreateUser(string) (User, error)
}

type AccountInterface interface {
	CreateAccount(string) (Account, error)
}
