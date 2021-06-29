package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
修改商家商品 API请求
alibaba.wdk.merchant.item.update

修改商家商品
*/
type AlibabaWdkMerchantItemUpdateRequest struct {
    model.Params
    // 商品编码
    _skuCode   string
    // 门店编码
    _merchantCode   string
    // 修改字段的json
    _params   string
}

// 初始化AlibabaWdkMerchantItemUpdateRequest对象
func NewAlibabaWdkMerchantItemUpdateRequest() *AlibabaWdkMerchantItemUpdateRequest{
    return &AlibabaWdkMerchantItemUpdateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkMerchantItemUpdateRequest) GetApiMethodName() string {
    return "alibaba.wdk.merchant.item.update"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkMerchantItemUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// SkuCode Setter
// 商品编码
func (r *AlibabaWdkMerchantItemUpdateRequest) SetSkuCode(_skuCode string) error {
    r._skuCode = _skuCode
    r.Set("sku_code", _skuCode)
    return nil
}

// SkuCode Getter
func (r AlibabaWdkMerchantItemUpdateRequest) GetSkuCode() string {
    return r._skuCode
}
// MerchantCode Setter
// 门店编码
func (r *AlibabaWdkMerchantItemUpdateRequest) SetMerchantCode(_merchantCode string) error {
    r._merchantCode = _merchantCode
    r.Set("merchant_code", _merchantCode)
    return nil
}

// MerchantCode Getter
func (r AlibabaWdkMerchantItemUpdateRequest) GetMerchantCode() string {
    return r._merchantCode
}
// Params Setter
// 修改字段的json
func (r *AlibabaWdkMerchantItemUpdateRequest) SetParams(_params string) error {
    r._params = _params
    r.Set("params", _params)
    return nil
}

// Params Getter
func (r AlibabaWdkMerchantItemUpdateRequest) GetParams() string {
    return r._params
}
