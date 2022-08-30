package ascp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDchainIsvWmsOrderprocessReportAPIRequest 仓作业信息同步 API请求
// alibaba.dchain.isv.wms.orderprocess.report
//
// 仓作业信息同步
type AlibabaDchainIsvWmsOrderprocessReportAPIRequest struct {
	model.Params
	// 请求入参
	_orderProcessReportRequest *WmsOrderProcessReportRequest
}

// NewAlibabaDchainIsvWmsOrderprocessReportRequest 初始化AlibabaDchainIsvWmsOrderprocessReportAPIRequest对象
func NewAlibabaDchainIsvWmsOrderprocessReportRequest() *AlibabaDchainIsvWmsOrderprocessReportAPIRequest {
	return &AlibabaDchainIsvWmsOrderprocessReportAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaDchainIsvWmsOrderprocessReportAPIRequest) GetApiMethodName() string {
	return "alibaba.dchain.isv.wms.orderprocess.report"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaDchainIsvWmsOrderprocessReportAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetOrderProcessReportRequest is OrderProcessReportRequest Setter
// 请求入参
func (r *AlibabaDchainIsvWmsOrderprocessReportAPIRequest) SetOrderProcessReportRequest(_orderProcessReportRequest *WmsOrderProcessReportRequest) error {
	r._orderProcessReportRequest = _orderProcessReportRequest
	r.Set("order_process_report_request", _orderProcessReportRequest)
	return nil
}

// GetOrderProcessReportRequest OrderProcessReportRequest Getter
func (r AlibabaDchainIsvWmsOrderprocessReportAPIRequest) GetOrderProcessReportRequest() *WmsOrderProcessReportRequest {
	return r._orderProcessReportRequest
}
