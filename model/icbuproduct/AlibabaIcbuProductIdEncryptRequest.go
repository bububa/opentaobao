package icbuproduct

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
ICBU国际站商品加密接口 API请求
alibaba.icbu.product.id.encrypt

ICBU国际站，对混淆的产品ID加密。
*/
type AlibabaIcbuProductIdEncryptRequest struct {
    model.Params
    // 语种
    language   string
    // 明文id
    productId   int64
}

// 初始化AlibabaIcbuProductIdEncryptRequest对象
func NewAlibabaIcbuProductIdEncryptRequest() *AlibabaIcbuProductIdEncryptRequest{
    return &AlibabaIcbuProductIdEncryptRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaIcbuProductIdEncryptRequest) GetApiMethodName() string {
    return "alibaba.icbu.product.id.encrypt"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaIcbuProductIdEncryptRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Language Setter
// 语种
func (r *AlibabaIcbuProductIdEncryptRequest) SetLanguage(language string) error {
    r.language = language
    r.Set("language", language)
    return nil
}

// Language Getter
func (r AlibabaIcbuProductIdEncryptRequest) GetLanguage() string {
    return r.language
}
// ProductId Setter
// 明文id
func (r *AlibabaIcbuProductIdEncryptRequest) SetProductId(productId int64) error {
    r.productId = productId
    r.Set("product_id", productId)
    return nil
}

// ProductId Getter
func (r AlibabaIcbuProductIdEncryptRequest) GetProductId() int64 {
    return r.productId
}
