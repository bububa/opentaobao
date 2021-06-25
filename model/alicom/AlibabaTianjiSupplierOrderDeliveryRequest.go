package alicom

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
天机供应商发货 APIRequest
alibaba.tianji.supplier.order.delivery

天机供应商发货
*/
type AlibabaTianjiSupplierOrderDeliveryRequest struct {
    model.Params

    // 物流信息
    paramDistributionOrderLogisticsModel   *DistributionOrderLogisticsModel 

}

func NewAlibabaTianjiSupplierOrderDeliveryRequest() *AlibabaTianjiSupplierOrderDeliveryRequest{
    return &AlibabaTianjiSupplierOrderDeliveryRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaTianjiSupplierOrderDeliveryRequest) GetApiMethodName() string {
    return "alibaba.tianji.supplier.order.delivery"
}

func (r AlibabaTianjiSupplierOrderDeliveryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaTianjiSupplierOrderDeliveryRequest) SetParamDistributionOrderLogisticsModel(paramDistributionOrderLogisticsModel *DistributionOrderLogisticsModel) error {
    r.paramDistributionOrderLogisticsModel = paramDistributionOrderLogisticsModel
    r.Set("param_distribution_order_logistics_model", paramDistributionOrderLogisticsModel)
    return nil
}

func (r AlibabaTianjiSupplierOrderDeliveryRequest) GetParamDistributionOrderLogisticsModel() *DistributionOrderLogisticsModel {
    return r.paramDistributionOrderLogisticsModel
}

