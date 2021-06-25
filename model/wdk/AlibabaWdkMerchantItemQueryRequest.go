package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商家商品查询 APIRequest
alibaba.wdk.merchant.item.query

商家商品查询
*/
type AlibabaWdkMerchantItemQueryRequest struct {
    model.Params

    // 商品编码
    skuCode   string 

    // 商家编码
    merchantCode   string 

}

func NewAlibabaWdkMerchantItemQueryRequest() *AlibabaWdkMerchantItemQueryRequest{
    return &AlibabaWdkMerchantItemQueryRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaWdkMerchantItemQueryRequest) GetApiMethodName() string {
    return "alibaba.wdk.merchant.item.query"
}

func (r AlibabaWdkMerchantItemQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaWdkMerchantItemQueryRequest) SetSkuCode(skuCode string) error {
    r.skuCode = skuCode
    r.Set("sku_code", skuCode)
    return nil
}

func (r AlibabaWdkMerchantItemQueryRequest) GetSkuCode() string {
    return r.skuCode
}

func (r *AlibabaWdkMerchantItemQueryRequest) SetMerchantCode(merchantCode string) error {
    r.merchantCode = merchantCode
    r.Set("merchant_code", merchantCode)
    return nil
}

func (r AlibabaWdkMerchantItemQueryRequest) GetMerchantCode() string {
    return r.merchantCode
}

