var url;
var url_host = "http://localhost:5000/";
var ccId;
var counterList;
var user_name;

//議案の登録
function regist() {
	url = url_host + "chaincode";
	var number = $("#subject__no").val();
	var subject = $("#subject__title").val();
	var content = $("#subject__content").val();
	var sub = [number, subject, content];

	var JSONdata = createJSONdataForShareApp("invoke","regist", sub, 1);
	console.log(JSON.stringify(JSONdata)); //値をテスト確認
	executeJsonRpc(url, JSONdata,
    	function success(data) {
        	console.log("success!");
        },
        function error(data) {
        	console.log("error");
        }
	);
}

//JSONメッセージ作成
function createJSONdataForShareApp(method, functionName, args, id) {
	var JSONdata = {
		jsonrpc: "2.0",
		method: method,
        params: {
        	type: 1,
            ctorMsg: {
            	function: functionName,
                          args: args
                },
                secureContext: user_name,
            },
            id: id
        };
        // チェーンコードIDを設定
        JSONdata.params["chaincodeID"] = {
        	path: "github.com/hyperledger/fabric/examples/chaincode/go/chaincode_counter"
        };
        return JSONdata;
}

//JSON-RPC実行
function executeJsonRpc(url_exec, JSONdata, success, error) {
	$.ajax({
    	type: 'post',
        url: url_exec,
        data: JSON.stringify(JSONdata),
        contentType: 'application/JSON;',
        dataType: 'JSON',
        scriptCharset: 'utf-8',
        success: function(data) {
        	success(data);
        },
        error: function(data) {
        	error(data);
        }
	});
}
