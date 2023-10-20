package wdk

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaWdkUmsShiftGetAPIRequest) Reset() {
	r._warehouseCode = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkUmsShiftGetAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.ums.shift.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkUmsShiftGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaWdkUmsShiftGetAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolAlibabaWdkUmsShiftGetAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaWdkUmsShiftGetRequest()
	},
}

// GetAlibabaWdkUmsShiftGetRequest 从 sync.Pool 获取 AlibabaWdkUmsShiftGetAPIRequest
func GetAlibabaWdkUmsShiftGetAPIRequest() *AlibabaWdkUmsShiftGetAPIRequest {
	return poolAlibabaWdkUmsShiftGetAPIRequest.Get().(*AlibabaWdkUmsShiftGetAPIRequest)
}

// ReleaseAlibabaWdkUmsShiftGetAPIRequest 将 AlibabaWdkUmsShiftGetAPIRequest 放入 sync.Pool
func ReleaseAlibabaWdkUmsShiftGetAPIRequest(v *AlibabaWdkUmsShiftGetAPIRequest) {
	v.Reset()
	poolAlibabaWdkUmsShiftGetAPIRequest.Put(v)
}
