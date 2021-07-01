package promotion

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/promotion"
)

/* 
抽奖平台抽奖接口 
alibaba.marketing.lottery.draw.dodraw

抽奖平台PC端抽奖接口
*/
func AlibabaMarketingLotteryDrawDodraw(clt *core.SDKClient, req *promotion.AlibabaMarketingLotteryDrawDodrawAPIRequest, session string) (*promotion.AlibabaMarketingLotteryDrawDodrawAPIResponse, error) {
    var resp promotion.AlibabaMarketingLotteryDrawDodrawAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
