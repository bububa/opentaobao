package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkUmsShiftGetAPIRequest 移库单获取 API请求
// alibaba.wdk.ums.shift.get
//
// 移库单获取
type AlibabaWdkUmsShiftGetAPIRequest struct {
	model.Params
	// 店仓code，指的是库调对象，对应一个物理店或仓编码
	_warehouseCode string
}

// NewAlibabaWdkUmsShiftGetRequest 初始化AlibabaWdkUmsShiftGetAPIRequest对象
func NewAlibabaWdkUmsShiftGetRequest() *AlibabaWdkUmsShiftGetAPIRequest {
	return &AlibabaWdkUmsShiftGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkUmsShiftGetAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.ums.shift.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkUmsShiftGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetWarehouseCode is WarehouseCode Setter
// 店仓code，指的是库调对象，对应一个物理店或仓编码
func (r *AlibabaWdkUmsShiftGetAPIRequest) SetWarehouseCode(_warehouseCode string) error {
	r._warehouseCode = _warehouseCode
	r.Set("warehouse_code", _warehouseCode)
	return nil
}

// GetWarehouseCode WarehouseCode Getter
func (r AlibabaWdkUmsShiftGetAPIRequest) GetWarehouseCode() string {
	return r._warehouseCode
}
