package alicom

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/alicom"
)

/* 
判断支付宝用户是否绑定淘宝账号 
alibaba.aliqin.flow.alipay.isbindingtbaccount

判断支付宝用户是否绑定淘宝账号
*/
func AlibabaAliqinFlowAlipayIsbindingtbaccount(clt *core.SDKClient, req *alicom.AlibabaAliqinFlowAlipayIsbindingtbaccountAPIRequest, session string) (*alicom.AlibabaAliqinFlowAlipayIsbindingtbaccountAPIResponse, error) {
    var resp alicom.AlibabaAliqinFlowAlipayIsbindingtbaccountAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
