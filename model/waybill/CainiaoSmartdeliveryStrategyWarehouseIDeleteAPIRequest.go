package waybill

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* CainiaoSmartdeliveryStrategyWarehouseIDeleteAPIRequest
删除智能发货引擎仓策略 API请求
cainiao.smartdelivery.strategy.warehouse.i.delete

删除智能发货引擎仓策略 */
type CainiaoSmartdeliveryStrategyWarehouseIDeleteAPIRequest struct {
	model.Params
	// 仓id
	_warehouseId int64
}

// NewCainiaoSmartdeliveryStrategyWarehouseIDeleteRequest 初始化CainiaoSmartdeliveryStrategyWarehouseIDeleteAPIRequest对象
func NewCainiaoSmartdeliveryStrategyWarehouseIDeleteRequest() *CainiaoSmartdeliveryStrategyWarehouseIDeleteAPIRequest {
	return &CainiaoSmartdeliveryStrategyWarehouseIDeleteAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r CainiaoSmartdeliveryStrategyWarehouseIDeleteAPIRequest) GetApiMethodName() string {
	return "cainiao.smartdelivery.strategy.warehouse.i.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r CainiaoSmartdeliveryStrategyWarehouseIDeleteAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is WarehouseId Setter
// 仓id
func (r *CainiaoSmartdeliveryStrategyWarehouseIDeleteAPIRequest) SetWarehouseId(_warehouseId int64) error {
	r._warehouseId = _warehouseId
	r.Set("warehouse_id", _warehouseId)
	return nil
}

// Get WarehouseId Getter
func (r CainiaoSmartdeliveryStrategyWarehouseIDeleteAPIRequest) GetWarehouseId() int64 {
	return r._warehouseId
}
