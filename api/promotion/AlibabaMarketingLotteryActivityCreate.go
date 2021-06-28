package promotion

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/promotion"
)

/* 
抽奖平台奖池创建接口 
alibaba.marketing.lottery.activity.create

抽奖平台奖池创建接口
*/
func AlibabaMarketingLotteryActivityCreate(clt *core.SDKClient, req *promotion.AlibabaMarketingLotteryActivityCreateRequest, session string) (*promotion.AlibabaMarketingLotteryActivityCreateAPIResponse, error) {
    var resp promotion.AlibabaMarketingLotteryActivityCreateAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
