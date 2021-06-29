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
type AlibabaPricePromotionCreateRequest struct {
    model.Params
    // 档期活动参数
    _promotionActivityDo   *PromotionActivityDO
}

// 初始化AlibabaPricePromotionCreateRequest对象
func NewAlibabaPricePromotionCreateRequest() *AlibabaPricePromotionCreateRequest{
    return &AlibabaPricePromotionCreateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaPricePromotionCreateRequest) GetApiMethodName() string {
    return "alibaba.price.promotion.create"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaPricePromotionCreateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// PromotionActivityDo Setter
// 档期活动参数
func (r *AlibabaPricePromotionCreateRequest) SetPromotionActivityDo(_promotionActivityDo *PromotionActivityDO) error {
    r._promotionActivityDo = _promotionActivityDo
    r.Set("promotion_activity_do", _promotionActivityDo)
    return nil
}

// PromotionActivityDo Getter
func (r AlibabaPricePromotionCreateRequest) GetPromotionActivityDo() *PromotionActivityDO {
    return r._promotionActivityDo
}
