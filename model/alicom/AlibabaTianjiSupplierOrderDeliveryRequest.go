package alicom

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
天机供应商发货 API请求
alibaba.tianji.supplier.order.delivery

天机供应商发货
*/
type AlibabaTianjiSupplierOrderDeliveryRequest struct {
    model.Params
    // 物流信息
    _paramDistributionOrderLogisticsModel   *DistributionOrderLogisticsModel
}

// 初始化AlibabaTianjiSupplierOrderDeliveryRequest对象
func NewAlibabaTianjiSupplierOrderDeliveryRequest() *AlibabaTianjiSupplierOrderDeliveryRequest{
    return &AlibabaTianjiSupplierOrderDeliveryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaTianjiSupplierOrderDeliveryRequest) GetApiMethodName() string {
    return "alibaba.tianji.supplier.order.delivery"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaTianjiSupplierOrderDeliveryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamDistributionOrderLogisticsModel Setter
// 物流信息
func (r *AlibabaTianjiSupplierOrderDeliveryRequest) SetParamDistributionOrderLogisticsModel(_paramDistributionOrderLogisticsModel *DistributionOrderLogisticsModel) error {
    r._paramDistributionOrderLogisticsModel = _paramDistributionOrderLogisticsModel
    r.Set("param_distribution_order_logistics_model", _paramDistributionOrderLogisticsModel)
    return nil
}

// ParamDistributionOrderLogisticsModel Getter
func (r AlibabaTianjiSupplierOrderDeliveryRequest) GetParamDistributionOrderLogisticsModel() *DistributionOrderLogisticsModel {
    return r._paramDistributionOrderLogisticsModel
}
