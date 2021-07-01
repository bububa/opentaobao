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
type AlibabaWdkItemMerchantskuQueryAPIRequest struct {
    model.Params
    // 商品编码
    _skuCode   string
    // 机构编码
    _orgCode   string
}

// 初始化AlibabaWdkItemMerchantskuQueryAPIRequest对象
func NewAlibabaWdkItemMerchantskuQueryRequest() *AlibabaWdkItemMerchantskuQueryAPIRequest{
    return &AlibabaWdkItemMerchantskuQueryAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkItemMerchantskuQueryAPIRequest) GetApiMethodName() string {
    return "alibaba.wdk.item.merchantsku.query"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkItemMerchantskuQueryAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// SkuCode Setter
// 商品编码
func (r *AlibabaWdkItemMerchantskuQueryAPIRequest) SetSkuCode(_skuCode string) error {
    r._skuCode = _skuCode
    r.Set("sku_code", _skuCode)
    return nil
}

// SkuCode Getter
func (r AlibabaWdkItemMerchantskuQueryAPIRequest) GetSkuCode() string {
    return r._skuCode
}
// OrgCode Setter
// 机构编码
func (r *AlibabaWdkItemMerchantskuQueryAPIRequest) SetOrgCode(_orgCode string) error {
    r._orgCode = _orgCode
    r.Set("org_code", _orgCode)
    return nil
}

// OrgCode Getter
func (r AlibabaWdkItemMerchantskuQueryAPIRequest) GetOrgCode() string {
    return r._orgCode
}
