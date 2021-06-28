package alicom

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/alicom"
)

/* 
运营商获得用户身份信息 
alibaba.aliqin.tcc.trade.identity.get

天猫网厅运营商官方旗舰店获取用户身份信息
*/
func AlibabaAliqinTccTradeIdentityGet(clt *core.SDKClient, req *alicom.AlibabaAliqinTccTradeIdentityGetRequest, session string) (*alicom.AlibabaAliqinTccTradeIdentityGetAPIResponse, error) {
    var resp alicom.AlibabaAliqinTccTradeIdentityGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
