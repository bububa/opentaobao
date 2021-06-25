package promotion

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/promotion"
)

/* 
定向优惠活动名称与描述违禁词检查 
taobao.marketing.promotion.kfc

活动名称与描述违禁词检查
*/
func TaobaoMarketingPromotionKfc(clt *core.SDKClient, req *promotion.TaobaoMarketingPromotionKfcRequest, session string) (*promotion.TaobaoMarketingPromotionKfcResponse, error) {
    var resp promotion.TaobaoMarketingPromotionKfcAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
