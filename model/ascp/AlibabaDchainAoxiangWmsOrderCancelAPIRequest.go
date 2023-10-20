package ascp

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDchainAoxiangWmsOrderCancelAPIRequest 回传发货单取消通知 API请求
// alibaba.dchain.aoxiang.wms.order.cancel
//
// 回传发货单取消通知
type AlibabaDchainAoxiangWmsOrderCancelAPIRequest struct {
	model.Params
	// 单据取消回传上报请求
	_orderCancelReportRequest *OrderCancelReportRequest
}

// NewAlibabaDchainAoxiangWmsOrderCancelRequest 初始化AlibabaDchainAoxiangWmsOrderCancelAPIRequest对象
func NewAlibabaDchainAoxiangWmsOrderCancelRequest() *AlibabaDchainAoxiangWmsOrderCancelAPIRequest {
	return &AlibabaDchainAoxiangWmsOrderCancelAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaDchainAoxiangWmsOrderCancelAPIRequest) Reset() {
	r._orderCancelReportRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaDchainAoxiangWmsOrderCancelAPIRequest) GetApiMethodName() string {
	return "alibaba.dchain.aoxiang.wms.order.cancel"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaDchainAoxiangWmsOrderCancelAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaDchainAoxiangWmsOrderCancelAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOrderCancelReportRequest is OrderCancelReportRequest Setter
// 单据取消回传上报请求
func (r *AlibabaDchainAoxiangWmsOrderCancelAPIRequest) SetOrderCancelReportRequest(_orderCancelReportRequest *OrderCancelReportRequest) error {
	r._orderCancelReportRequest = _orderCancelReportRequest
	r.Set("order_cancel_report_request", _orderCancelReportRequest)
	return nil
}

// GetOrderCancelReportRequest OrderCancelReportRequest Getter
func (r AlibabaDchainAoxiangWmsOrderCancelAPIRequest) GetOrderCancelReportRequest() *OrderCancelReportRequest {
	return r._orderCancelReportRequest
}

var poolAlibabaDchainAoxiangWmsOrderCancelAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaDchainAoxiangWmsOrderCancelRequest()
	},
}

// GetAlibabaDchainAoxiangWmsOrderCancelRequest 从 sync.Pool 获取 AlibabaDchainAoxiangWmsOrderCancelAPIRequest
func GetAlibabaDchainAoxiangWmsOrderCancelAPIRequest() *AlibabaDchainAoxiangWmsOrderCancelAPIRequest {
	return poolAlibabaDchainAoxiangWmsOrderCancelAPIRequest.Get().(*AlibabaDchainAoxiangWmsOrderCancelAPIRequest)
}

// ReleaseAlibabaDchainAoxiangWmsOrderCancelAPIRequest 将 AlibabaDchainAoxiangWmsOrderCancelAPIRequest 放入 sync.Pool
func ReleaseAlibabaDchainAoxiangWmsOrderCancelAPIRequest(v *AlibabaDchainAoxiangWmsOrderCancelAPIRequest) {
	v.Reset()
	poolAlibabaDchainAoxiangWmsOrderCancelAPIRequest.Put(v)
}
