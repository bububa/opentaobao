package ascp

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDchainAoxiangWmsDeliveryorderCreateAPIRequest 回传仓库接单通知 API请求
// alibaba.dchain.aoxiang.wms.deliveryorder.create
//
// WMS上报仓库接单节点状态信息，代表接单环节。
type AlibabaDchainAoxiangWmsDeliveryorderCreateAPIRequest struct {
	model.Params
	// 接单回传上报请求
	_deliveryOrderReportRequest *DeliveryOrderReportRequest
}

// NewAlibabaDchainAoxiangWmsDeliveryorderCreateRequest 初始化AlibabaDchainAoxiangWmsDeliveryorderCreateAPIRequest对象
func NewAlibabaDchainAoxiangWmsDeliveryorderCreateRequest() *AlibabaDchainAoxiangWmsDeliveryorderCreateAPIRequest {
	return &AlibabaDchainAoxiangWmsDeliveryorderCreateAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaDchainAoxiangWmsDeliveryorderCreateAPIRequest) Reset() {
	r._deliveryOrderReportRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaDchainAoxiangWmsDeliveryorderCreateAPIRequest) GetApiMethodName() string {
	return "alibaba.dchain.aoxiang.wms.deliveryorder.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaDchainAoxiangWmsDeliveryorderCreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaDchainAoxiangWmsDeliveryorderCreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDeliveryOrderReportRequest is DeliveryOrderReportRequest Setter
// 接单回传上报请求
func (r *AlibabaDchainAoxiangWmsDeliveryorderCreateAPIRequest) SetDeliveryOrderReportRequest(_deliveryOrderReportRequest *DeliveryOrderReportRequest) error {
	r._deliveryOrderReportRequest = _deliveryOrderReportRequest
	r.Set("delivery_order_report_request", _deliveryOrderReportRequest)
	return nil
}

// GetDeliveryOrderReportRequest DeliveryOrderReportRequest Getter
func (r AlibabaDchainAoxiangWmsDeliveryorderCreateAPIRequest) GetDeliveryOrderReportRequest() *DeliveryOrderReportRequest {
	return r._deliveryOrderReportRequest
}

var poolAlibabaDchainAoxiangWmsDeliveryorderCreateAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaDchainAoxiangWmsDeliveryorderCreateRequest()
	},
}

// GetAlibabaDchainAoxiangWmsDeliveryorderCreateRequest 从 sync.Pool 获取 AlibabaDchainAoxiangWmsDeliveryorderCreateAPIRequest
func GetAlibabaDchainAoxiangWmsDeliveryorderCreateAPIRequest() *AlibabaDchainAoxiangWmsDeliveryorderCreateAPIRequest {
	return poolAlibabaDchainAoxiangWmsDeliveryorderCreateAPIRequest.Get().(*AlibabaDchainAoxiangWmsDeliveryorderCreateAPIRequest)
}

// ReleaseAlibabaDchainAoxiangWmsDeliveryorderCreateAPIRequest 将 AlibabaDchainAoxiangWmsDeliveryorderCreateAPIRequest 放入 sync.Pool
func ReleaseAlibabaDchainAoxiangWmsDeliveryorderCreateAPIRequest(v *AlibabaDchainAoxiangWmsDeliveryorderCreateAPIRequest) {
	v.Reset()
	poolAlibabaDchainAoxiangWmsDeliveryorderCreateAPIRequest.Put(v)
}
