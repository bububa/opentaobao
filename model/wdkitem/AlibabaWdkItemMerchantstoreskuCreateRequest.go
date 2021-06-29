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
    _storeId   string
    // 商品编码
    _skuCode   string
    // 新建参数，是个json字符串
    _params   string
    // 机构
    _orgCode   string
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
func (r *AlibabaWdkItemMerchantstoreskuCreateRequest) SetStoreId(_storeId string) error {
    r._storeId = _storeId
    r.Set("store_id", _storeId)
    return nil
}

// StoreId Getter
func (r AlibabaWdkItemMerchantstoreskuCreateRequest) GetStoreId() string {
    return r._storeId
}
// SkuCode Setter
// 商品编码
func (r *AlibabaWdkItemMerchantstoreskuCreateRequest) SetSkuCode(_skuCode string) error {
    r._skuCode = _skuCode
    r.Set("sku_code", _skuCode)
    return nil
}

// SkuCode Getter
func (r AlibabaWdkItemMerchantstoreskuCreateRequest) GetSkuCode() string {
    return r._skuCode
}
// Params Setter
// 新建参数，是个json字符串
func (r *AlibabaWdkItemMerchantstoreskuCreateRequest) SetParams(_params string) error {
    r._params = _params
    r.Set("params", _params)
    return nil
}

// Params Getter
func (r AlibabaWdkItemMerchantstoreskuCreateRequest) GetParams() string {
    return r._params
}
// OrgCode Setter
// 机构
func (r *AlibabaWdkItemMerchantstoreskuCreateRequest) SetOrgCode(_orgCode string) error {
    r._orgCode = _orgCode
    r.Set("org_code", _orgCode)
    return nil
}

// OrgCode Getter
func (r AlibabaWdkItemMerchantstoreskuCreateRequest) GetOrgCode() string {
    return r._orgCode
}
