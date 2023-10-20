package ascp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabadchainaoxiangwmsdeliveryorderconfirmAPIRequest 回传发货单确认 API请求
// alibaba.dchain.aoxiang.wms.deliveryorder.confirm
//
// 回传发货单确认
type AlibabadchainaoxiangwmsdeliveryorderconfirmAPIRequest struct {
	model.Params
	// 确认接单回传上报请求
	_deliveryOrderConfirmReportRequest *DeliveryOrderConfirmReportRequest
}

// NewAlibabadchainaoxiangwmsdeliveryorderconfirmRequest 初始化AlibabadchainaoxiangwmsdeliveryorderconfirmAPIRequest对象
func NewAlibabadchainaoxiangwmsdeliveryorderconfirmRequest() *AlibabadchainaoxiangwmsdeliveryorderconfirmAPIRequest {
	return &AlibabadchainaoxiangwmsdeliveryorderconfirmAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabadchainaoxiangwmsdeliveryorderconfirmAPIRequest) GetApiMethodName() string {
	return "alibaba.dchain.aoxiang.wms.deliveryorder.confirm"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabadchainaoxiangwmsdeliveryorderconfirmAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabadchainaoxiangwmsdeliveryorderconfirmAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDeliveryOrderConfirmReportRequest is DeliveryOrderConfirmReportRequest Setter
// 确认接单回传上报请求
func (r *AlibabadchainaoxiangwmsdeliveryorderconfirmAPIRequest) SetDeliveryOrderConfirmReportRequest(_deliveryOrderConfirmReportRequest *DeliveryOrderConfirmReportRequest) error {
	r._deliveryOrderConfirmReportRequest = _deliveryOrderConfirmReportRequest
	r.Set("delivery_order_confirm_report_request", _deliveryOrderConfirmReportRequest)
	return nil
}

// GetDeliveryOrderConfirmReportRequest DeliveryOrderConfirmReportRequest Getter
func (r AlibabadchainaoxiangwmsdeliveryorderconfirmAPIRequest) GetDeliveryOrderConfirmReportRequest() *DeliveryOrderConfirmReportRequest {
	return r._deliveryOrderConfirmReportRequest
}
