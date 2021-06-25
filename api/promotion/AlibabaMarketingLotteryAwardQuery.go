package promotion

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/promotion"
)

/* 
抽奖平台查询可用奖品接口 
alibaba.marketing.lottery.award.query

抽奖平台查询可用奖品接口
*/
func AlibabaMarketingLotteryAwardQuery(clt *core.SDKClient, req *promotion.AlibabaMarketingLotteryAwardQueryRequest, session string) (*promotion.AlibabaMarketingLotteryAwardQueryResponse, error) {
    var resp promotion.AlibabaMarketingLotteryAwardQueryAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
