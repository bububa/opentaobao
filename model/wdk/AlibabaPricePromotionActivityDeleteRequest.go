package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
删除档期活动 APIRequest
alibaba.price.promotion.activity.delete

删除盒马帮档期活动
*/
type AlibabaPricePromotionActivityDeleteRequest struct {
    model.Params

    // 外部主键
    outerPromotionCode   string 

    // 经营店OU
    ouCode   string 

}

func NewAlibabaPricePromotionActivityDeleteRequest() *AlibabaPricePromotionActivityDeleteRequest{
    return &AlibabaPricePromotionActivityDeleteRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaPricePromotionActivityDeleteRequest) GetApiMethodName() string {
    return "alibaba.price.promotion.activity.delete"
}

func (r AlibabaPricePromotionActivityDeleteRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaPricePromotionActivityDeleteRequest) SetOuterPromotionCode(outerPromotionCode string) error {
    r.outerPromotionCode = outerPromotionCode
    r.Set("outer_promotion_code", outerPromotionCode)
    return nil
}

func (r AlibabaPricePromotionActivityDeleteRequest) GetOuterPromotionCode() string {
    return r.outerPromotionCode
}

func (r *AlibabaPricePromotionActivityDeleteRequest) SetOuCode(ouCode string) error {
    r.ouCode = ouCode
    r.Set("ou_code", ouCode)
    return nil
}

func (r AlibabaPricePromotionActivityDeleteRequest) GetOuCode() string {
    return r.ouCode
}

