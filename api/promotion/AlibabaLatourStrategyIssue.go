package promotion

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/promotion"
)

/* 
阿里巴巴权益发放接口 
alibaba.latour.strategy.issue

阿里巴巴权益平台权益发放接口
*/
func AlibabaLatourStrategyIssue(clt *core.SDKClient, req *promotion.AlibabaLatourStrategyIssueAPIRequest, session string) (*promotion.AlibabaLatourStrategyIssueAPIResponse, error) {
    var resp promotion.AlibabaLatourStrategyIssueAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
