package main

import(
	"errors"
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

type SubjectChaincode struct {
}

// 議案(subject)と内容(content)
/*type Subject struct {
	Subject string `json:"subject"`
	Content string `json:"content"`
}*/

// 初期処理(なし)
func (cc *SubjectChaincode) Init(stub *shim.ChaincodeStub, function string, args []string) ([]byte, error) {
	return nil, nil
}

// 議案を登録
func (cc *SubjectChaincode) Invoke(stub *shim.ChaincodeStub,function string,args []string) ([]byte, error) {
	// function名でハンドリング
	if function == "regist" {
		// 登録実行
		return cc.regist(stub, args)
	}

	return nil, errors.New("Received unknown function")
}

// 議案を参照
func (cc *SubjectChaincode) Query(stub *shim.ChaincodeStub, function string, args []string) ([]byte, error) {
	// function名でハンドリング
	if function == "getsubject" {
		// 議案を取得
		return cc.getsubject(stub, args)
	}

	return nil, errors.New("Received unknown function")
}

// 登録を実行
func (cc *SubjectChaincode) regist(stub *shim.ChaincodeStub, args []string) ([]byte, error) {
	var key, value string
	var err error

	if len(args) != 2 {
		return nil, errors.New("Incorrect number of arguments. Expecting 2. name of the key and value to set")
	}

	// ワールドステートに追加
	key = args[0]
	value = args[1]
	err = stub.PutState(key, []byte(value))
	if err != nil {
		return nil, err
	}
	return nil, nil
}

// 議案の取得
func (cc *SubjectChaincode) getsubject(stub *shim.ChaincodeStub,args []string) ([]byte, error) {
	var key string
	var err error

	if len(args) != 1 {
        return nil, errors.New("Incorrect number of arguments. Expecting name of the key to query")
    	}

    	key = args[0]
    	value, err := stub.GetState(key)
    	if err != nil {
        	return nil, errors.New("No exist")
    	}

    	return value, nil
}

// Validating Peerに接続し、チェーンコードを実行
func main() {
	err := shim.Start(new(SubjectChaincode))
	if err != nil {
		fmt.Printf("Error starting chaincode: %s",err)
	}
}