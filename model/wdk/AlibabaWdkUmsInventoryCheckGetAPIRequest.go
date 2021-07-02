package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkUmsInventoryCheckGetAPIRequest 盘点结果单-回流单 API请求
// alibaba.wdk.ums.inventory.check.get
//
// 盘点结果单-回流单
type AlibabaWdkUmsInventoryCheckGetAPIRequest struct {
	model.Params
	// 店仓code，指的是库调对象，对应一个物理店或仓编码
	_warehouseCode string
}

// NewAlibabaWdkUmsInventoryCheckGetRequest 初始化AlibabaWdkUmsInventoryCheckGetAPIRequest对象
func NewAlibabaWdkUmsInventoryCheckGetRequest() *AlibabaWdkUmsInventoryCheckGetAPIRequest {
	return &AlibabaWdkUmsInventoryCheckGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkUmsInventoryCheckGetAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.ums.inventory.check.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkUmsInventoryCheckGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is WarehouseCode Setter
// 店仓code，指的是库调对象，对应一个物理店或仓编码
func (r *AlibabaWdkUmsInventoryCheckGetAPIRequest) SetWarehouseCode(_warehouseCode string) error {
	r._warehouseCode = _warehouseCode
	r.Set("warehouse_code", _warehouseCode)
	return nil
}

// Get WarehouseCode Getter
func (r AlibabaWdkUmsInventoryCheckGetAPIRequest) GetWarehouseCode() string {
	return r._warehouseCode
}
