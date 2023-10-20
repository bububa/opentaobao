package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkskuwarehouseskuqueryAPIRequest 仓商品查询接口(指定商品编码) API请求
// alibaba.wdk.sku.warehousesku.query
//
// 提供指定仓商品编码查询
type AlibabawdkskuwarehouseskuqueryAPIRequest struct {
	model.Params
	// 商品编码
	_skuCodeList []string
	// 仓编码
	_warehouseCode string
}

// NewAlibabawdkskuwarehouseskuqueryRequest 初始化AlibabawdkskuwarehouseskuqueryAPIRequest对象
func NewAlibabawdkskuwarehouseskuqueryRequest() *AlibabawdkskuwarehouseskuqueryAPIRequest {
	return &AlibabawdkskuwarehouseskuqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabawdkskuwarehouseskuqueryAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.sku.warehousesku.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabawdkskuwarehouseskuqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabawdkskuwarehouseskuqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSkuCodeList is SkuCodeList Setter
// 商品编码
func (r *AlibabawdkskuwarehouseskuqueryAPIRequest) SetSkuCodeList(_skuCodeList []string) error {
	r._skuCodeList = _skuCodeList
	r.Set("sku_code_list", _skuCodeList)
	return nil
}

// GetSkuCodeList SkuCodeList Getter
func (r AlibabawdkskuwarehouseskuqueryAPIRequest) GetSkuCodeList() []string {
	return r._skuCodeList
}

// SetWarehouseCode is WarehouseCode Setter
// 仓编码
func (r *AlibabawdkskuwarehouseskuqueryAPIRequest) SetWarehouseCode(_warehouseCode string) error {
	r._warehouseCode = _warehouseCode
	r.Set("warehouse_code", _warehouseCode)
	return nil
}

// GetWarehouseCode WarehouseCode Getter
func (r AlibabawdkskuwarehouseskuqueryAPIRequest) GetWarehouseCode() string {
	return r._warehouseCode
}
