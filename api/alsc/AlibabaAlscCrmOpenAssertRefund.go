package alsc

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/alsc"
)

/* 
资产核销回退接口 
alibaba.alsc.crm.open.assert.refund

回退已经核销的储值积分券资产
*/
func AlibabaAlscCrmOpenAssertRefund(clt *core.SDKClient, req *alsc.AlibabaAlscCrmOpenAssertRefundAPIRequest, session string) (*alsc.AlibabaAlscCrmOpenAssertRefundAPIResponse, error) {
    var resp alsc.AlibabaAlscCrmOpenAssertRefundAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
