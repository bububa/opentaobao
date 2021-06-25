package icbu

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商品ID解密 APIRequest
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

func NewAlibabaIcbuProductIdDecryptRequest() *AlibabaIcbuProductIdDecryptRequest{
    return &AlibabaIcbuProductIdDecryptRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaIcbuProductIdDecryptRequest) GetApiMethodName() string {
    return "alibaba.icbu.product.id.decrypt"
}

func (r AlibabaIcbuProductIdDecryptRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaIcbuProductIdDecryptRequest) SetLanguage(language string) error {
    r.language = language
    r.Set("language", language)
    return nil
}

func (r AlibabaIcbuProductIdDecryptRequest) GetLanguage() string {
    return r.language
}

func (r *AlibabaIcbuProductIdDecryptRequest) SetProductId(productId string) error {
    r.productId = productId
    r.Set("product_id", productId)
    return nil
}

func (r AlibabaIcbuProductIdDecryptRequest) GetProductId() string {
    return r.productId
}

