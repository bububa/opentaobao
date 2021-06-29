package wdkitem

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商家商品信息查询 API请求
alibaba.wdk.item.merchantsku.query

商家商品信息查询
*/
type AlibabaWdkItemMerchantskuQueryRequest struct {
    model.Params
    // 商品编码
    skuCode   string
    // 机构编码
    orgCode   string
}

// 初始化AlibabaWdkItemMerchantskuQueryRequest对象
func NewAlibabaWdkItemMerchantskuQueryRequest() *AlibabaWdkItemMerchantskuQueryRequest{
    return &AlibabaWdkItemMerchantskuQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkItemMerchantskuQueryRequest) GetApiMethodName() string {
    return "alibaba.wdk.item.merchantsku.query"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkItemMerchantskuQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// SkuCode Setter
// 商品编码
func (r *AlibabaWdkItemMerchantskuQueryRequest) SetSkuCode(skuCode string) error {
    r.skuCode = skuCode
    r.Set("sku_code", skuCode)
    return nil
}

// SkuCode Getter
func (r AlibabaWdkItemMerchantskuQueryRequest) GetSkuCode() string {
    return r.skuCode
}
// OrgCode Setter
// 机构编码
func (r *AlibabaWdkItemMerchantskuQueryRequest) SetOrgCode(orgCode string) error {
    r.orgCode = orgCode
    r.Set("org_code", orgCode)
    return nil
}

// OrgCode Getter
func (r AlibabaWdkItemMerchantskuQueryRequest) GetOrgCode() string {
    return r.orgCode
}
