package ascp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDchainAoxiangWmsOrderprocessReportAPIRequest 回传发货单流水通知 API请求
// alibaba.dchain.aoxiang.wms.orderprocess.report
//
// 回传发货单流水通知
type AlibabaDchainAoxiangWmsOrderprocessReportAPIRequest struct {
	model.Params
	// 回传单据流水请求
	_orderProcessReportRequest *OrderProcessReportRequest
}

// NewAlibabaDchainAoxiangWmsOrderprocessReportRequest 初始化AlibabaDchainAoxiangWmsOrderprocessReportAPIRequest对象
func NewAlibabaDchainAoxiangWmsOrderprocessReportRequest() *AlibabaDchainAoxiangWmsOrderprocessReportAPIRequest {
	return &AlibabaDchainAoxiangWmsOrderprocessReportAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaDchainAoxiangWmsOrderprocessReportAPIRequest) GetApiMethodName() string {
	return "alibaba.dchain.aoxiang.wms.orderprocess.report"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaDchainAoxiangWmsOrderprocessReportAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaDchainAoxiangWmsOrderprocessReportAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOrderProcessReportRequest is OrderProcessReportRequest Setter
// 回传单据流水请求
func (r *AlibabaDchainAoxiangWmsOrderprocessReportAPIRequest) SetOrderProcessReportRequest(_orderProcessReportRequest *OrderProcessReportRequest) error {
	r._orderProcessReportRequest = _orderProcessReportRequest
	r.Set("order_process_report_request", _orderProcessReportRequest)
	return nil
}

// GetOrderProcessReportRequest OrderProcessReportRequest Getter
func (r AlibabaDchainAoxiangWmsOrderprocessReportAPIRequest) GetOrderProcessReportRequest() *OrderProcessReportRequest {
	return r._orderProcessReportRequest
}
