package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
删除档期活动 API请求
alibaba.price.promotion.activity.delete

删除盒马帮档期活动
*/
type AlibabaPricePromotionActivityDeleteRequest struct {
    model.Params
    // 外部主键
    _outerPromotionCode   string
    // 经营店OU
    _ouCode   string
}

// 初始化AlibabaPricePromotionActivityDeleteRequest对象
func NewAlibabaPricePromotionActivityDeleteRequest() *AlibabaPricePromotionActivityDeleteRequest{
    return &AlibabaPricePromotionActivityDeleteRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaPricePromotionActivityDeleteRequest) GetApiMethodName() string {
    return "alibaba.price.promotion.activity.delete"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaPricePromotionActivityDeleteRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OuterPromotionCode Setter
// 外部主键
func (r *AlibabaPricePromotionActivityDeleteRequest) SetOuterPromotionCode(_outerPromotionCode string) error {
    r._outerPromotionCode = _outerPromotionCode
    r.Set("outer_promotion_code", _outerPromotionCode)
    return nil
}

// OuterPromotionCode Getter
func (r AlibabaPricePromotionActivityDeleteRequest) GetOuterPromotionCode() string {
    return r._outerPromotionCode
}
// OuCode Setter
// 经营店OU
func (r *AlibabaPricePromotionActivityDeleteRequest) SetOuCode(_ouCode string) error {
    r._ouCode = _ouCode
    r.Set("ou_code", _ouCode)
    return nil
}

// OuCode Getter
func (r AlibabaPricePromotionActivityDeleteRequest) GetOuCode() string {
    return r._ouCode
}
