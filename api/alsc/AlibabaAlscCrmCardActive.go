package alsc

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/alsc"
)

/* 
标准激活卡 
alibaba.alsc.crm.card.active

激活卡
*/
func AlibabaAlscCrmCardActive(clt *core.SDKClient, req *alsc.AlibabaAlscCrmCardActiveRequest, session string) (*alsc.AlibabaAlscCrmCardActiveAPIResponse, error) {
    var resp alsc.AlibabaAlscCrmCardActiveAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
