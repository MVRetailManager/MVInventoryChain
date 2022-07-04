package cli

import (
	"flag"
	"fmt"
	"os"
	"runtime"
	"time"

	blockchainPKG "github.com/MVRetailManager/MVInventoryChain/blockchain"
	"github.com/MVRetailManager/MVInventoryChain/logging"
	"github.com/MVRetailManager/MVInventoryChain/wallet"
)

type CLI struct{}

var bc blockchainPKG.Blockchain

func (cli *CLI) printUsage() {
	logging.InfoLogger.Println("Usage command executed.")

	fmt.Println("Usage:")
	fmt.Println(" getbalance -address ADDRESS - Get balance of ADDRESS")
	fmt.Println(" createblockchain -address ADDRESS - Create a new blockchain and save it to disk, ADDRESS will be the address of the coinbase transaction")
	fmt.Println(" printchain - Print the blockchain")
	fmt.Println(" send -from FROM -to TO- amount AMOUNT - Send AMOUNT of coins from FROM to TO")
	fmt.Println(" createwallet - Create a new wallet")
	fmt.Println(" listaddresses - List all addresses stored within wallet file")
}

func (cli *CLI) validateArgs() {
	if len(os.Args) < 2 {
		cli.printUsage()
		runtime.Goexit()
	}
}

func (cli *CLI) printChain() {
	bc.ContinueBlockchain("")
	defer bc.Database.Close()

	iter := bc.Iterator()

	for {
		block, err := iter.Next()
		blockchainPKG.HandleError(err)

		fmt.Printf("\nS==========%d==========S\n", block.Index)
		fmt.Printf("Previous Hash: 	%x\n", block.PreviousHash)
		fmt.Printf("Hash: 			%x\n", block.Hash)
		fmt.Printf("Nonce: 			%d\n", block.Nonce)
		fmt.Printf("Difficulty: 		%d\n", block.Difficulty)
		fmt.Printf("TimeStamp: 		%d\n", block.UnixTimeStamp)
		fmt.Printf("Transactions:\n")

		for _, tx := range block.Transaction {
			fmt.Println(tx)
		}

		fmt.Printf("E==========%d==========E\n", block.Index)
	}
}

func (cli *CLI) createBlockchain(address string) {
	if !wallet.ValidateAddress(address) {
		fmt.Printf("%s is not a valid address\n", address)
		logging.ErrorLogger.Printf("%s is not a valid address", address)
		runtime.Goexit()
	}

	bc.InitBlockchain(address)
	bc.Database.Close()
	fmt.Println("Success!")
}

func (cli *CLI) getBalance(address string) {
	if !wallet.ValidateAddress(address) {
		fmt.Printf("%s is not a valid address\n", address)
		logging.ErrorLogger.Printf("%s is not a valid address", address)
		runtime.Goexit()
	}

	bc.ContinueBlockchain(address)
	defer bc.Database.Close()

	balance := 0

	finalHash := wallet.Base58Encode([]byte(address))
	publicKeyHash := finalHash[1 : len(finalHash)-4]

	UTXOs := bc.HandleUnspentTxs(publicKeyHash)

	for _, out := range UTXOs {
		balance += out.Value
	}

	fmt.Printf("Balance of '%s': %d\n", address, balance)
}

func (cli *CLI) send(from, to string, amount int) {
	if !wallet.ValidateAddress(to) {
		fmt.Printf("%s is not a valid address\n", to)
		logging.ErrorLogger.Printf("%s is not a valid address", to)
		runtime.Goexit()
	}
	if !wallet.ValidateAddress(from) {
		fmt.Printf("%s is not a valid address\n", from)
		logging.ErrorLogger.Printf("%s is not a valid address", from)
		runtime.Goexit()
	}

	bc.ContinueBlockchain(from)
	defer bc.Database.Close()

	tx := blockchainPKG.NewTransaction(from, to, amount, &bc)

	nbIndex, _ := bc.Database.Size()
	bc.AddBlock(*blockchainPKG.NewBlock(int(nbIndex), time.Now().UTC().UnixNano(), bc.LastHash, []*blockchainPKG.Transaction{tx}))

	fmt.Print("Success!")
}

func (cli *CLI) createWallet() {
	wallets, _ := wallet.CreateWallets()
	address := wallets.AddWallet()
	wallets.SaveFile()

	fmt.Printf("New address: %s\n", address)
}

func (cli *CLI) listAddresses() {
	wallets, _ := wallet.CreateWallets()
	addresses := wallets.GetAllAddresses()

	for _, address := range addresses {
		fmt.Println(address)
	}
}

func (cli *CLI) Run() {
	bc = blockchainPKG.Blockchain{}

	cli.validateArgs()

	getBalanceCmd := flag.NewFlagSet("getbalance", flag.ExitOnError)
	createBlockchainCmd := flag.NewFlagSet("createblockchain", flag.ExitOnError)
	printChainCmd := flag.NewFlagSet("printchain", flag.ExitOnError)
	sendCmd := flag.NewFlagSet("send", flag.ExitOnError)
	createWalletsCmd := flag.NewFlagSet("createwallet", flag.ExitOnError)
	listAddressesCmd := flag.NewFlagSet("listaddresses", flag.ExitOnError)

	getBalanceAddress := getBalanceCmd.String("address", "", "The address to get balance for")
	createBlockchainAddress := createBlockchainCmd.String("address", "", "The address to send genesis block reward to")
	sendFrom := sendCmd.String("from", "", "Source wallet address")
	sendTo := sendCmd.String("to", "", "Destination wallet address")
	sendAmount := sendCmd.Int("amount", 0, "Amount to send")

	switch os.Args[1] {
	case "getbalance":
		err := getBalanceCmd.Parse(os.Args[2:])
		if err != nil {
			logging.ErrorLogger.Panic(err)
		}
	case "createblockchain":
		err := createBlockchainCmd.Parse(os.Args[2:])
		if err != nil {
			logging.ErrorLogger.Panic(err)
		}
	case "printchain":
		err := printChainCmd.Parse(os.Args[2:])
		if err != nil {
			logging.ErrorLogger.Panic(err)
		}
	case "send":
		err := sendCmd.Parse(os.Args[2:])
		if err != nil {
			logging.ErrorLogger.Panic(err)
		}
	case "createwallet":
		err := createWalletsCmd.Parse(os.Args[2:])
		if err != nil {
			logging.ErrorLogger.Panic(err)
		}
	case "listaddresses":
		err := listAddressesCmd.Parse(os.Args[2:])
		if err != nil {
			logging.ErrorLogger.Panic(err)
		}
	default:
		cli.printUsage()
		runtime.Goexit()
	}

	if getBalanceCmd.Parsed() {
		if *getBalanceAddress == "" {
			getBalanceCmd.Usage()
			runtime.Goexit()
		}
		cli.getBalance(*getBalanceAddress)
	}

	if createBlockchainCmd.Parsed() {
		if *createBlockchainAddress == "" {
			createBlockchainCmd.Usage()
			runtime.Goexit()
		}
		cli.createBlockchain(*createBlockchainAddress)
	}

	if printChainCmd.Parsed() {
		cli.printChain()
	}

	if sendCmd.Parsed() {
		if *sendFrom == "" || *sendTo == "" || *sendAmount <= 0 {
			sendCmd.Usage()
			runtime.Goexit()
		}

		cli.send(*sendFrom, *sendTo, *sendAmount)
	}

	if createWalletsCmd.Parsed() {
		cli.createWallet()
	}

	if listAddressesCmd.Parsed() {
		cli.listAddresses()
	}

	runtime.Goexit()
}
