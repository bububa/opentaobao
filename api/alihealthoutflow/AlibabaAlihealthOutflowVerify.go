package alihealthoutflow

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/alihealthoutflow"
)

/* 
处方外流药店通过核销码核销处方 
alibaba.alihealth.outflow.verify

处方外流药店通过核销码核销处方
*/
func AlibabaAlihealthOutflowVerify(clt *core.SDKClient, req *alihealthoutflow.AlibabaAlihealthOutflowVerifyRequest, session string) (*alihealthoutflow.AlibabaAlihealthOutflowVerifyAPIResponse, error) {
    var resp alihealthoutflow.AlibabaAlihealthOutflowVerifyAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
