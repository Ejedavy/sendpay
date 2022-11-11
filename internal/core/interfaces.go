package core

type UserInterface interface {
	CreateUser(string) (User, error)
}

type AccountInterface interface {
	CreateAccount(string) (Account, error)
	TopUp(Account, float32) (Account, error)
	MakeTransfer(Account, Account) (Account, Transaction, error)
	MakeRefund(Transaction) (Transaction, error)
}
