package waybill

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
智能发货引擎策略仓设置 APIRequest
cainiao.smartdelivery.strategy.warehouse.i.update

智能发货引擎发货策略设置仓维度
*/
type CainiaoSmartdeliveryStrategyWarehouseIUpdateRequest struct {
    model.Params

    // 智能发货设置请求参数
    deliveryStrategySetRequest   *DeliveryStrategySetRequest 

}

func NewCainiaoSmartdeliveryStrategyWarehouseIUpdateRequest() *CainiaoSmartdeliveryStrategyWarehouseIUpdateRequest{
    return &CainiaoSmartdeliveryStrategyWarehouseIUpdateRequest{
        Params: model.NewParams(),
    }
}

func (r CainiaoSmartdeliveryStrategyWarehouseIUpdateRequest) GetApiMethodName() string {
    return "cainiao.smartdelivery.strategy.warehouse.i.update"
}

func (r CainiaoSmartdeliveryStrategyWarehouseIUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *CainiaoSmartdeliveryStrategyWarehouseIUpdateRequest) SetDeliveryStrategySetRequest(deliveryStrategySetRequest *DeliveryStrategySetRequest) error {
    r.deliveryStrategySetRequest = deliveryStrategySetRequest
    r.Set("delivery_strategy_set_request", deliveryStrategySetRequest)
    return nil
}

func (r CainiaoSmartdeliveryStrategyWarehouseIUpdateRequest) GetDeliveryStrategySetRequest() *DeliveryStrategySetRequest {
    return r.deliveryStrategySetRequest
}

