package wdk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkSkuWarehouseskuQueryAPIRequest 仓商品查询接口(指定商品编码) API请求
// alibaba.wdk.sku.warehousesku.query
//
// 提供指定仓商品编码查询
type AlibabaWdkSkuWarehouseskuQueryAPIRequest struct {
	model.Params
	// 商品编码
	_skuCodeList []string
	// 仓编码
	_warehouseCode string
}

// NewAlibabaWdkSkuWarehouseskuQueryRequest 初始化AlibabaWdkSkuWarehouseskuQueryAPIRequest对象
func NewAlibabaWdkSkuWarehouseskuQueryRequest() *AlibabaWdkSkuWarehouseskuQueryAPIRequest {
	return &AlibabaWdkSkuWarehouseskuQueryAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaWdkSkuWarehouseskuQueryAPIRequest) Reset() {
	r._skuCodeList = r._skuCodeList[:0]
	r._warehouseCode = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkSkuWarehouseskuQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.sku.warehousesku.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkSkuWarehouseskuQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaWdkSkuWarehouseskuQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSkuCodeList is SkuCodeList Setter
// 商品编码
func (r *AlibabaWdkSkuWarehouseskuQueryAPIRequest) SetSkuCodeList(_skuCodeList []string) error {
	r._skuCodeList = _skuCodeList
	r.Set("sku_code_list", _skuCodeList)
	return nil
}

// GetSkuCodeList SkuCodeList Getter
func (r AlibabaWdkSkuWarehouseskuQueryAPIRequest) GetSkuCodeList() []string {
	return r._skuCodeList
}

// SetWarehouseCode is WarehouseCode Setter
// 仓编码
func (r *AlibabaWdkSkuWarehouseskuQueryAPIRequest) SetWarehouseCode(_warehouseCode string) error {
	r._warehouseCode = _warehouseCode
	r.Set("warehouse_code", _warehouseCode)
	return nil
}

// GetWarehouseCode WarehouseCode Getter
func (r AlibabaWdkSkuWarehouseskuQueryAPIRequest) GetWarehouseCode() string {
	return r._warehouseCode
}

var poolAlibabaWdkSkuWarehouseskuQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaWdkSkuWarehouseskuQueryRequest()
	},
}

// GetAlibabaWdkSkuWarehouseskuQueryRequest 从 sync.Pool 获取 AlibabaWdkSkuWarehouseskuQueryAPIRequest
func GetAlibabaWdkSkuWarehouseskuQueryAPIRequest() *AlibabaWdkSkuWarehouseskuQueryAPIRequest {
	return poolAlibabaWdkSkuWarehouseskuQueryAPIRequest.Get().(*AlibabaWdkSkuWarehouseskuQueryAPIRequest)
}

// ReleaseAlibabaWdkSkuWarehouseskuQueryAPIRequest 将 AlibabaWdkSkuWarehouseskuQueryAPIRequest 放入 sync.Pool
func ReleaseAlibabaWdkSkuWarehouseskuQueryAPIRequest(v *AlibabaWdkSkuWarehouseskuQueryAPIRequest) {
	v.Reset()
	poolAlibabaWdkSkuWarehouseskuQueryAPIRequest.Put(v)
}
