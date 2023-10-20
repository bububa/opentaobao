package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkumshandlinggetAPIRequest 加工单-回流单（新接口） API请求
// alibaba.wdk.ums.handling.get
//
// 加工单-回流单（新接口）
type AlibabawdkumshandlinggetAPIRequest struct {
	model.Params
	// 仓库编码
	_warehouseCode string
}

// NewAlibabawdkumshandlinggetRequest 初始化AlibabawdkumshandlinggetAPIRequest对象
func NewAlibabawdkumshandlinggetRequest() *AlibabawdkumshandlinggetAPIRequest {
	return &AlibabawdkumshandlinggetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabawdkumshandlinggetAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.ums.handling.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabawdkumshandlinggetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabawdkumshandlinggetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetWarehouseCode is WarehouseCode Setter
// 仓库编码
func (r *AlibabawdkumshandlinggetAPIRequest) SetWarehouseCode(_warehouseCode string) error {
	r._warehouseCode = _warehouseCode
	r.Set("warehouse_code", _warehouseCode)
	return nil
}

// GetWarehouseCode WarehouseCode Getter
func (r AlibabawdkumshandlinggetAPIRequest) GetWarehouseCode() string {
	return r._warehouseCode
}
