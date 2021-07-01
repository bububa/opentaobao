package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
营销档期活动创建 API请求
alibaba.price.promotion.create

大润发-盒马帮提供新增创建营销活动
*/
type AlibabaPricePromotionCreateAPIRequest struct {
    model.Params
    // 档期活动参数
    _promotionActivityDo   *PromotionActivityDo
}

// 初始化AlibabaPricePromotionCreateAPIRequest对象
func NewAlibabaPricePromotionCreateRequest() *AlibabaPricePromotionCreateAPIRequest{
    return &AlibabaPricePromotionCreateAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaPricePromotionCreateAPIRequest) GetApiMethodName() string {
    return "alibaba.price.promotion.create"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaPricePromotionCreateAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// PromotionActivityDo Setter
// 档期活动参数
func (r *AlibabaPricePromotionCreateAPIRequest) SetPromotionActivityDo(_promotionActivityDo *PromotionActivityDo) error {
    r._promotionActivityDo = _promotionActivityDo
    r.Set("promotion_activity_do", _promotionActivityDo)
    return nil
}

// PromotionActivityDo Getter
func (r AlibabaPricePromotionCreateAPIRequest) GetPromotionActivityDo() *PromotionActivityDo {
    return r._promotionActivityDo
}
