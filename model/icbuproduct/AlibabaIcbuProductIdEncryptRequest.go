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
    _language   string
    // 明文id
    _productId   int64
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
func (r *AlibabaIcbuProductIdEncryptRequest) SetLanguage(_language string) error {
    r._language = _language
    r.Set("language", _language)
    return nil
}

// Language Getter
func (r AlibabaIcbuProductIdEncryptRequest) GetLanguage() string {
    return r._language
}
// ProductId Setter
// 明文id
func (r *AlibabaIcbuProductIdEncryptRequest) SetProductId(_productId int64) error {
    r._productId = _productId
    r.Set("product_id", _productId)
    return nil
}

// ProductId Getter
func (r AlibabaIcbuProductIdEncryptRequest) GetProductId() int64 {
    return r._productId
}
