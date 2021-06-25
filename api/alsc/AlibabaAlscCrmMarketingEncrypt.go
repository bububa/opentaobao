package alsc

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/alsc"
)

/* 
加密 
alibaba.alsc.crm.marketing.encrypt

加密
*/
func AlibabaAlscCrmMarketingEncrypt(clt *core.SDKClient, req *alsc.AlibabaAlscCrmMarketingEncryptRequest, session string) (*alsc.AlibabaAlscCrmMarketingEncryptResponse, error) {
    var resp alsc.AlibabaAlscCrmMarketingEncryptAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
