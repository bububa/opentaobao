package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
仓商品查询接口(指定商品编码) API请求
alibaba.wdk.sku.warehousesku.query

提供指定仓商品编码查询
*/
type AlibabaWdkSkuWarehouseskuQueryAPIRequest struct {
    model.Params
    // 商品编码
    _skuCodeList   []string
    // 仓编码
    _warehouseCode   string
}

// 初始化AlibabaWdkSkuWarehouseskuQueryAPIRequest对象
func NewAlibabaWdkSkuWarehouseskuQueryRequest() *AlibabaWdkSkuWarehouseskuQueryAPIRequest{
    return &AlibabaWdkSkuWarehouseskuQueryAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkSkuWarehouseskuQueryAPIRequest) GetApiMethodName() string {
    return "alibaba.wdk.sku.warehousesku.query"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkSkuWarehouseskuQueryAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// SkuCodeList Setter
// 商品编码
func (r *AlibabaWdkSkuWarehouseskuQueryAPIRequest) SetSkuCodeList(_skuCodeList []string) error {
    r._skuCodeList = _skuCodeList
    r.Set("sku_code_list", _skuCodeList)
    return nil
}

// SkuCodeList Getter
func (r AlibabaWdkSkuWarehouseskuQueryAPIRequest) GetSkuCodeList() []string {
    return r._skuCodeList
}
// WarehouseCode Setter
// 仓编码
func (r *AlibabaWdkSkuWarehouseskuQueryAPIRequest) SetWarehouseCode(_warehouseCode string) error {
    r._warehouseCode = _warehouseCode
    r.Set("warehouse_code", _warehouseCode)
    return nil
}

// WarehouseCode Getter
func (r AlibabaWdkSkuWarehouseskuQueryAPIRequest) GetWarehouseCode() string {
    return r._warehouseCode
}
