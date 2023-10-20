package waybill

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// CainiaosmartdeliverystrategywarehouseideleteAPIRequest 删除智能发货引擎仓策略 API请求
// cainiao.smartdelivery.strategy.warehouse.i.delete
//
// 删除智能发货引擎仓策略
type CainiaosmartdeliverystrategywarehouseideleteAPIRequest struct {
	model.Params
	// 仓id
	_warehouseId int64
}

// NewCainiaosmartdeliverystrategywarehouseideleteRequest 初始化CainiaosmartdeliverystrategywarehouseideleteAPIRequest对象
func NewCainiaosmartdeliverystrategywarehouseideleteRequest() *CainiaosmartdeliverystrategywarehouseideleteAPIRequest {
	return &CainiaosmartdeliverystrategywarehouseideleteAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r CainiaosmartdeliverystrategywarehouseideleteAPIRequest) GetApiMethodName() string {
	return "cainiao.smartdelivery.strategy.warehouse.i.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r CainiaosmartdeliverystrategywarehouseideleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r CainiaosmartdeliverystrategywarehouseideleteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetWarehouseId is WarehouseId Setter
// 仓id
func (r *CainiaosmartdeliverystrategywarehouseideleteAPIRequest) SetWarehouseId(_warehouseId int64) error {
	r._warehouseId = _warehouseId
	r.Set("warehouse_id", _warehouseId)
	return nil
}

// GetWarehouseId WarehouseId Getter
func (r CainiaosmartdeliverystrategywarehouseideleteAPIRequest) GetWarehouseId() int64 {
	return r._warehouseId
}
