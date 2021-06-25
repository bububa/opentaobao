package promotion

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/promotion"
)

/* 
抽奖平台抽奖方案创建接口 
alibaba.marketing.lottery.schema.create

抽奖平台抽奖方案创建接口
*/
func AlibabaMarketingLotterySchemaCreate(clt *core.SDKClient, req *promotion.AlibabaMarketingLotterySchemaCreateRequest, session string) (*promotion.AlibabaMarketingLotterySchemaCreateResponse, error) {
    var resp promotion.AlibabaMarketingLotterySchemaCreateAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
