package waybill

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// CainiaoSmartdeliveryStrategyWarehouseIUpdateAPIRequest 智能发货引擎策略仓设置 API请求
// cainiao.smartdelivery.strategy.warehouse.i.update
//
// 智能发货引擎发货策略设置仓维度
type CainiaoSmartdeliveryStrategyWarehouseIUpdateAPIRequest struct {
	model.Params
	// 智能发货设置请求参数
	_deliveryStrategySetRequest *DeliveryStrategySetRequest
}

// NewCainiaoSmartdeliveryStrategyWarehouseIUpdateRequest 初始化CainiaoSmartdeliveryStrategyWarehouseIUpdateAPIRequest对象
func NewCainiaoSmartdeliveryStrategyWarehouseIUpdateRequest() *CainiaoSmartdeliveryStrategyWarehouseIUpdateAPIRequest {
	return &CainiaoSmartdeliveryStrategyWarehouseIUpdateAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *CainiaoSmartdeliveryStrategyWarehouseIUpdateAPIRequest) Reset() {
	r._deliveryStrategySetRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r CainiaoSmartdeliveryStrategyWarehouseIUpdateAPIRequest) GetApiMethodName() string {
	return "cainiao.smartdelivery.strategy.warehouse.i.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r CainiaoSmartdeliveryStrategyWarehouseIUpdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r CainiaoSmartdeliveryStrategyWarehouseIUpdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDeliveryStrategySetRequest is DeliveryStrategySetRequest Setter
// 智能发货设置请求参数
func (r *CainiaoSmartdeliveryStrategyWarehouseIUpdateAPIRequest) SetDeliveryStrategySetRequest(_deliveryStrategySetRequest *DeliveryStrategySetRequest) error {
	r._deliveryStrategySetRequest = _deliveryStrategySetRequest
	r.Set("delivery_strategy_set_request", _deliveryStrategySetRequest)
	return nil
}

// GetDeliveryStrategySetRequest DeliveryStrategySetRequest Getter
func (r CainiaoSmartdeliveryStrategyWarehouseIUpdateAPIRequest) GetDeliveryStrategySetRequest() *DeliveryStrategySetRequest {
	return r._deliveryStrategySetRequest
}

var poolCainiaoSmartdeliveryStrategyWarehouseIUpdateAPIRequest = sync.Pool{
	New: func() any {
		return NewCainiaoSmartdeliveryStrategyWarehouseIUpdateRequest()
	},
}

// GetCainiaoSmartdeliveryStrategyWarehouseIUpdateRequest 从 sync.Pool 获取 CainiaoSmartdeliveryStrategyWarehouseIUpdateAPIRequest
func GetCainiaoSmartdeliveryStrategyWarehouseIUpdateAPIRequest() *CainiaoSmartdeliveryStrategyWarehouseIUpdateAPIRequest {
	return poolCainiaoSmartdeliveryStrategyWarehouseIUpdateAPIRequest.Get().(*CainiaoSmartdeliveryStrategyWarehouseIUpdateAPIRequest)
}

// ReleaseCainiaoSmartdeliveryStrategyWarehouseIUpdateAPIRequest 将 CainiaoSmartdeliveryStrategyWarehouseIUpdateAPIRequest 放入 sync.Pool
func ReleaseCainiaoSmartdeliveryStrategyWarehouseIUpdateAPIRequest(v *CainiaoSmartdeliveryStrategyWarehouseIUpdateAPIRequest) {
	v.Reset()
	poolCainiaoSmartdeliveryStrategyWarehouseIUpdateAPIRequest.Put(v)
}
