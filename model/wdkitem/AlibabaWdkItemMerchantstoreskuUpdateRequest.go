package wdkitem

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
门店商品信息修改 API请求
alibaba.wdk.item.merchantstoresku.update

门店商品信息修改
*/
type AlibabaWdkItemMerchantstoreskuUpdateRequest struct {
    model.Params
    // 门店ID
    _storeId   string
    // 商品编码
    _skuCode   string
    // 修改参数，是个json字符串
    _params   string
}

// 初始化AlibabaWdkItemMerchantstoreskuUpdateRequest对象
func NewAlibabaWdkItemMerchantstoreskuUpdateRequest() *AlibabaWdkItemMerchantstoreskuUpdateRequest{
    return &AlibabaWdkItemMerchantstoreskuUpdateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkItemMerchantstoreskuUpdateRequest) GetApiMethodName() string {
    return "alibaba.wdk.item.merchantstoresku.update"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkItemMerchantstoreskuUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// StoreId Setter
// 门店ID
func (r *AlibabaWdkItemMerchantstoreskuUpdateRequest) SetStoreId(_storeId string) error {
    r._storeId = _storeId
    r.Set("store_id", _storeId)
    return nil
}

// StoreId Getter
func (r AlibabaWdkItemMerchantstoreskuUpdateRequest) GetStoreId() string {
    return r._storeId
}
// SkuCode Setter
// 商品编码
func (r *AlibabaWdkItemMerchantstoreskuUpdateRequest) SetSkuCode(_skuCode string) error {
    r._skuCode = _skuCode
    r.Set("sku_code", _skuCode)
    return nil
}

// SkuCode Getter
func (r AlibabaWdkItemMerchantstoreskuUpdateRequest) GetSkuCode() string {
    return r._skuCode
}
// Params Setter
// 修改参数，是个json字符串
func (r *AlibabaWdkItemMerchantstoreskuUpdateRequest) SetParams(_params string) error {
    r._params = _params
    r.Set("params", _params)
    return nil
}

// Params Getter
func (r AlibabaWdkItemMerchantstoreskuUpdateRequest) GetParams() string {
    return r._params
}
