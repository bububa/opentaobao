package promotion

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
定向优惠活动名称与描述违禁词检查 APIRequest
taobao.marketing.promotion.kfc

活动名称与描述违禁词检查
*/
type TaobaoMarketingPromotionKfcRequest struct {
    model.Params

    // 活动名称
    promotionTitle   string 

    // 活动描述
    promotionDesc   string 

}

func NewTaobaoMarketingPromotionKfcRequest() *TaobaoMarketingPromotionKfcRequest{
    return &TaobaoMarketingPromotionKfcRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoMarketingPromotionKfcRequest) GetApiMethodName() string {
    return "taobao.marketing.promotion.kfc"
}

func (r TaobaoMarketingPromotionKfcRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoMarketingPromotionKfcRequest) SetPromotionTitle(promotionTitle string) error {
    r.promotionTitle = promotionTitle
    r.Set("promotion_title", promotionTitle)
    return nil
}

func (r TaobaoMarketingPromotionKfcRequest) GetPromotionTitle() string {
    return r.promotionTitle
}

func (r *TaobaoMarketingPromotionKfcRequest) SetPromotionDesc(promotionDesc string) error {
    r.promotionDesc = promotionDesc
    r.Set("promotion_desc", promotionDesc)
    return nil
}

func (r TaobaoMarketingPromotionKfcRequest) GetPromotionDesc() string {
    return r.promotionDesc
}

