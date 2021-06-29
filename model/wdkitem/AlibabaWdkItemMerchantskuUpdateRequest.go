package wdkitem

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商家商品修改 API请求
alibaba.wdk.item.merchantsku.update

商家商品修改
*/
type AlibabaWdkItemMerchantskuUpdateRequest struct {
    model.Params
    // 商品编码
    _skuCode   string
    // 参数json
    _params   string
    // 机构编码
    _orgCode   string
}

// 初始化AlibabaWdkItemMerchantskuUpdateRequest对象
func NewAlibabaWdkItemMerchantskuUpdateRequest() *AlibabaWdkItemMerchantskuUpdateRequest{
    return &AlibabaWdkItemMerchantskuUpdateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkItemMerchantskuUpdateRequest) GetApiMethodName() string {
    return "alibaba.wdk.item.merchantsku.update"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkItemMerchantskuUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// SkuCode Setter
// 商品编码
func (r *AlibabaWdkItemMerchantskuUpdateRequest) SetSkuCode(_skuCode string) error {
    r._skuCode = _skuCode
    r.Set("sku_code", _skuCode)
    return nil
}

// SkuCode Getter
func (r AlibabaWdkItemMerchantskuUpdateRequest) GetSkuCode() string {
    return r._skuCode
}
// Params Setter
// 参数json
func (r *AlibabaWdkItemMerchantskuUpdateRequest) SetParams(_params string) error {
    r._params = _params
    r.Set("params", _params)
    return nil
}

// Params Getter
func (r AlibabaWdkItemMerchantskuUpdateRequest) GetParams() string {
    return r._params
}
// OrgCode Setter
// 机构编码
func (r *AlibabaWdkItemMerchantskuUpdateRequest) SetOrgCode(_orgCode string) error {
    r._orgCode = _orgCode
    r.Set("org_code", _orgCode)
    return nil
}

// OrgCode Getter
func (r AlibabaWdkItemMerchantskuUpdateRequest) GetOrgCode() string {
    return r._orgCode
}
