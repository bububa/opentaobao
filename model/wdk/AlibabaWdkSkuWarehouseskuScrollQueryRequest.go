package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
仓商品遍历接口(游标) API请求
alibaba.wdk.sku.warehousesku.scroll.query

提供仓商品数据接口查询
*/
type AlibabaWdkSkuWarehouseskuScrollQueryAPIRequest struct {
    model.Params
    // 仓库编码
    _warehouseCode   string
    // 游标
    _scrollId   string
}

// 初始化AlibabaWdkSkuWarehouseskuScrollQueryAPIRequest对象
func NewAlibabaWdkSkuWarehouseskuScrollQueryRequest() *AlibabaWdkSkuWarehouseskuScrollQueryAPIRequest{
    return &AlibabaWdkSkuWarehouseskuScrollQueryAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkSkuWarehouseskuScrollQueryAPIRequest) GetApiMethodName() string {
    return "alibaba.wdk.sku.warehousesku.scroll.query"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkSkuWarehouseskuScrollQueryAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// WarehouseCode Setter
// 仓库编码
func (r *AlibabaWdkSkuWarehouseskuScrollQueryAPIRequest) SetWarehouseCode(_warehouseCode string) error {
    r._warehouseCode = _warehouseCode
    r.Set("warehouse_code", _warehouseCode)
    return nil
}

// WarehouseCode Getter
func (r AlibabaWdkSkuWarehouseskuScrollQueryAPIRequest) GetWarehouseCode() string {
    return r._warehouseCode
}
// ScrollId Setter
// 游标
func (r *AlibabaWdkSkuWarehouseskuScrollQueryAPIRequest) SetScrollId(_scrollId string) error {
    r._scrollId = _scrollId
    r.Set("scroll_id", _scrollId)
    return nil
}

// ScrollId Getter
func (r AlibabaWdkSkuWarehouseskuScrollQueryAPIRequest) GetScrollId() string {
    return r._scrollId
}
