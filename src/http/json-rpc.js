function CallJsonRpc(sServiceMethod, oParams, _oSuccessCb, _oErrorCb) {

    var url = "rpc";

    var req = new XMLHttpRequest;

    var rpc = {};
    rpc.method = sServiceMethod;
    rpc.params = [ oParams ];
    rpc.id = 1;
    rpc.jsonrpc = "2.0";

    rpc = JSON.stringify(rpc);

    req.open("POST", encodeURIComponent(url), true);

    req.setRequestHeader("Content-Type","application/json");

    req.onload = function() {
        if (req.status == 400) {
            //if (req.error && req.error != null) {
                _oErrorCb(req.response);
            //}
        } else if (req.status == 200) {
            var oResponse = JSON.parse(req.response)
            if (oResponse.error && oResponse.error != null) {
                _oErrorCb(oResponse.error);
            }
            else {
                var oResult = oResponse.result;
                _oSuccessCb(oResult);
            }
        }
    }

    req.send(rpc);

}
