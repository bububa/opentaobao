package alicom

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaTianjiSupplierOrderDeliveryAPIRequest 天机供应商发货 API请求
// alibaba.tianji.supplier.order.delivery
//
// 天机供应商发货
type AlibabaTianjiSupplierOrderDeliveryAPIRequest struct {
	model.Params
	// 物流信息
	_paramDistributionOrderLogisticsModel *DistributionOrderLogisticsModel
}

// NewAlibabaTianjiSupplierOrderDeliveryRequest 初始化AlibabaTianjiSupplierOrderDeliveryAPIRequest对象
func NewAlibabaTianjiSupplierOrderDeliveryRequest() *AlibabaTianjiSupplierOrderDeliveryAPIRequest {
	return &AlibabaTianjiSupplierOrderDeliveryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaTianjiSupplierOrderDeliveryAPIRequest) GetApiMethodName() string {
	return "alibaba.tianji.supplier.order.delivery"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaTianjiSupplierOrderDeliveryAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is ParamDistributionOrderLogisticsModel Setter
// 物流信息
func (r *AlibabaTianjiSupplierOrderDeliveryAPIRequest) SetParamDistributionOrderLogisticsModel(_paramDistributionOrderLogisticsModel *DistributionOrderLogisticsModel) error {
	r._paramDistributionOrderLogisticsModel = _paramDistributionOrderLogisticsModel
	r.Set("param_distribution_order_logistics_model", _paramDistributionOrderLogisticsModel)
	return nil
}

// Get ParamDistributionOrderLogisticsModel Getter
func (r AlibabaTianjiSupplierOrderDeliveryAPIRequest) GetParamDistributionOrderLogisticsModel() *DistributionOrderLogisticsModel {
	return r._paramDistributionOrderLogisticsModel
}
