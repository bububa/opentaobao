package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
修改商家商品 APIRequest
alibaba.wdk.merchant.item.update

修改商家商品
*/
type AlibabaWdkMerchantItemUpdateRequest struct {
    model.Params

    // 商品编码
    skuCode   string 

    // 门店编码
    merchantCode   string 

    // 修改字段的json
    params   string 

}

func NewAlibabaWdkMerchantItemUpdateRequest() *AlibabaWdkMerchantItemUpdateRequest{
    return &AlibabaWdkMerchantItemUpdateRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaWdkMerchantItemUpdateRequest) GetApiMethodName() string {
    return "alibaba.wdk.merchant.item.update"
}

func (r AlibabaWdkMerchantItemUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaWdkMerchantItemUpdateRequest) SetSkuCode(skuCode string) error {
    r.skuCode = skuCode
    r.Set("sku_code", skuCode)
    return nil
}

func (r AlibabaWdkMerchantItemUpdateRequest) GetSkuCode() string {
    return r.skuCode
}

func (r *AlibabaWdkMerchantItemUpdateRequest) SetMerchantCode(merchantCode string) error {
    r.merchantCode = merchantCode
    r.Set("merchant_code", merchantCode)
    return nil
}

func (r AlibabaWdkMerchantItemUpdateRequest) GetMerchantCode() string {
    return r.merchantCode
}

func (r *AlibabaWdkMerchantItemUpdateRequest) SetParams(params string) error {
    r.params = params
    r.Set("params", params)
    return nil
}

func (r AlibabaWdkMerchantItemUpdateRequest) GetParams() string {
    return r.params
}

