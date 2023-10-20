package waybill

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// CainiaosmartdeliverystrategywarehouseiupdateAPIRequest 智能发货引擎策略仓设置 API请求
// cainiao.smartdelivery.strategy.warehouse.i.update
//
// 智能发货引擎发货策略设置仓维度
type CainiaosmartdeliverystrategywarehouseiupdateAPIRequest struct {
	model.Params
	// 智能发货设置请求参数
	_deliveryStrategySetRequest *DeliveryStrategySetRequest
}

// NewCainiaosmartdeliverystrategywarehouseiupdateRequest 初始化CainiaosmartdeliverystrategywarehouseiupdateAPIRequest对象
func NewCainiaosmartdeliverystrategywarehouseiupdateRequest() *CainiaosmartdeliverystrategywarehouseiupdateAPIRequest {
	return &CainiaosmartdeliverystrategywarehouseiupdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r CainiaosmartdeliverystrategywarehouseiupdateAPIRequest) GetApiMethodName() string {
	return "cainiao.smartdelivery.strategy.warehouse.i.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r CainiaosmartdeliverystrategywarehouseiupdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r CainiaosmartdeliverystrategywarehouseiupdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDeliveryStrategySetRequest is DeliveryStrategySetRequest Setter
// 智能发货设置请求参数
func (r *CainiaosmartdeliverystrategywarehouseiupdateAPIRequest) SetDeliveryStrategySetRequest(_deliveryStrategySetRequest *DeliveryStrategySetRequest) error {
	r._deliveryStrategySetRequest = _deliveryStrategySetRequest
	r.Set("delivery_strategy_set_request", _deliveryStrategySetRequest)
	return nil
}

// GetDeliveryStrategySetRequest DeliveryStrategySetRequest Getter
func (r CainiaosmartdeliverystrategywarehouseiupdateAPIRequest) GetDeliveryStrategySetRequest() *DeliveryStrategySetRequest {
	return r._deliveryStrategySetRequest
}
