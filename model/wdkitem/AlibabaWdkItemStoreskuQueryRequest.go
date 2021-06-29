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
    skuCode   string
    // 门店ID
    storeId   string
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
func (r *AlibabaWdkItemStoreskuQueryRequest) SetSkuCode(skuCode string) error {
    r.skuCode = skuCode
    r.Set("sku_code", skuCode)
    return nil
}

// SkuCode Getter
func (r AlibabaWdkItemStoreskuQueryRequest) GetSkuCode() string {
    return r.skuCode
}
// StoreId Setter
// 门店ID
func (r *AlibabaWdkItemStoreskuQueryRequest) SetStoreId(storeId string) error {
    r.storeId = storeId
    r.Set("store_id", storeId)
    return nil
}

// StoreId Getter
func (r AlibabaWdkItemStoreskuQueryRequest) GetStoreId() string {
    return r.storeId
}
