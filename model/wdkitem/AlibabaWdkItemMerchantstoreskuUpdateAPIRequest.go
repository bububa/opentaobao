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
type AlibabaWdkItemMerchantstoreskuUpdateAPIRequest struct {
    model.Params
    // 门店ID
    _storeId   string
    // 商品编码
    _skuCode   string
    // 修改参数，是个json字符串
    _params   string
}

// 初始化AlibabaWdkItemMerchantstoreskuUpdateAPIRequest对象
func NewAlibabaWdkItemMerchantstoreskuUpdateRequest() *AlibabaWdkItemMerchantstoreskuUpdateAPIRequest{
    return &AlibabaWdkItemMerchantstoreskuUpdateAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkItemMerchantstoreskuUpdateAPIRequest) GetApiMethodName() string {
    return "alibaba.wdk.item.merchantstoresku.update"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkItemMerchantstoreskuUpdateAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// StoreId Setter
// 门店ID
func (r *AlibabaWdkItemMerchantstoreskuUpdateAPIRequest) SetStoreId(_storeId string) error {
    r._storeId = _storeId
    r.Set("store_id", _storeId)
    return nil
}

// StoreId Getter
func (r AlibabaWdkItemMerchantstoreskuUpdateAPIRequest) GetStoreId() string {
    return r._storeId
}
// SkuCode Setter
// 商品编码
func (r *AlibabaWdkItemMerchantstoreskuUpdateAPIRequest) SetSkuCode(_skuCode string) error {
    r._skuCode = _skuCode
    r.Set("sku_code", _skuCode)
    return nil
}

// SkuCode Getter
func (r AlibabaWdkItemMerchantstoreskuUpdateAPIRequest) GetSkuCode() string {
    return r._skuCode
}
// Params Setter
// 修改参数，是个json字符串
func (r *AlibabaWdkItemMerchantstoreskuUpdateAPIRequest) SetParams(_params string) error {
    r._params = _params
    r.Set("params", _params)
    return nil
}

// Params Getter
func (r AlibabaWdkItemMerchantstoreskuUpdateAPIRequest) GetParams() string {
    return r._params
}
