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
type AlibabaWdkMerchantStoreitemUpdateAPIRequest struct {
    model.Params
    // 商品编码
    _skuCode   string
    // 商家编码
    _merchantCode   string
    // 门店编码
    _storeId   string
    // 修改参数的json
    _params   string
}

// 初始化AlibabaWdkMerchantStoreitemUpdateAPIRequest对象
func NewAlibabaWdkMerchantStoreitemUpdateRequest() *AlibabaWdkMerchantStoreitemUpdateAPIRequest{
    return &AlibabaWdkMerchantStoreitemUpdateAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkMerchantStoreitemUpdateAPIRequest) GetApiMethodName() string {
    return "alibaba.wdk.merchant.storeitem.update"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkMerchantStoreitemUpdateAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// SkuCode Setter
// 商品编码
func (r *AlibabaWdkMerchantStoreitemUpdateAPIRequest) SetSkuCode(_skuCode string) error {
    r._skuCode = _skuCode
    r.Set("sku_code", _skuCode)
    return nil
}

// SkuCode Getter
func (r AlibabaWdkMerchantStoreitemUpdateAPIRequest) GetSkuCode() string {
    return r._skuCode
}
// MerchantCode Setter
// 商家编码
func (r *AlibabaWdkMerchantStoreitemUpdateAPIRequest) SetMerchantCode(_merchantCode string) error {
    r._merchantCode = _merchantCode
    r.Set("merchant_code", _merchantCode)
    return nil
}

// MerchantCode Getter
func (r AlibabaWdkMerchantStoreitemUpdateAPIRequest) GetMerchantCode() string {
    return r._merchantCode
}
// StoreId Setter
// 门店编码
func (r *AlibabaWdkMerchantStoreitemUpdateAPIRequest) SetStoreId(_storeId string) error {
    r._storeId = _storeId
    r.Set("store_id", _storeId)
    return nil
}

// StoreId Getter
func (r AlibabaWdkMerchantStoreitemUpdateAPIRequest) GetStoreId() string {
    return r._storeId
}
// Params Setter
// 修改参数的json
func (r *AlibabaWdkMerchantStoreitemUpdateAPIRequest) SetParams(_params string) error {
    r._params = _params
    r.Set("params", _params)
    return nil
}

// Params Getter
func (r AlibabaWdkMerchantStoreitemUpdateAPIRequest) GetParams() string {
    return r._params
}
