package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
修改门店商品 APIRequest
alibaba.wdk.merchant.storeitem.update

修改门店商品
*/
type AlibabaWdkMerchantStoreitemUpdateRequest struct {
    model.Params

    // 商品编码
    skuCode   string 

    // 商家编码
    merchantCode   string 

    // 门店编码
    storeId   string 

    // 修改参数的json
    params   string 

}

func NewAlibabaWdkMerchantStoreitemUpdateRequest() *AlibabaWdkMerchantStoreitemUpdateRequest{
    return &AlibabaWdkMerchantStoreitemUpdateRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaWdkMerchantStoreitemUpdateRequest) GetApiMethodName() string {
    return "alibaba.wdk.merchant.storeitem.update"
}

func (r AlibabaWdkMerchantStoreitemUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaWdkMerchantStoreitemUpdateRequest) SetSkuCode(skuCode string) error {
    r.skuCode = skuCode
    r.Set("sku_code", skuCode)
    return nil
}

func (r AlibabaWdkMerchantStoreitemUpdateRequest) GetSkuCode() string {
    return r.skuCode
}

func (r *AlibabaWdkMerchantStoreitemUpdateRequest) SetMerchantCode(merchantCode string) error {
    r.merchantCode = merchantCode
    r.Set("merchant_code", merchantCode)
    return nil
}

func (r AlibabaWdkMerchantStoreitemUpdateRequest) GetMerchantCode() string {
    return r.merchantCode
}

func (r *AlibabaWdkMerchantStoreitemUpdateRequest) SetStoreId(storeId string) error {
    r.storeId = storeId
    r.Set("store_id", storeId)
    return nil
}

func (r AlibabaWdkMerchantStoreitemUpdateRequest) GetStoreId() string {
    return r.storeId
}

func (r *AlibabaWdkMerchantStoreitemUpdateRequest) SetParams(params string) error {
    r.params = params
    r.Set("params", params)
    return nil
}

func (r AlibabaWdkMerchantStoreitemUpdateRequest) GetParams() string {
    return r.params
}

