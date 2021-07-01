package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
门店商品信心查询 API请求
alibaba.wdk.merchant.storeitem.query

门店商品信心查询
*/
type AlibabaWdkMerchantStoreitemQueryAPIRequest struct {
    model.Params
    // 商品编码
    _skuCode   string
    // 商家编码
    _merchantCode   string
    // 门店编码
    _storeId   string
}

// 初始化AlibabaWdkMerchantStoreitemQueryAPIRequest对象
func NewAlibabaWdkMerchantStoreitemQueryRequest() *AlibabaWdkMerchantStoreitemQueryAPIRequest{
    return &AlibabaWdkMerchantStoreitemQueryAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkMerchantStoreitemQueryAPIRequest) GetApiMethodName() string {
    return "alibaba.wdk.merchant.storeitem.query"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkMerchantStoreitemQueryAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// SkuCode Setter
// 商品编码
func (r *AlibabaWdkMerchantStoreitemQueryAPIRequest) SetSkuCode(_skuCode string) error {
    r._skuCode = _skuCode
    r.Set("sku_code", _skuCode)
    return nil
}

// SkuCode Getter
func (r AlibabaWdkMerchantStoreitemQueryAPIRequest) GetSkuCode() string {
    return r._skuCode
}
// MerchantCode Setter
// 商家编码
func (r *AlibabaWdkMerchantStoreitemQueryAPIRequest) SetMerchantCode(_merchantCode string) error {
    r._merchantCode = _merchantCode
    r.Set("merchant_code", _merchantCode)
    return nil
}

// MerchantCode Getter
func (r AlibabaWdkMerchantStoreitemQueryAPIRequest) GetMerchantCode() string {
    return r._merchantCode
}
// StoreId Setter
// 门店编码
func (r *AlibabaWdkMerchantStoreitemQueryAPIRequest) SetStoreId(_storeId string) error {
    r._storeId = _storeId
    r.Set("store_id", _storeId)
    return nil
}

// StoreId Getter
func (r AlibabaWdkMerchantStoreitemQueryAPIRequest) GetStoreId() string {
    return r._storeId
}
