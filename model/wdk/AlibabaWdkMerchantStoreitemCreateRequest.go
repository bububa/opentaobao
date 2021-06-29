package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
新建门店商品 API请求
alibaba.wdk.merchant.storeitem.create

新建门店商品
*/
type AlibabaWdkMerchantStoreitemCreateRequest struct {
    model.Params
    // 门店id
    storeId   string
    // 商品编码
    skuCode   string
    // 商家编码
    merchantCode   string
    // 新增门店商品参数，具体字段详见文档
    params   string
}

// 初始化AlibabaWdkMerchantStoreitemCreateRequest对象
func NewAlibabaWdkMerchantStoreitemCreateRequest() *AlibabaWdkMerchantStoreitemCreateRequest{
    return &AlibabaWdkMerchantStoreitemCreateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkMerchantStoreitemCreateRequest) GetApiMethodName() string {
    return "alibaba.wdk.merchant.storeitem.create"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkMerchantStoreitemCreateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// StoreId Setter
// 门店id
func (r *AlibabaWdkMerchantStoreitemCreateRequest) SetStoreId(storeId string) error {
    r.storeId = storeId
    r.Set("store_id", storeId)
    return nil
}

// StoreId Getter
func (r AlibabaWdkMerchantStoreitemCreateRequest) GetStoreId() string {
    return r.storeId
}
// SkuCode Setter
// 商品编码
func (r *AlibabaWdkMerchantStoreitemCreateRequest) SetSkuCode(skuCode string) error {
    r.skuCode = skuCode
    r.Set("sku_code", skuCode)
    return nil
}

// SkuCode Getter
func (r AlibabaWdkMerchantStoreitemCreateRequest) GetSkuCode() string {
    return r.skuCode
}
// MerchantCode Setter
// 商家编码
func (r *AlibabaWdkMerchantStoreitemCreateRequest) SetMerchantCode(merchantCode string) error {
    r.merchantCode = merchantCode
    r.Set("merchant_code", merchantCode)
    return nil
}

// MerchantCode Getter
func (r AlibabaWdkMerchantStoreitemCreateRequest) GetMerchantCode() string {
    return r.merchantCode
}
// Params Setter
// 新增门店商品参数，具体字段详见文档
func (r *AlibabaWdkMerchantStoreitemCreateRequest) SetParams(params string) error {
    r.params = params
    r.Set("params", params)
    return nil
}

// Params Getter
func (r AlibabaWdkMerchantStoreitemCreateRequest) GetParams() string {
    return r.params
}
