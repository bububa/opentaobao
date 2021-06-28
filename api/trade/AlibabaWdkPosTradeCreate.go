package trade

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/trade"
)

/* 
轻pos品牌营销下单接口 
alibaba.wdk.pos.trade.create

提供给石基进行轻pos品牌营销下单
*/
func AlibabaWdkPosTradeCreate(clt *core.SDKClient, req *trade.AlibabaWdkPosTradeCreateRequest, session string) (*trade.AlibabaWdkPosTradeCreateAPIResponse, error) {
    var resp trade.AlibabaWdkPosTradeCreateAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
