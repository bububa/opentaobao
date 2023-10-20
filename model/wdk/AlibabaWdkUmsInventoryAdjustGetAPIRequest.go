package wdk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkUmsInventoryAdjustGetAPIRequest 库调单-回流单 API请求
// alibaba.wdk.ums.inventory.adjust.get
//
// 库调单-回流单
type AlibabaWdkUmsInventoryAdjustGetAPIRequest struct {
	model.Params
	// 店仓code，指的是库调对象，对应一个物理店或仓编码
	_warehouseCode string
}

// NewAlibabaWdkUmsInventoryAdjustGetRequest 初始化AlibabaWdkUmsInventoryAdjustGetAPIRequest对象
func NewAlibabaWdkUmsInventoryAdjustGetRequest() *AlibabaWdkUmsInventoryAdjustGetAPIRequest {
	return &AlibabaWdkUmsInventoryAdjustGetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaWdkUmsInventoryAdjustGetAPIRequest) Reset() {
	r._warehouseCode = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkUmsInventoryAdjustGetAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.ums.inventory.adjust.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkUmsInventoryAdjustGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaWdkUmsInventoryAdjustGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetWarehouseCode is WarehouseCode Setter
// 店仓code，指的是库调对象，对应一个物理店或仓编码
func (r *AlibabaWdkUmsInventoryAdjustGetAPIRequest) SetWarehouseCode(_warehouseCode string) error {
	r._warehouseCode = _warehouseCode
	r.Set("warehouse_code", _warehouseCode)
	return nil
}

// GetWarehouseCode WarehouseCode Getter
func (r AlibabaWdkUmsInventoryAdjustGetAPIRequest) GetWarehouseCode() string {
	return r._warehouseCode
}

var poolAlibabaWdkUmsInventoryAdjustGetAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaWdkUmsInventoryAdjustGetRequest()
	},
}

// GetAlibabaWdkUmsInventoryAdjustGetRequest 从 sync.Pool 获取 AlibabaWdkUmsInventoryAdjustGetAPIRequest
func GetAlibabaWdkUmsInventoryAdjustGetAPIRequest() *AlibabaWdkUmsInventoryAdjustGetAPIRequest {
	return poolAlibabaWdkUmsInventoryAdjustGetAPIRequest.Get().(*AlibabaWdkUmsInventoryAdjustGetAPIRequest)
}

// ReleaseAlibabaWdkUmsInventoryAdjustGetAPIRequest 将 AlibabaWdkUmsInventoryAdjustGetAPIRequest 放入 sync.Pool
func ReleaseAlibabaWdkUmsInventoryAdjustGetAPIRequest(v *AlibabaWdkUmsInventoryAdjustGetAPIRequest) {
	v.Reset()
	poolAlibabaWdkUmsInventoryAdjustGetAPIRequest.Put(v)
}
