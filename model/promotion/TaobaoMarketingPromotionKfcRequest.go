package promotion

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
定向优惠活动名称与描述违禁词检查 API请求
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

// 初始化TaobaoMarketingPromotionKfcRequest对象
func NewTaobaoMarketingPromotionKfcRequest() *TaobaoMarketingPromotionKfcRequest{
    return &TaobaoMarketingPromotionKfcRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoMarketingPromotionKfcRequest) GetApiMethodName() string {
    return "taobao.marketing.promotion.kfc"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoMarketingPromotionKfcRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// PromotionTitle Setter
// 活动名称
func (r *TaobaoMarketingPromotionKfcRequest) SetPromotionTitle(promotionTitle string) error {
    r.promotionTitle = promotionTitle
    r.Set("promotion_title", promotionTitle)
    return nil
}

// PromotionTitle Getter
func (r TaobaoMarketingPromotionKfcRequest) GetPromotionTitle() string {
    return r.promotionTitle
}
// PromotionDesc Setter
// 活动描述
func (r *TaobaoMarketingPromotionKfcRequest) SetPromotionDesc(promotionDesc string) error {
    r.promotionDesc = promotionDesc
    r.Set("promotion_desc", promotionDesc)
    return nil
}

// PromotionDesc Getter
func (r TaobaoMarketingPromotionKfcRequest) GetPromotionDesc() string {
    return r.promotionDesc
}
