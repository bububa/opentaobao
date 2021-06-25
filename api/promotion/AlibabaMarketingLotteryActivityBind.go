package promotion

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/promotion"
)

/* 
抽奖平台奖池绑定接口 
alibaba.marketing.lottery.activity.bind

抽奖平台奖池关联接口
*/
func AlibabaMarketingLotteryActivityBind(clt *core.SDKClient, req *promotion.AlibabaMarketingLotteryActivityBindRequest, session string) (*promotion.AlibabaMarketingLotteryActivityBindResponse, error) {
    var resp promotion.AlibabaMarketingLotteryActivityBindAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
