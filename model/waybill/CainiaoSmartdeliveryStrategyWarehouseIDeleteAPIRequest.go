package waybill

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// CainiaoSmartdeliveryStrategyWarehouseIDeleteAPIRequest 删除智能发货引擎仓策略 API请求
// cainiao.smartdelivery.strategy.warehouse.i.delete
//
// 删除智能发货引擎仓策略
type CainiaoSmartdeliveryStrategyWarehouseIDeleteAPIRequest struct {
	model.Params
	// 仓id
	_warehouseId int64
}

// NewCainiaoSmartdeliveryStrategyWarehouseIDeleteRequest 初始化CainiaoSmartdeliveryStrategyWarehouseIDeleteAPIRequest对象
func NewCainiaoSmartdeliveryStrategyWarehouseIDeleteRequest() *CainiaoSmartdeliveryStrategyWarehouseIDeleteAPIRequest {
	return &CainiaoSmartdeliveryStrategyWarehouseIDeleteAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *CainiaoSmartdeliveryStrategyWarehouseIDeleteAPIRequest) Reset() {
	r._warehouseId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r CainiaoSmartdeliveryStrategyWarehouseIDeleteAPIRequest) GetApiMethodName() string {
	return "cainiao.smartdelivery.strategy.warehouse.i.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r CainiaoSmartdeliveryStrategyWarehouseIDeleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r CainiaoSmartdeliveryStrategyWarehouseIDeleteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetWarehouseId is WarehouseId Setter
// 仓id
func (r *CainiaoSmartdeliveryStrategyWarehouseIDeleteAPIRequest) SetWarehouseId(_warehouseId int64) error {
	r._warehouseId = _warehouseId
	r.Set("warehouse_id", _warehouseId)
	return nil
}

// GetWarehouseId WarehouseId Getter
func (r CainiaoSmartdeliveryStrategyWarehouseIDeleteAPIRequest) GetWarehouseId() int64 {
	return r._warehouseId
}

var poolCainiaoSmartdeliveryStrategyWarehouseIDeleteAPIRequest = sync.Pool{
	New: func() any {
		return NewCainiaoSmartdeliveryStrategyWarehouseIDeleteRequest()
	},
}

// GetCainiaoSmartdeliveryStrategyWarehouseIDeleteRequest 从 sync.Pool 获取 CainiaoSmartdeliveryStrategyWarehouseIDeleteAPIRequest
func GetCainiaoSmartdeliveryStrategyWarehouseIDeleteAPIRequest() *CainiaoSmartdeliveryStrategyWarehouseIDeleteAPIRequest {
	return poolCainiaoSmartdeliveryStrategyWarehouseIDeleteAPIRequest.Get().(*CainiaoSmartdeliveryStrategyWarehouseIDeleteAPIRequest)
}

// ReleaseCainiaoSmartdeliveryStrategyWarehouseIDeleteAPIRequest 将 CainiaoSmartdeliveryStrategyWarehouseIDeleteAPIRequest 放入 sync.Pool
func ReleaseCainiaoSmartdeliveryStrategyWarehouseIDeleteAPIRequest(v *CainiaoSmartdeliveryStrategyWarehouseIDeleteAPIRequest) {
	v.Reset()
	poolCainiaoSmartdeliveryStrategyWarehouseIDeleteAPIRequest.Put(v)
}
