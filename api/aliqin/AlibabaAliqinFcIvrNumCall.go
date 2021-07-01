package aliqin

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/aliqin"
)

/* 
ivr呼叫 
alibaba.aliqin.fc.ivr.num.call

ivr呼叫
*/
func AlibabaAliqinFcIvrNumCall(clt *core.SDKClient, req *aliqin.AlibabaAliqinFcIvrNumCallAPIRequest, session string) (*aliqin.AlibabaAliqinFcIvrNumCallAPIResponse, error) {
    var resp aliqin.AlibabaAliqinFcIvrNumCallAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
