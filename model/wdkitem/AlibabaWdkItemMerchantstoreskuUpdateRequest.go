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
    storeId   string
    // 商品编码
    skuCode   string
    // 修改参数，是个json字符串
    params   string
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
func (r *AlibabaWdkItemMerchantstoreskuUpdateRequest) SetStoreId(storeId string) error {
    r.storeId = storeId
    r.Set("store_id", storeId)
    return nil
}

// StoreId Getter
func (r AlibabaWdkItemMerchantstoreskuUpdateRequest) GetStoreId() string {
    return r.storeId
}
// SkuCode Setter
// 商品编码
func (r *AlibabaWdkItemMerchantstoreskuUpdateRequest) SetSkuCode(skuCode string) error {
    r.skuCode = skuCode
    r.Set("sku_code", skuCode)
    return nil
}

// SkuCode Getter
func (r AlibabaWdkItemMerchantstoreskuUpdateRequest) GetSkuCode() string {
    return r.skuCode
}
// Params Setter
// 修改参数，是个json字符串
func (r *AlibabaWdkItemMerchantstoreskuUpdateRequest) SetParams(params string) error {
    r.params = params
    r.Set("params", params)
    return nil
}

// Params Getter
func (r AlibabaWdkItemMerchantstoreskuUpdateRequest) GetParams() string {
    return r.params
}
