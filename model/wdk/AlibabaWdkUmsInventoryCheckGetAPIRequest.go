package wdk

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaWdkUmsInventoryCheckGetAPIRequest) Reset() {
	r._warehouseCode = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkUmsInventoryCheckGetAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.ums.inventory.check.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkUmsInventoryCheckGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaWdkUmsInventoryCheckGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetWarehouseCode is WarehouseCode Setter
// 店仓code，指的是库调对象，对应一个物理店或仓编码
func (r *AlibabaWdkUmsInventoryCheckGetAPIRequest) SetWarehouseCode(_warehouseCode string) error {
	r._warehouseCode = _warehouseCode
	r.Set("warehouse_code", _warehouseCode)
	return nil
}

// GetWarehouseCode WarehouseCode Getter
func (r AlibabaWdkUmsInventoryCheckGetAPIRequest) GetWarehouseCode() string {
	return r._warehouseCode
}

var poolAlibabaWdkUmsInventoryCheckGetAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaWdkUmsInventoryCheckGetRequest()
	},
}

// GetAlibabaWdkUmsInventoryCheckGetRequest 从 sync.Pool 获取 AlibabaWdkUmsInventoryCheckGetAPIRequest
func GetAlibabaWdkUmsInventoryCheckGetAPIRequest() *AlibabaWdkUmsInventoryCheckGetAPIRequest {
	return poolAlibabaWdkUmsInventoryCheckGetAPIRequest.Get().(*AlibabaWdkUmsInventoryCheckGetAPIRequest)
}

// ReleaseAlibabaWdkUmsInventoryCheckGetAPIRequest 将 AlibabaWdkUmsInventoryCheckGetAPIRequest 放入 sync.Pool
func ReleaseAlibabaWdkUmsInventoryCheckGetAPIRequest(v *AlibabaWdkUmsInventoryCheckGetAPIRequest) {
	v.Reset()
	poolAlibabaWdkUmsInventoryCheckGetAPIRequest.Put(v)
}
