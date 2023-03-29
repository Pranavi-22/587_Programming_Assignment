MyRpc interface {
	
	Deposit(user string) (balance float64)
	
	Withdraw(amount float64,user string) (newBalance float64)
	
	DepositBalance(amount1 float64,user string)(updatedepositbalance float64)
	
	Transfer(user1 string,user2 string,amounttobesend float64)(mebalance float64)
}
