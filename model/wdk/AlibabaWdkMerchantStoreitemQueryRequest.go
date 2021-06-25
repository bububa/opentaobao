package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
门店商品信心查询 APIRequest
alibaba.wdk.merchant.storeitem.query

门店商品信心查询
*/
type AlibabaWdkMerchantStoreitemQueryRequest struct {
    model.Params

    // 商品编码
    skuCode   string 

    // 商家编码
    merchantCode   string 

    // 门店编码
    storeId   string 

}

func NewAlibabaWdkMerchantStoreitemQueryRequest() *AlibabaWdkMerchantStoreitemQueryRequest{
    return &AlibabaWdkMerchantStoreitemQueryRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaWdkMerchantStoreitemQueryRequest) GetApiMethodName() string {
    return "alibaba.wdk.merchant.storeitem.query"
}

func (r AlibabaWdkMerchantStoreitemQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaWdkMerchantStoreitemQueryRequest) SetSkuCode(skuCode string) error {
    r.skuCode = skuCode
    r.Set("sku_code", skuCode)
    return nil
}

func (r AlibabaWdkMerchantStoreitemQueryRequest) GetSkuCode() string {
    return r.skuCode
}

func (r *AlibabaWdkMerchantStoreitemQueryRequest) SetMerchantCode(merchantCode string) error {
    r.merchantCode = merchantCode
    r.Set("merchant_code", merchantCode)
    return nil
}

func (r AlibabaWdkMerchantStoreitemQueryRequest) GetMerchantCode() string {
    return r.merchantCode
}

func (r *AlibabaWdkMerchantStoreitemQueryRequest) SetStoreId(storeId string) error {
    r.storeId = storeId
    r.Set("store_id", storeId)
    return nil
}

func (r AlibabaWdkMerchantStoreitemQueryRequest) GetStoreId() string {
    return r.storeId
}

