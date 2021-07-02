package waybill

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r CainiaoSmartdeliveryStrategyWarehouseIUpdateAPIRequest) GetApiMethodName() string {
	return "cainiao.smartdelivery.strategy.warehouse.i.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r CainiaoSmartdeliveryStrategyWarehouseIUpdateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
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
