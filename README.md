### 587 HomeWork 2 (Ethos Client Server Banking)



#### Project Description:

This project consists of two programs one is banking server and the other is banking client.We have four operations in this project that is balance,transfer between users,withdraw and deposit.

All this operations are written as interfaces in MyRpc.t files.
Bankingserver.go file consists of the methods implemented according to information given by the client. It intializes all the set of accounts and processes the request.
Bankingclient file consists of methods to send the information about the amount to be withdraw and deposited.
We have two users one is the bydefault user that is "me" and the other user taken is "bennett".

#### Exceution

1st Terminal

```bash
  >> make && make clean && make install
  >> cd client
  >> sudo -E ethosRun
```
2nd Terminal for one user

```bash
  >> cd client
  >> etAl client.ethos 
  "After this step go to other terminal to execute another user and login by bennett then execute the below step"
  >> bankingclient
```

3rd Terminal for second user with name bennett(2nd instance)

```bash
  >> login:bennett
  >> pswd:ethos
  >> cd ..
  >> cd me
  >> cd bank_587_project
  >> cd client
  >> et client.ethos
  >> bankingclient

```


4th Terminal 
```bash
  >> cd client
  >> ethosLog .
```


#### Log File


Log_file1.txt consists the entire log that is generated.

This is the process for me user.

- Intially he as 500 dollars firstly he is withdrawing 64 and depositing 30 and transferring of 50 dollars to another account.
```bash


Client Side 
 >> 1679988769.146362310 [bankingclient.O] the user name is me
 >> 1679988769.175570932 [bankingclient.O] withdraw amount 436.000000
 >> 1679988769.185397912 [bankingclient.O] updatedBalance is 466.000000
 >> 1679988769.199361071 [bankingclient.O] person balance after adding amount1050.000000
 >> 1679988769.205181206 [bankingclient.O] Balance amount is 416.000000
 >> 1679988769.210087102 [bankingclient.O] syncClient: done

Server side

 >> 1679988684.593763986 [bankingserver.O] Entering main
 >> 1679988769.169438497 [bankingserver.O] new connection accepted
 >> 1679988769.172212704 [bankingserver.O] withdraw amount from the user given is 436
 >> 1679988769.178302288 [bankingserver.O] new connection accepted
 >> 1679988769.181788565 [bankingserver.O] update balance after adding amount 466
 >> 1679988769.188838893 [bankingserver.O] new connection accepted
 >> 1679988769.191269388 [bankingserver.O] transfer from me 416
 >> 1679988769.192787257 [bankingserver.O] transfer to user given 1050
 >> 1679988769.200769994 [bankingserver.O] new connection accepted
 >> 1679988769.202165069 [bankingserver.O] called balance 416
```

This is the process for bennett user.

- Intially he as 3000 dollars firstly he is withdrawing 64 and depositing 30 and transferring of 50 dollars to another account.

```bash


Client Side 
>> 1679988788.411148015 [bankingclient.O] the user name is bennett
>> 1679988788.429362185 [bankingclient.O] withdraw amount 2936.000000
>> 1679988788.438044673 [bankingclient.O] updatedBalance is 2966.000000
>> 1679988788.450294982 [bankingclient.O] person balance after adding amount1100.000000
>> 1679988788.456258683 [bankingclient.O] Balance amount is 2916.000000
>> 1679988788.461425642 [bankingclient.O] syncClient: done

Server side
>> 1679988788.413545762 [bankingserver.O] new connection accepted
>> 1679988788.426112764 [bankingserver.O] withdraw amount from the user given is 2936
>> 1679988788.433065792 [bankingserver.O] new connection accepted
>> 1679988788.435405579 [bankingserver.O] update balance after adding amount 2966
>> 1679988788.440607967 [bankingserver.O] new connection accepted
>> 1679988788.443309872 [bankingserver.O] transfer from me 2916
>> 1679988788.445407455 [bankingserver.O] transfer to user given 1100
>> 1679988788.452317397 [bankingserver.O] new connection accepted
>> 1679988788.454471381 [bankingserver.O] called balance 2916
```