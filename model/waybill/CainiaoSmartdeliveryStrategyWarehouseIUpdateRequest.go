package waybill

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
智能发货引擎策略仓设置 API请求
cainiao.smartdelivery.strategy.warehouse.i.update

智能发货引擎发货策略设置仓维度
*/
type CainiaoSmartdeliveryStrategyWarehouseIUpdateRequest struct {
    model.Params
    // 智能发货设置请求参数
    _deliveryStrategySetRequest   *DeliveryStrategySetRequest
}

// 初始化CainiaoSmartdeliveryStrategyWarehouseIUpdateRequest对象
func NewCainiaoSmartdeliveryStrategyWarehouseIUpdateRequest() *CainiaoSmartdeliveryStrategyWarehouseIUpdateRequest{
    return &CainiaoSmartdeliveryStrategyWarehouseIUpdateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r CainiaoSmartdeliveryStrategyWarehouseIUpdateRequest) GetApiMethodName() string {
    return "cainiao.smartdelivery.strategy.warehouse.i.update"
}

// IRequest interface 方法, 获取API参数
func (r CainiaoSmartdeliveryStrategyWarehouseIUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// DeliveryStrategySetRequest Setter
// 智能发货设置请求参数
func (r *CainiaoSmartdeliveryStrategyWarehouseIUpdateRequest) SetDeliveryStrategySetRequest(_deliveryStrategySetRequest *DeliveryStrategySetRequest) error {
    r._deliveryStrategySetRequest = _deliveryStrategySetRequest
    r.Set("delivery_strategy_set_request", _deliveryStrategySetRequest)
    return nil
}

// DeliveryStrategySetRequest Getter
func (r CainiaoSmartdeliveryStrategyWarehouseIUpdateRequest) GetDeliveryStrategySetRequest() *DeliveryStrategySetRequest {
    return r._deliveryStrategySetRequest
}
