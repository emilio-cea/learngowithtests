package main

type Transaction struct {
	From string
	To   string
	Sum  int
}

func BalanceFor(transactions []Transaction, name string) int {
	adjustBalance := func(currentBalance int, t Transaction) int {
		if t.From == name {
			return currentBalance - t.Sum
		}
		if t.To == name {
			return currentBalance + t.Sum
		}
		return currentBalance
	}
	return Reduce(transactions, adjustBalance, 0)
}
