package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商家商品查询 API请求
alibaba.wdk.merchant.item.query

商家商品查询
*/
type AlibabaWdkMerchantItemQueryRequest struct {
    model.Params
    // 商品编码
    _skuCode   string
    // 商家编码
    _merchantCode   string
}

// 初始化AlibabaWdkMerchantItemQueryRequest对象
func NewAlibabaWdkMerchantItemQueryRequest() *AlibabaWdkMerchantItemQueryRequest{
    return &AlibabaWdkMerchantItemQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkMerchantItemQueryRequest) GetApiMethodName() string {
    return "alibaba.wdk.merchant.item.query"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkMerchantItemQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// SkuCode Setter
// 商品编码
func (r *AlibabaWdkMerchantItemQueryRequest) SetSkuCode(_skuCode string) error {
    r._skuCode = _skuCode
    r.Set("sku_code", _skuCode)
    return nil
}

// SkuCode Getter
func (r AlibabaWdkMerchantItemQueryRequest) GetSkuCode() string {
    return r._skuCode
}
// MerchantCode Setter
// 商家编码
func (r *AlibabaWdkMerchantItemQueryRequest) SetMerchantCode(_merchantCode string) error {
    r._merchantCode = _merchantCode
    r.Set("merchant_code", _merchantCode)
    return nil
}

// MerchantCode Getter
func (r AlibabaWdkMerchantItemQueryRequest) GetMerchantCode() string {
    return r._merchantCode
}
