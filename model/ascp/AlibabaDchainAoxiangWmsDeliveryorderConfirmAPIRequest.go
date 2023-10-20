package ascp

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDchainAoxiangWmsDeliveryorderConfirmAPIRequest 回传发货单确认 API请求
// alibaba.dchain.aoxiang.wms.deliveryorder.confirm
//
// 回传发货单确认
type AlibabaDchainAoxiangWmsDeliveryorderConfirmAPIRequest struct {
	model.Params
	// 确认接单回传上报请求
	_deliveryOrderConfirmReportRequest *DeliveryOrderConfirmReportRequest
}

// NewAlibabaDchainAoxiangWmsDeliveryorderConfirmRequest 初始化AlibabaDchainAoxiangWmsDeliveryorderConfirmAPIRequest对象
func NewAlibabaDchainAoxiangWmsDeliveryorderConfirmRequest() *AlibabaDchainAoxiangWmsDeliveryorderConfirmAPIRequest {
	return &AlibabaDchainAoxiangWmsDeliveryorderConfirmAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaDchainAoxiangWmsDeliveryorderConfirmAPIRequest) Reset() {
	r._deliveryOrderConfirmReportRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaDchainAoxiangWmsDeliveryorderConfirmAPIRequest) GetApiMethodName() string {
	return "alibaba.dchain.aoxiang.wms.deliveryorder.confirm"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaDchainAoxiangWmsDeliveryorderConfirmAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaDchainAoxiangWmsDeliveryorderConfirmAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDeliveryOrderConfirmReportRequest is DeliveryOrderConfirmReportRequest Setter
// 确认接单回传上报请求
func (r *AlibabaDchainAoxiangWmsDeliveryorderConfirmAPIRequest) SetDeliveryOrderConfirmReportRequest(_deliveryOrderConfirmReportRequest *DeliveryOrderConfirmReportRequest) error {
	r._deliveryOrderConfirmReportRequest = _deliveryOrderConfirmReportRequest
	r.Set("delivery_order_confirm_report_request", _deliveryOrderConfirmReportRequest)
	return nil
}

// GetDeliveryOrderConfirmReportRequest DeliveryOrderConfirmReportRequest Getter
func (r AlibabaDchainAoxiangWmsDeliveryorderConfirmAPIRequest) GetDeliveryOrderConfirmReportRequest() *DeliveryOrderConfirmReportRequest {
	return r._deliveryOrderConfirmReportRequest
}

var poolAlibabaDchainAoxiangWmsDeliveryorderConfirmAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaDchainAoxiangWmsDeliveryorderConfirmRequest()
	},
}

// GetAlibabaDchainAoxiangWmsDeliveryorderConfirmRequest 从 sync.Pool 获取 AlibabaDchainAoxiangWmsDeliveryorderConfirmAPIRequest
func GetAlibabaDchainAoxiangWmsDeliveryorderConfirmAPIRequest() *AlibabaDchainAoxiangWmsDeliveryorderConfirmAPIRequest {
	return poolAlibabaDchainAoxiangWmsDeliveryorderConfirmAPIRequest.Get().(*AlibabaDchainAoxiangWmsDeliveryorderConfirmAPIRequest)
}

// ReleaseAlibabaDchainAoxiangWmsDeliveryorderConfirmAPIRequest 将 AlibabaDchainAoxiangWmsDeliveryorderConfirmAPIRequest 放入 sync.Pool
func ReleaseAlibabaDchainAoxiangWmsDeliveryorderConfirmAPIRequest(v *AlibabaDchainAoxiangWmsDeliveryorderConfirmAPIRequest) {
	v.Reset()
	poolAlibabaDchainAoxiangWmsDeliveryorderConfirmAPIRequest.Put(v)
}
