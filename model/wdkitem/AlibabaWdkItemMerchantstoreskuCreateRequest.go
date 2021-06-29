package wdkitem

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
门店商品信息新建 API请求
alibaba.wdk.item.merchantstoresku.create

门店商品信息新建
*/
type AlibabaWdkItemMerchantstoreskuCreateRequest struct {
    model.Params
    // 门店编码
    storeId   string
    // 商品编码
    skuCode   string
    // 新建参数，是个json字符串
    params   string
    // 机构
    orgCode   string
}

// 初始化AlibabaWdkItemMerchantstoreskuCreateRequest对象
func NewAlibabaWdkItemMerchantstoreskuCreateRequest() *AlibabaWdkItemMerchantstoreskuCreateRequest{
    return &AlibabaWdkItemMerchantstoreskuCreateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkItemMerchantstoreskuCreateRequest) GetApiMethodName() string {
    return "alibaba.wdk.item.merchantstoresku.create"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkItemMerchantstoreskuCreateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// StoreId Setter
// 门店编码
func (r *AlibabaWdkItemMerchantstoreskuCreateRequest) SetStoreId(storeId string) error {
    r.storeId = storeId
    r.Set("store_id", storeId)
    return nil
}

// StoreId Getter
func (r AlibabaWdkItemMerchantstoreskuCreateRequest) GetStoreId() string {
    return r.storeId
}
// SkuCode Setter
// 商品编码
func (r *AlibabaWdkItemMerchantstoreskuCreateRequest) SetSkuCode(skuCode string) error {
    r.skuCode = skuCode
    r.Set("sku_code", skuCode)
    return nil
}

// SkuCode Getter
func (r AlibabaWdkItemMerchantstoreskuCreateRequest) GetSkuCode() string {
    return r.skuCode
}
// Params Setter
// 新建参数，是个json字符串
func (r *AlibabaWdkItemMerchantstoreskuCreateRequest) SetParams(params string) error {
    r.params = params
    r.Set("params", params)
    return nil
}

// Params Getter
func (r AlibabaWdkItemMerchantstoreskuCreateRequest) GetParams() string {
    return r.params
}
// OrgCode Setter
// 机构
func (r *AlibabaWdkItemMerchantstoreskuCreateRequest) SetOrgCode(orgCode string) error {
    r.orgCode = orgCode
    r.Set("org_code", orgCode)
    return nil
}

// OrgCode Getter
func (r AlibabaWdkItemMerchantstoreskuCreateRequest) GetOrgCode() string {
    return r.orgCode
}
