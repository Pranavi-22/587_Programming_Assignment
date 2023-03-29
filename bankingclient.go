package main

import (
	"ethos/altEthos"
	"ethos/syscall"
	"ethos/myRpc"
	"log"
)

//withdraw amount
var withdrawamount float64=64
//amount to be deposited
var updatedBalance float64=30

func init() {
	myRpc.SetupMyRpcDepositReply(depositReply)
	myRpc.SetupMyRpcWithdrawReply(withdrawReply)
	myRpc.SetupMyRpcDepositBalanceReply(depositBalanceReply)
	myRpc.SetupMyRpcTransferReply(transfer)
}

//balance amount 
func depositReply(balance float64) (myRpc.MyRpcProcedure) {
	
	log.Printf("Balance amount is %f\n",balance)
	return nil
}

// withdraw amount 
func withdrawReply(newBalance float64)(myRpc.MyRpcProcedure){
	log.Printf("withdraw amount %f\n",newBalance)
	return nil

}
// deposit amount 
func depositBalanceReply(updatedBalance float64)(myRpc.MyRpcProcedure){
	log.Printf("updatedBalance is %f\n",updatedBalance)
	return nil
}

//transfer amount

func transfer(mebalance float64)(myRpc.MyRpcProcedure){
	log.Printf("person balance after adding amount%f\n",mebalance)
	return nil
}

func main () {

	altEthos.LogToDirectory("test/bankingclient")
	var name=altEthos.GetUser()
	log.Println("the user name is",name)
	var name2="user1"
    
	fd, status1 := altEthos.IpcRepeat("myRpc", "", nil)
	if status1 != syscall.StatusOk {
		log.Printf("Ipc failed: %v\n", status1)
		altEthos.Exit(status1)
	}

	call1 := myRpc.MyRpcWithdraw{withdrawamount,name}
	status1 = altEthos.ClientCall(fd, &call1)
	if status1 != syscall.StatusOk {
		log.Printf("clientCall failed: %v\n", status1)
		altEthos.Exit(status1)
	}

	fd, status2 := altEthos.IpcRepeat("myRpc", "", nil)
	if status2 != syscall.StatusOk {
		log.Printf("Ipc failed: %v\n", status2)
		altEthos.Exit(status2)
	}

	call2 := myRpc.MyRpcDepositBalance{updatedBalance,name}
	status2 = altEthos.ClientCall(fd, &call2)
	if status2 != syscall.StatusOk {
		log.Printf("clientCall failed: %v\n", status2)
		altEthos.Exit(status2)
	}

	fd, status3 := altEthos.IpcRepeat("myRpc", "", nil)
	if status3 != syscall.StatusOk {
		log.Printf("Ipc failed: %v\n", status3)
		altEthos.Exit(status3)
	}

	call3 := myRpc.MyRpcTransfer{name,name2,50}
	status3 = altEthos.ClientCall(fd, &call3)
	if status3 != syscall.StatusOk {
		log.Printf("clientCall failed: %v\n", status3)
		altEthos.Exit(status3)
	}

	fd, status := altEthos.IpcRepeat("myRpc", "", nil)
	if status != syscall.StatusOk {
		log.Printf("Ipc failed: %v\n", status)
		altEthos.Exit(status)
	}

	call := myRpc.MyRpcDeposit{name}
	status = altEthos.ClientCall(fd, &call)
	if status != syscall.StatusOk {
		log.Printf("clientCall failed: %v\n", status)
		altEthos.Exit(status)
	}
	log.Println("syncClient: done")
}