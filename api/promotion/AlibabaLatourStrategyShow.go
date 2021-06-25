package promotion

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/promotion"
)

/* 
阿里巴巴权益投放接口 
alibaba.latour.strategy.show

阿里巴巴权益平台权益投放接口
*/
func AlibabaLatourStrategyShow(clt *core.SDKClient, req *promotion.AlibabaLatourStrategyShowRequest, session string) (*promotion.AlibabaLatourStrategyShowResponse, error) {
    var resp promotion.AlibabaLatourStrategyShowAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
