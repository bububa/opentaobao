package ascp

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaDchainAoxiangWmsOrderCancelAPIRequest) GetApiMethodName() string {
	return "alibaba.dchain.aoxiang.wms.order.cancel"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaDchainAoxiangWmsOrderCancelAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
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
