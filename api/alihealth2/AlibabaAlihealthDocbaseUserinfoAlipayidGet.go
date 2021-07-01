package alihealth2

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/alihealth2"
)

/* 
根据健康ID获取支付宝ID 
alibaba.alihealth.docbase.userinfo.alipayid.get

根据健康ID获取支付宝ID
*/
func AlibabaAlihealthDocbaseUserinfoAlipayidGet(clt *core.SDKClient, req *alihealth2.AlibabaAlihealthDocbaseUserinfoAlipayidGetAPIRequest, session string) (*alihealth2.AlibabaAlihealthDocbaseUserinfoAlipayidGetAPIResponse, error) {
    var resp alihealth2.AlibabaAlihealthDocbaseUserinfoAlipayidGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
