package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkumsinventoryadjustgetAPIRequest 库调单-回流单 API请求
// alibaba.wdk.ums.inventory.adjust.get
//
// 库调单-回流单
type AlibabawdkumsinventoryadjustgetAPIRequest struct {
	model.Params
	// 店仓code，指的是库调对象，对应一个物理店或仓编码
	_warehouseCode string
}

// NewAlibabawdkumsinventoryadjustgetRequest 初始化AlibabawdkumsinventoryadjustgetAPIRequest对象
func NewAlibabawdkumsinventoryadjustgetRequest() *AlibabawdkumsinventoryadjustgetAPIRequest {
	return &AlibabawdkumsinventoryadjustgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabawdkumsinventoryadjustgetAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.ums.inventory.adjust.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabawdkumsinventoryadjustgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabawdkumsinventoryadjustgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetWarehouseCode is WarehouseCode Setter
// 店仓code，指的是库调对象，对应一个物理店或仓编码
func (r *AlibabawdkumsinventoryadjustgetAPIRequest) SetWarehouseCode(_warehouseCode string) error {
	r._warehouseCode = _warehouseCode
	r.Set("warehouse_code", _warehouseCode)
	return nil
}

// GetWarehouseCode WarehouseCode Getter
func (r AlibabawdkumsinventoryadjustgetAPIRequest) GetWarehouseCode() string {
	return r._warehouseCode
}
