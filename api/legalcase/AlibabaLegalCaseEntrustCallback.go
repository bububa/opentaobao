package legalcase

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/legalcase"
)

/* 
委托回调接口 
alibaba.legal.case.entrust.callback

委托回调接口
*/
func AlibabaLegalCaseEntrustCallback(clt *core.SDKClient, req *legalcase.AlibabaLegalCaseEntrustCallbackRequest, session string) (*legalcase.AlibabaLegalCaseEntrustCallbackResponse, error) {
    var resp legalcase.AlibabaLegalCaseEntrustCallbackAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
