package alsc

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/alsc"
)

/* 
批量激活卡 
alibaba.alsc.crm.card.batch.active

批量激活卡
*/
func AlibabaAlscCrmCardBatchActive(clt *core.SDKClient, req *alsc.AlibabaAlscCrmCardBatchActiveRequest, session string) (*alsc.AlibabaAlscCrmCardBatchActiveResponse, error) {
    var resp alsc.AlibabaAlscCrmCardBatchActiveAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
