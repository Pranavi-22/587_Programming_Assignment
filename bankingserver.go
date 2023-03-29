package main

import (
	"ethos/altEthos"
	"ethos/myRpc"
	"ethos/syscall"
	"log"
)

var userbalance float64 = 1000

var map2 = map[string]float64{
	"me":      500,
	"user1":   1000,
	"bennett": 3000,
}

func init() {
	myRpc.SetupMyRpcDeposit(deposit)
	myRpc.SetupMyRpcWithdraw(withdraw)
	myRpc.SetupMyRpcDepositBalance(updatedBalance)
	myRpc.SetupMyRpcTransfer(transfer)
}

// balance amount will be displayed here
func deposit(user string) myRpc.MyRpcProcedure {
	log.Println("called balance", map2[user])

	return &myRpc.MyRpcDepositReply{map2[user]}
}

// withdraw will done from the account the amount given by the user will be subracted
func withdraw(amount float64, user string) myRpc.MyRpcProcedure {

	if map2[user] < amount {
		log.Println("withdraw cannot be done")
	} else {
		map2[user] = map2[user] - amount

	}

	log.Println("withdraw amount from the user given is", map2[user])

	return &myRpc.MyRpcWithdrawReply{map2[user]}
}

// the balance will be updated here that is the amount given will be added
func updatedBalance(amount1 float64, user string) myRpc.MyRpcProcedure {
	map2[user] = map2[user] + amount1
	log.Println("update balance after adding amount", map2[user])

	return &myRpc.MyRpcDepositBalanceReply{map2[user]}
}

// transfer will be done here from one one user account to another user account
func transfer(user1 string, user2 string, amount float64) myRpc.MyRpcProcedure {

	if map2[user1] > amount {
		map2[user1] = map2[user1] - amount
		map2[user2] = map2[user2] + amount
	}

	log.Println("withdraw from me", map2[user1])
	log.Println("withdraw to user given", map2[user2])

	return &myRpc.MyRpcTransferReply{map2[user2]}
}

func main() {

	altEthos.LogToDirectory("test/bankingserver")

	log.Printf("Entering main")

	listeningFd, status := altEthos.Advertise("myRpc")
	if status != syscall.StatusOk {
		log.Println("Advertising service failed: ", status)
		altEthos.Exit(status)
	}

	for {
		_, fd, status := altEthos.Import(listeningFd)
		if status != syscall.StatusOk {
			log.Printf("Error calling Import: %v\n", status)
			altEthos.Exit(status)
		}

		log.Println("new connection accepted")

		t := myRpc.MyRpc{}
		altEthos.Handle(fd, &t)
	}

}
