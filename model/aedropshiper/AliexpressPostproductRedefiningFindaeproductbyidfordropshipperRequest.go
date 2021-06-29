package aedropshiper

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
Dropshipper查找商品信息接口 API请求
aliexpress.postproduct.redefining.findaeproductbyidfordropshipper

提供给Dropshipper的通过商品ID查找商品信息的接口，只有特定买家可以使用
*/
type AliexpressPostproductRedefiningFindaeproductbyidfordropshipperRequest struct {
    model.Params
    // 商品ID
    _productId   int64
    // 国家
    _localCountry   string
    // 语言
    _localLanguage   string
}

// 初始化AliexpressPostproductRedefiningFindaeproductbyidfordropshipperRequest对象
func NewAliexpressPostproductRedefiningFindaeproductbyidfordropshipperRequest() *AliexpressPostproductRedefiningFindaeproductbyidfordropshipperRequest{
    return &AliexpressPostproductRedefiningFindaeproductbyidfordropshipperRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AliexpressPostproductRedefiningFindaeproductbyidfordropshipperRequest) GetApiMethodName() string {
    return "aliexpress.postproduct.redefining.findaeproductbyidfordropshipper"
}

// IRequest interface 方法, 获取API参数
func (r AliexpressPostproductRedefiningFindaeproductbyidfordropshipperRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ProductId Setter
// 商品ID
func (r *AliexpressPostproductRedefiningFindaeproductbyidfordropshipperRequest) SetProductId(_productId int64) error {
    r._productId = _productId
    r.Set("product_id", _productId)
    return nil
}

// ProductId Getter
func (r AliexpressPostproductRedefiningFindaeproductbyidfordropshipperRequest) GetProductId() int64 {
    return r._productId
}
// LocalCountry Setter
// 国家
func (r *AliexpressPostproductRedefiningFindaeproductbyidfordropshipperRequest) SetLocalCountry(_localCountry string) error {
    r._localCountry = _localCountry
    r.Set("local_country", _localCountry)
    return nil
}

// LocalCountry Getter
func (r AliexpressPostproductRedefiningFindaeproductbyidfordropshipperRequest) GetLocalCountry() string {
    return r._localCountry
}
// LocalLanguage Setter
// 语言
func (r *AliexpressPostproductRedefiningFindaeproductbyidfordropshipperRequest) SetLocalLanguage(_localLanguage string) error {
    r._localLanguage = _localLanguage
    r.Set("local_language", _localLanguage)
    return nil
}

// LocalLanguage Getter
func (r AliexpressPostproductRedefiningFindaeproductbyidfordropshipperRequest) GetLocalLanguage() string {
    return r._localLanguage
}
