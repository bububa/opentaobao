package wdk

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/wdk"
)

/* 
翱象ERP核销回调 
alibaba.tcls.aelophy.bill.verificate.callback

翱象ERP核销回调
*/
func AlibabaTclsAelophyBillVerificateCallback(clt *core.SDKClient, req *wdk.AlibabaTclsAelophyBillVerificateCallbackRequest, session string) (*wdk.AlibabaTclsAelophyBillVerificateCallbackResponse, error) {
    var resp wdk.AlibabaTclsAelophyBillVerificateCallbackAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
