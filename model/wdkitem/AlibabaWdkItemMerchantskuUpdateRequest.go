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
    skuCode   string
    // 参数json
    params   string
    // 机构编码
    orgCode   string
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
func (r *AlibabaWdkItemMerchantskuUpdateRequest) SetSkuCode(skuCode string) error {
    r.skuCode = skuCode
    r.Set("sku_code", skuCode)
    return nil
}

// SkuCode Getter
func (r AlibabaWdkItemMerchantskuUpdateRequest) GetSkuCode() string {
    return r.skuCode
}
// Params Setter
// 参数json
func (r *AlibabaWdkItemMerchantskuUpdateRequest) SetParams(params string) error {
    r.params = params
    r.Set("params", params)
    return nil
}

// Params Getter
func (r AlibabaWdkItemMerchantskuUpdateRequest) GetParams() string {
    return r.params
}
// OrgCode Setter
// 机构编码
func (r *AlibabaWdkItemMerchantskuUpdateRequest) SetOrgCode(orgCode string) error {
    r.orgCode = orgCode
    r.Set("org_code", orgCode)
    return nil
}

// OrgCode Getter
func (r AlibabaWdkItemMerchantskuUpdateRequest) GetOrgCode() string {
    return r.orgCode
}
