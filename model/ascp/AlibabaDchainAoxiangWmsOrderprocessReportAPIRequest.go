package ascp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabadchainaoxiangwmsorderprocessreportAPIRequest 回传发货单流水通知 API请求
// alibaba.dchain.aoxiang.wms.orderprocess.report
//
// 回传发货单流水通知
type AlibabadchainaoxiangwmsorderprocessreportAPIRequest struct {
	model.Params
	// 回传单据流水请求
	_orderProcessReportRequest *OrderProcessReportRequest
}

// NewAlibabadchainaoxiangwmsorderprocessreportRequest 初始化AlibabadchainaoxiangwmsorderprocessreportAPIRequest对象
func NewAlibabadchainaoxiangwmsorderprocessreportRequest() *AlibabadchainaoxiangwmsorderprocessreportAPIRequest {
	return &AlibabadchainaoxiangwmsorderprocessreportAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabadchainaoxiangwmsorderprocessreportAPIRequest) GetApiMethodName() string {
	return "alibaba.dchain.aoxiang.wms.orderprocess.report"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabadchainaoxiangwmsorderprocessreportAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabadchainaoxiangwmsorderprocessreportAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOrderProcessReportRequest is OrderProcessReportRequest Setter
// 回传单据流水请求
func (r *AlibabadchainaoxiangwmsorderprocessreportAPIRequest) SetOrderProcessReportRequest(_orderProcessReportRequest *OrderProcessReportRequest) error {
	r._orderProcessReportRequest = _orderProcessReportRequest
	r.Set("order_process_report_request", _orderProcessReportRequest)
	return nil
}

// GetOrderProcessReportRequest OrderProcessReportRequest Getter
func (r AlibabadchainaoxiangwmsorderprocessreportAPIRequest) GetOrderProcessReportRequest() *OrderProcessReportRequest {
	return r._orderProcessReportRequest
}
