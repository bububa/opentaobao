package wdkitem

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
门店商品信息查询 API请求
alibaba.wdk.item.storesku.query

门店商品信息查询
*/
type AlibabaWdkItemStoreskuQueryRequest struct {
    model.Params
    // 商品编码
    _skuCode   string
    // 门店ID
    _storeId   string
}

// 初始化AlibabaWdkItemStoreskuQueryRequest对象
func NewAlibabaWdkItemStoreskuQueryRequest() *AlibabaWdkItemStoreskuQueryRequest{
    return &AlibabaWdkItemStoreskuQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkItemStoreskuQueryRequest) GetApiMethodName() string {
    return "alibaba.wdk.item.storesku.query"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkItemStoreskuQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// SkuCode Setter
// 商品编码
func (r *AlibabaWdkItemStoreskuQueryRequest) SetSkuCode(_skuCode string) error {
    r._skuCode = _skuCode
    r.Set("sku_code", _skuCode)
    return nil
}

// SkuCode Getter
func (r AlibabaWdkItemStoreskuQueryRequest) GetSkuCode() string {
    return r._skuCode
}
// StoreId Setter
// 门店ID
func (r *AlibabaWdkItemStoreskuQueryRequest) SetStoreId(_storeId string) error {
    r._storeId = _storeId
    r.Set("store_id", _storeId)
    return nil
}

// StoreId Getter
func (r AlibabaWdkItemStoreskuQueryRequest) GetStoreId() string {
    return r._storeId
}
