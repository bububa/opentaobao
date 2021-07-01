package trade

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/trade"
)

/* 
轻pos品牌营销支付接口 
alibaba.wdk.pos.trade.pay

轻pos场景，外部商家支付后调用开放平台把支付信息回传给五道口交易
*/
func AlibabaWdkPosTradePay(clt *core.SDKClient, req *trade.AlibabaWdkPosTradePayAPIRequest, session string) (*trade.AlibabaWdkPosTradePayAPIResponse, error) {
    var resp trade.AlibabaWdkPosTradePayAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
