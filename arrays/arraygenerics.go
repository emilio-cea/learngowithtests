package main

type Transaction struct {
	From string
	To   string
	Sum  int
}

func BalanceFor(transactions []Transaction, name string) int {
	var balance int
	for _, transaction := range transactions {
		if transaction.From == name {
			balance -= transaction.Sum
		}
		if transaction.To == name {
			balance += transaction.Sum
		}
	}
	return balance
}
