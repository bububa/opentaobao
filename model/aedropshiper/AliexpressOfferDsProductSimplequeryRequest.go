package aedropshiper

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
Dropshipper查询单个商品的简易信息 API请求
aliexpress.offer.ds.product.simplequery

提供给Dropshipper的通过商品ID查找商品简易信息（包括SKU-价格/库存、产品状态等）的接口，只有特定买家可以使用
*/
type AliexpressOfferDsProductSimplequeryRequest struct {
    model.Params
    // 商品ID
    productId   int64
    // 国家
    localCountry   string
    // 语言
    localLanguage   string
}

// 初始化AliexpressOfferDsProductSimplequeryRequest对象
func NewAliexpressOfferDsProductSimplequeryRequest() *AliexpressOfferDsProductSimplequeryRequest{
    return &AliexpressOfferDsProductSimplequeryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AliexpressOfferDsProductSimplequeryRequest) GetApiMethodName() string {
    return "aliexpress.offer.ds.product.simplequery"
}

// IRequest interface 方法, 获取API参数
func (r AliexpressOfferDsProductSimplequeryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ProductId Setter
// 商品ID
func (r *AliexpressOfferDsProductSimplequeryRequest) SetProductId(productId int64) error {
    r.productId = productId
    r.Set("product_id", productId)
    return nil
}

// ProductId Getter
func (r AliexpressOfferDsProductSimplequeryRequest) GetProductId() int64 {
    return r.productId
}
// LocalCountry Setter
// 国家
func (r *AliexpressOfferDsProductSimplequeryRequest) SetLocalCountry(localCountry string) error {
    r.localCountry = localCountry
    r.Set("local_country", localCountry)
    return nil
}

// LocalCountry Getter
func (r AliexpressOfferDsProductSimplequeryRequest) GetLocalCountry() string {
    return r.localCountry
}
// LocalLanguage Setter
// 语言
func (r *AliexpressOfferDsProductSimplequeryRequest) SetLocalLanguage(localLanguage string) error {
    r.localLanguage = localLanguage
    r.Set("local_language", localLanguage)
    return nil
}

// LocalLanguage Getter
func (r AliexpressOfferDsProductSimplequeryRequest) GetLocalLanguage() string {
    return r.localLanguage
}
