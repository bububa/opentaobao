package trade

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/trade"
)

/* 
轻pos品牌营销查询接口 
alibaba.wdk.pos.trade.query

轻pos品牌营销场景，外部商家查询营销信息
*/
func AlibabaWdkPosTradeQuery(clt *core.SDKClient, req *trade.AlibabaWdkPosTradeQueryAPIRequest, session string) (*trade.AlibabaWdkPosTradeQueryAPIResponse, error) {
    var resp trade.AlibabaWdkPosTradeQueryAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
