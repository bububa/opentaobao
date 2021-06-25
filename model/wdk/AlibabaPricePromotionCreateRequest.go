package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
营销档期活动创建 APIRequest
alibaba.price.promotion.create

大润发-盒马帮提供新增创建营销活动
*/
type AlibabaPricePromotionCreateRequest struct {
    model.Params

    // 档期活动参数
    promotionActivityDo   *PromotionActivityDo 

}

func NewAlibabaPricePromotionCreateRequest() *AlibabaPricePromotionCreateRequest{
    return &AlibabaPricePromotionCreateRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaPricePromotionCreateRequest) GetApiMethodName() string {
    return "alibaba.price.promotion.create"
}

func (r AlibabaPricePromotionCreateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaPricePromotionCreateRequest) SetPromotionActivityDo(promotionActivityDo *PromotionActivityDo) error {
    r.promotionActivityDo = promotionActivityDo
    r.Set("promotion_activity_do", promotionActivityDo)
    return nil
}

func (r AlibabaPricePromotionCreateRequest) GetPromotionActivityDo() *PromotionActivityDo {
    return r.promotionActivityDo
}

