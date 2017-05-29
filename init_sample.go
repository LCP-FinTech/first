package main

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

// SimpleChaincode example simple Chaincode implementation
type SimpleChaincode struct {
}
// ユーザーの設定
var userA, userB, userC int //string
//var Aval, Bval, Cval int
var err error

//init処理
func (t *SimpleChaincode) Init(stub *shim.ChaincodeStub, function string, args []string) ([]byte, error) {
   
   userA = 100
   userB = 200
   userC = 400
   
    //Aval = 100
    //Bval = 200
    //Cval = 400
    
    //ユーザーアカウントと割り振りは別途管理が？？今回の仕様ならユーザー名入力して数値設定すればOKな気がする。
    
        return nil, errors.New("Incorrect number of arguments. Expecting 0")
    }

    return nil, nil
}
