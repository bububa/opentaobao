package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkumsshiftgetAPIRequest 移库单获取 API请求
// alibaba.wdk.ums.shift.get
//
// 移库单获取
type AlibabawdkumsshiftgetAPIRequest struct {
	model.Params
	// 店仓code，指的是库调对象，对应一个物理店或仓编码
	_warehouseCode string
}

// NewAlibabawdkumsshiftgetRequest 初始化AlibabawdkumsshiftgetAPIRequest对象
func NewAlibabawdkumsshiftgetRequest() *AlibabawdkumsshiftgetAPIRequest {
	return &AlibabawdkumsshiftgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabawdkumsshiftgetAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.ums.shift.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabawdkumsshiftgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabawdkumsshiftgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetWarehouseCode is WarehouseCode Setter
// 店仓code，指的是库调对象，对应一个物理店或仓编码
func (r *AlibabawdkumsshiftgetAPIRequest) SetWarehouseCode(_warehouseCode string) error {
	r._warehouseCode = _warehouseCode
	r.Set("warehouse_code", _warehouseCode)
	return nil
}

// GetWarehouseCode WarehouseCode Getter
func (r AlibabawdkumsshiftgetAPIRequest) GetWarehouseCode() string {
	return r._warehouseCode
}
