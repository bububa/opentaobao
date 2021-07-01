package alsc

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/alsc"
)

/* 
批量开卡（售卡） 
alibaba.alsc.crm.card.batch.sell

批量开卡（售卡）
*/
func AlibabaAlscCrmCardBatchSell(clt *core.SDKClient, req *alsc.AlibabaAlscCrmCardBatchSellAPIRequest, session string) (*alsc.AlibabaAlscCrmCardBatchSellAPIResponse, error) {
    var resp alsc.AlibabaAlscCrmCardBatchSellAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
