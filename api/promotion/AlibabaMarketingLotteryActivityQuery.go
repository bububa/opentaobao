package promotion

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/promotion"
)

/* 
抽奖平台奖池查询接口 
alibaba.marketing.lottery.activity.query

抽奖平台奖池查询接口
*/
func AlibabaMarketingLotteryActivityQuery(clt *core.SDKClient, req *promotion.AlibabaMarketingLotteryActivityQueryRequest, session string) (*promotion.AlibabaMarketingLotteryActivityQueryAPIResponse, error) {
    var resp promotion.AlibabaMarketingLotteryActivityQueryAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
