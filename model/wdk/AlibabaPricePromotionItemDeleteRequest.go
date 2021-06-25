package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
批量删除档期 APIRequest
alibaba.price.promotion.item.delete

盒马帮批量删除档期商品
*/
type AlibabaPricePromotionItemDeleteRequest struct {
    model.Params

    // 商品code
    skuCodes   []String 

    // toB渠道店OU
    ouCode   string 

    // 外部档期编码
    outerPromotionCode   string 

    // 盒马唯一标识
    uniqueId   string 

}

func NewAlibabaPricePromotionItemDeleteRequest() *AlibabaPricePromotionItemDeleteRequest{
    return &AlibabaPricePromotionItemDeleteRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaPricePromotionItemDeleteRequest) GetApiMethodName() string {
    return "alibaba.price.promotion.item.delete"
}

func (r AlibabaPricePromotionItemDeleteRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaPricePromotionItemDeleteRequest) SetSkuCodes(skuCodes []String) error {
    r.skuCodes = skuCodes
    r.Set("sku_codes", skuCodes)
    return nil
}

func (r AlibabaPricePromotionItemDeleteRequest) GetSkuCodes() []String {
    return r.skuCodes
}

func (r *AlibabaPricePromotionItemDeleteRequest) SetOuCode(ouCode string) error {
    r.ouCode = ouCode
    r.Set("ou_code", ouCode)
    return nil
}

func (r AlibabaPricePromotionItemDeleteRequest) GetOuCode() string {
    return r.ouCode
}

func (r *AlibabaPricePromotionItemDeleteRequest) SetOuterPromotionCode(outerPromotionCode string) error {
    r.outerPromotionCode = outerPromotionCode
    r.Set("outer_promotion_code", outerPromotionCode)
    return nil
}

func (r AlibabaPricePromotionItemDeleteRequest) GetOuterPromotionCode() string {
    return r.outerPromotionCode
}

func (r *AlibabaPricePromotionItemDeleteRequest) SetUniqueId(uniqueId string) error {
    r.uniqueId = uniqueId
    r.Set("unique_id", uniqueId)
    return nil
}

func (r AlibabaPricePromotionItemDeleteRequest) GetUniqueId() string {
    return r.uniqueId
}

