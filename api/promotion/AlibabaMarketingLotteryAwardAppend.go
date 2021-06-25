package promotion

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/promotion"
)

/* 
抽奖平台奖品添加接口 
alibaba.marketing.lottery.award.append

抽奖平台奖品添加接口，目前仅用于奖池众筹项目
*/
func AlibabaMarketingLotteryAwardAppend(clt *core.SDKClient, req *promotion.AlibabaMarketingLotteryAwardAppendRequest, session string) (*promotion.AlibabaMarketingLotteryAwardAppendResponse, error) {
    var resp promotion.AlibabaMarketingLotteryAwardAppendAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
