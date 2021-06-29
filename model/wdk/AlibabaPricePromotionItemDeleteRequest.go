package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
批量删除档期 API请求
alibaba.price.promotion.item.delete

盒马帮批量删除档期商品
*/
type AlibabaPricePromotionItemDeleteRequest struct {
    model.Params
    // 商品code
    _skuCodes   []string
    // toB渠道店OU
    _ouCode   string
    // 外部档期编码
    _outerPromotionCode   string
    // 盒马唯一标识
    _uniqueId   string
}

// 初始化AlibabaPricePromotionItemDeleteRequest对象
func NewAlibabaPricePromotionItemDeleteRequest() *AlibabaPricePromotionItemDeleteRequest{
    return &AlibabaPricePromotionItemDeleteRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaPricePromotionItemDeleteRequest) GetApiMethodName() string {
    return "alibaba.price.promotion.item.delete"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaPricePromotionItemDeleteRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// SkuCodes Setter
// 商品code
func (r *AlibabaPricePromotionItemDeleteRequest) SetSkuCodes(_skuCodes []string) error {
    r._skuCodes = _skuCodes
    r.Set("sku_codes", _skuCodes)
    return nil
}

// SkuCodes Getter
func (r AlibabaPricePromotionItemDeleteRequest) GetSkuCodes() []string {
    return r._skuCodes
}
// OuCode Setter
// toB渠道店OU
func (r *AlibabaPricePromotionItemDeleteRequest) SetOuCode(_ouCode string) error {
    r._ouCode = _ouCode
    r.Set("ou_code", _ouCode)
    return nil
}

// OuCode Getter
func (r AlibabaPricePromotionItemDeleteRequest) GetOuCode() string {
    return r._ouCode
}
// OuterPromotionCode Setter
// 外部档期编码
func (r *AlibabaPricePromotionItemDeleteRequest) SetOuterPromotionCode(_outerPromotionCode string) error {
    r._outerPromotionCode = _outerPromotionCode
    r.Set("outer_promotion_code", _outerPromotionCode)
    return nil
}

// OuterPromotionCode Getter
func (r AlibabaPricePromotionItemDeleteRequest) GetOuterPromotionCode() string {
    return r._outerPromotionCode
}
// UniqueId Setter
// 盒马唯一标识
func (r *AlibabaPricePromotionItemDeleteRequest) SetUniqueId(_uniqueId string) error {
    r._uniqueId = _uniqueId
    r.Set("unique_id", _uniqueId)
    return nil
}

// UniqueId Getter
func (r AlibabaPricePromotionItemDeleteRequest) GetUniqueId() string {
    return r._uniqueId
}
