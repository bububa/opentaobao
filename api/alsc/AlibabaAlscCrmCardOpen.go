package alsc

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/alsc"
)

/* 
标准开卡流程 
alibaba.alsc.crm.card.open

标准开卡流程
*/
func AlibabaAlscCrmCardOpen(clt *core.SDKClient, req *alsc.AlibabaAlscCrmCardOpenAPIRequest, session string) (*alsc.AlibabaAlscCrmCardOpenAPIResponse, error) {
    var resp alsc.AlibabaAlscCrmCardOpenAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
