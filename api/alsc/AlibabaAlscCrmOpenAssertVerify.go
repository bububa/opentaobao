package alsc

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/alsc"
)

/* 
资产核销接口 
alibaba.alsc.crm.open.assert.verify

核销储值，积分，券资产
*/
func AlibabaAlscCrmOpenAssertVerify(clt *core.SDKClient, req *alsc.AlibabaAlscCrmOpenAssertVerifyAPIRequest, session string) (*alsc.AlibabaAlscCrmOpenAssertVerifyAPIResponse, error) {
    var resp alsc.AlibabaAlscCrmOpenAssertVerifyAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
