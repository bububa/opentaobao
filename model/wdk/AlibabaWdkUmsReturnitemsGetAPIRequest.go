package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkumsreturnitemsgetAPIRequest 退货库位商品查询（退货出库辅助）-回流单 API请求
// alibaba.wdk.ums.returnitems.get
//
// 退货库位商品查询（退货出库辅助）-回流单
type AlibabawdkumsreturnitemsgetAPIRequest struct {
	model.Params
	// 店仓code，指的是作业对象，对应一个物理店或仓编码
	_warehouseCode string
}

// NewAlibabawdkumsreturnitemsgetRequest 初始化AlibabawdkumsreturnitemsgetAPIRequest对象
func NewAlibabawdkumsreturnitemsgetRequest() *AlibabawdkumsreturnitemsgetAPIRequest {
	return &AlibabawdkumsreturnitemsgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabawdkumsreturnitemsgetAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.ums.returnitems.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabawdkumsreturnitemsgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabawdkumsreturnitemsgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetWarehouseCode is WarehouseCode Setter
// 店仓code，指的是作业对象，对应一个物理店或仓编码
func (r *AlibabawdkumsreturnitemsgetAPIRequest) SetWarehouseCode(_warehouseCode string) error {
	r._warehouseCode = _warehouseCode
	r.Set("warehouse_code", _warehouseCode)
	return nil
}

// GetWarehouseCode WarehouseCode Getter
func (r AlibabawdkumsreturnitemsgetAPIRequest) GetWarehouseCode() string {
	return r._warehouseCode
}
