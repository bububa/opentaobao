package promotion

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/promotion"
)

/* 
抽奖平台活动删除接口 
alibaba.marketing.lottery.activity.delete

抽奖平台活动删除接口
*/
func AlibabaMarketingLotteryActivityDelete(clt *core.SDKClient, req *promotion.AlibabaMarketingLotteryActivityDeleteRequest, session string) (*promotion.AlibabaMarketingLotteryActivityDeleteAPIResponse, error) {
    var resp promotion.AlibabaMarketingLotteryActivityDeleteAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
