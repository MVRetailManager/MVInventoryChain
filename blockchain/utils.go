package blockchain

import logging "github.com/MVRetailManager/MVInventoryChain/logging"

func HandleError(err error) {
	if err != nil {
		logging.ErrorLogger.Printf(err.Error())
	}
}
