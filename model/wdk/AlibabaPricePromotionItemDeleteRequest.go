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
    skuCodes   []string
    // toB渠道店OU
    ouCode   string
    // 外部档期编码
    outerPromotionCode   string
    // 盒马唯一标识
    uniqueId   string
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
func (r *AlibabaPricePromotionItemDeleteRequest) SetSkuCodes(skuCodes []string) error {
    r.skuCodes = skuCodes
    r.Set("sku_codes", skuCodes)
    return nil
}

// SkuCodes Getter
func (r AlibabaPricePromotionItemDeleteRequest) GetSkuCodes() []string {
    return r.skuCodes
}
// OuCode Setter
// toB渠道店OU
func (r *AlibabaPricePromotionItemDeleteRequest) SetOuCode(ouCode string) error {
    r.ouCode = ouCode
    r.Set("ou_code", ouCode)
    return nil
}

// OuCode Getter
func (r AlibabaPricePromotionItemDeleteRequest) GetOuCode() string {
    return r.ouCode
}
// OuterPromotionCode Setter
// 外部档期编码
func (r *AlibabaPricePromotionItemDeleteRequest) SetOuterPromotionCode(outerPromotionCode string) error {
    r.outerPromotionCode = outerPromotionCode
    r.Set("outer_promotion_code", outerPromotionCode)
    return nil
}

// OuterPromotionCode Getter
func (r AlibabaPricePromotionItemDeleteRequest) GetOuterPromotionCode() string {
    return r.outerPromotionCode
}
// UniqueId Setter
// 盒马唯一标识
func (r *AlibabaPricePromotionItemDeleteRequest) SetUniqueId(uniqueId string) error {
    r.uniqueId = uniqueId
    r.Set("unique_id", uniqueId)
    return nil
}

// UniqueId Getter
func (r AlibabaPricePromotionItemDeleteRequest) GetUniqueId() string {
    return r.uniqueId
}
