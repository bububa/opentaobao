package icbu

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商品ID解密 API请求
alibaba.icbu.product.id.decrypt

对混淆的产品ID解密
*/
type AlibabaIcbuProductIdDecryptRequest struct {
    model.Params
    // 语种
    language   string
    // 混淆后的商品ID
    productId   string
}

// 初始化AlibabaIcbuProductIdDecryptRequest对象
func NewAlibabaIcbuProductIdDecryptRequest() *AlibabaIcbuProductIdDecryptRequest{
    return &AlibabaIcbuProductIdDecryptRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaIcbuProductIdDecryptRequest) GetApiMethodName() string {
    return "alibaba.icbu.product.id.decrypt"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaIcbuProductIdDecryptRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Language Setter
// 语种
func (r *AlibabaIcbuProductIdDecryptRequest) SetLanguage(language string) error {
    r.language = language
    r.Set("language", language)
    return nil
}

// Language Getter
func (r AlibabaIcbuProductIdDecryptRequest) GetLanguage() string {
    return r.language
}
// ProductId Setter
// 混淆后的商品ID
func (r *AlibabaIcbuProductIdDecryptRequest) SetProductId(productId string) error {
    r.productId = productId
    r.Set("product_id", productId)
    return nil
}

// ProductId Getter
func (r AlibabaIcbuProductIdDecryptRequest) GetProductId() string {
    return r.productId
}
