package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkumsinventorycheckgetAPIRequest 盘点结果单-回流单 API请求
// alibaba.wdk.ums.inventory.check.get
//
// 盘点结果单-回流单
type AlibabawdkumsinventorycheckgetAPIRequest struct {
	model.Params
	// 店仓code，指的是库调对象，对应一个物理店或仓编码
	_warehouseCode string
}

// NewAlibabawdkumsinventorycheckgetRequest 初始化AlibabawdkumsinventorycheckgetAPIRequest对象
func NewAlibabawdkumsinventorycheckgetRequest() *AlibabawdkumsinventorycheckgetAPIRequest {
	return &AlibabawdkumsinventorycheckgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabawdkumsinventorycheckgetAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.ums.inventory.check.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabawdkumsinventorycheckgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabawdkumsinventorycheckgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetWarehouseCode is WarehouseCode Setter
// 店仓code，指的是库调对象，对应一个物理店或仓编码
func (r *AlibabawdkumsinventorycheckgetAPIRequest) SetWarehouseCode(_warehouseCode string) error {
	r._warehouseCode = _warehouseCode
	r.Set("warehouse_code", _warehouseCode)
	return nil
}

// GetWarehouseCode WarehouseCode Getter
func (r AlibabawdkumsinventorycheckgetAPIRequest) GetWarehouseCode() string {
	return r._warehouseCode
}
