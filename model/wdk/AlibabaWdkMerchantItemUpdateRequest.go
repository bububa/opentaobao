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
    skuCode   string
    // 门店编码
    merchantCode   string
    // 修改字段的json
    params   string
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
func (r *AlibabaWdkMerchantItemUpdateRequest) SetSkuCode(skuCode string) error {
    r.skuCode = skuCode
    r.Set("sku_code", skuCode)
    return nil
}

// SkuCode Getter
func (r AlibabaWdkMerchantItemUpdateRequest) GetSkuCode() string {
    return r.skuCode
}
// MerchantCode Setter
// 门店编码
func (r *AlibabaWdkMerchantItemUpdateRequest) SetMerchantCode(merchantCode string) error {
    r.merchantCode = merchantCode
    r.Set("merchant_code", merchantCode)
    return nil
}

// MerchantCode Getter
func (r AlibabaWdkMerchantItemUpdateRequest) GetMerchantCode() string {
    return r.merchantCode
}
// Params Setter
// 修改字段的json
func (r *AlibabaWdkMerchantItemUpdateRequest) SetParams(params string) error {
    r.params = params
    r.Set("params", params)
    return nil
}

// Params Getter
func (r AlibabaWdkMerchantItemUpdateRequest) GetParams() string {
    return r.params
}
