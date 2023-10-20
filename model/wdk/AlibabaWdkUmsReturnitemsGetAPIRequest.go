package wdk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkUmsReturnitemsGetAPIRequest 退货库位商品查询（退货出库辅助）-回流单 API请求
// alibaba.wdk.ums.returnitems.get
//
// 退货库位商品查询（退货出库辅助）-回流单
type AlibabaWdkUmsReturnitemsGetAPIRequest struct {
	model.Params
	// 店仓code，指的是作业对象，对应一个物理店或仓编码
	_warehouseCode string
}

// NewAlibabaWdkUmsReturnitemsGetRequest 初始化AlibabaWdkUmsReturnitemsGetAPIRequest对象
func NewAlibabaWdkUmsReturnitemsGetRequest() *AlibabaWdkUmsReturnitemsGetAPIRequest {
	return &AlibabaWdkUmsReturnitemsGetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaWdkUmsReturnitemsGetAPIRequest) Reset() {
	r._warehouseCode = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkUmsReturnitemsGetAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.ums.returnitems.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkUmsReturnitemsGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaWdkUmsReturnitemsGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetWarehouseCode is WarehouseCode Setter
// 店仓code，指的是作业对象，对应一个物理店或仓编码
func (r *AlibabaWdkUmsReturnitemsGetAPIRequest) SetWarehouseCode(_warehouseCode string) error {
	r._warehouseCode = _warehouseCode
	r.Set("warehouse_code", _warehouseCode)
	return nil
}

// GetWarehouseCode WarehouseCode Getter
func (r AlibabaWdkUmsReturnitemsGetAPIRequest) GetWarehouseCode() string {
	return r._warehouseCode
}

var poolAlibabaWdkUmsReturnitemsGetAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaWdkUmsReturnitemsGetRequest()
	},
}

// GetAlibabaWdkUmsReturnitemsGetRequest 从 sync.Pool 获取 AlibabaWdkUmsReturnitemsGetAPIRequest
func GetAlibabaWdkUmsReturnitemsGetAPIRequest() *AlibabaWdkUmsReturnitemsGetAPIRequest {
	return poolAlibabaWdkUmsReturnitemsGetAPIRequest.Get().(*AlibabaWdkUmsReturnitemsGetAPIRequest)
}

// ReleaseAlibabaWdkUmsReturnitemsGetAPIRequest 将 AlibabaWdkUmsReturnitemsGetAPIRequest 放入 sync.Pool
func ReleaseAlibabaWdkUmsReturnitemsGetAPIRequest(v *AlibabaWdkUmsReturnitemsGetAPIRequest) {
	v.Reset()
	poolAlibabaWdkUmsReturnitemsGetAPIRequest.Put(v)
}
