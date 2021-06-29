package aedropshiper

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
Dropshipper查找商品信息接口 APIRequest
aliexpress.postproduct.redefining.findaeproductbyidfordropshipper

提供给Dropshipper的通过商品ID查找商品信息的接口，只有特定买家可以使用
*/
type AliexpressPostproductRedefiningFindaeproductbyidfordropshipperRequest struct {
    model.Params

    // 商品ID
    productId   int64 

    // 国家
    localCountry   string 

    // 语言
    localLanguage   string 

}

func NewAliexpressPostproductRedefiningFindaeproductbyidfordropshipperRequest() *AliexpressPostproductRedefiningFindaeproductbyidfordropshipperRequest{
    return &AliexpressPostproductRedefiningFindaeproductbyidfordropshipperRequest{
        Params: model.NewParams(),
    }
}

func (r AliexpressPostproductRedefiningFindaeproductbyidfordropshipperRequest) GetApiMethodName() string {
    return "aliexpress.postproduct.redefining.findaeproductbyidfordropshipper"
}

func (r AliexpressPostproductRedefiningFindaeproductbyidfordropshipperRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AliexpressPostproductRedefiningFindaeproductbyidfordropshipperRequest) SetProductId(productId int64) error {
    r.productId = productId
    r.Set("product_id", productId)
    return nil
}

func (r AliexpressPostproductRedefiningFindaeproductbyidfordropshipperRequest) GetProductId() int64 {
    return r.productId
}

func (r *AliexpressPostproductRedefiningFindaeproductbyidfordropshipperRequest) SetLocalCountry(localCountry string) error {
    r.localCountry = localCountry
    r.Set("local_country", localCountry)
    return nil
}

func (r AliexpressPostproductRedefiningFindaeproductbyidfordropshipperRequest) GetLocalCountry() string {
    return r.localCountry
}

func (r *AliexpressPostproductRedefiningFindaeproductbyidfordropshipperRequest) SetLocalLanguage(localLanguage string) error {
    r.localLanguage = localLanguage
    r.Set("local_language", localLanguage)
    return nil
}

func (r AliexpressPostproductRedefiningFindaeproductbyidfordropshipperRequest) GetLocalLanguage() string {
    return r.localLanguage
}

