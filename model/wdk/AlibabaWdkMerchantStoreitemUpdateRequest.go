package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
修改门店商品 API请求
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

// 初始化AlibabaWdkMerchantStoreitemUpdateRequest对象
func NewAlibabaWdkMerchantStoreitemUpdateRequest() *AlibabaWdkMerchantStoreitemUpdateRequest{
    return &AlibabaWdkMerchantStoreitemUpdateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkMerchantStoreitemUpdateRequest) GetApiMethodName() string {
    return "alibaba.wdk.merchant.storeitem.update"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkMerchantStoreitemUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// SkuCode Setter
// 商品编码
func (r *AlibabaWdkMerchantStoreitemUpdateRequest) SetSkuCode(skuCode string) error {
    r.skuCode = skuCode
    r.Set("sku_code", skuCode)
    return nil
}

// SkuCode Getter
func (r AlibabaWdkMerchantStoreitemUpdateRequest) GetSkuCode() string {
    return r.skuCode
}
// MerchantCode Setter
// 商家编码
func (r *AlibabaWdkMerchantStoreitemUpdateRequest) SetMerchantCode(merchantCode string) error {
    r.merchantCode = merchantCode
    r.Set("merchant_code", merchantCode)
    return nil
}

// MerchantCode Getter
func (r AlibabaWdkMerchantStoreitemUpdateRequest) GetMerchantCode() string {
    return r.merchantCode
}
// StoreId Setter
// 门店编码
func (r *AlibabaWdkMerchantStoreitemUpdateRequest) SetStoreId(storeId string) error {
    r.storeId = storeId
    r.Set("store_id", storeId)
    return nil
}

// StoreId Getter
func (r AlibabaWdkMerchantStoreitemUpdateRequest) GetStoreId() string {
    return r.storeId
}
// Params Setter
// 修改参数的json
func (r *AlibabaWdkMerchantStoreitemUpdateRequest) SetParams(params string) error {
    r.params = params
    r.Set("params", params)
    return nil
}

// Params Getter
func (r AlibabaWdkMerchantStoreitemUpdateRequest) GetParams() string {
    return r.params
}
