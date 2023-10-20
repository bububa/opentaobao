package ascp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabadchainisvwmsorderprocessreportAPIRequest 仓作业信息同步 API请求
// alibaba.dchain.isv.wms.orderprocess.report
//
// 仓作业信息同步
type AlibabadchainisvwmsorderprocessreportAPIRequest struct {
	model.Params
	// 请求入参
	_orderProcessReportRequest *WmsOrderProcessReportRequest
}

// NewAlibabadchainisvwmsorderprocessreportRequest 初始化AlibabadchainisvwmsorderprocessreportAPIRequest对象
func NewAlibabadchainisvwmsorderprocessreportRequest() *AlibabadchainisvwmsorderprocessreportAPIRequest {
	return &AlibabadchainisvwmsorderprocessreportAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabadchainisvwmsorderprocessreportAPIRequest) GetApiMethodName() string {
	return "alibaba.dchain.isv.wms.orderprocess.report"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabadchainisvwmsorderprocessreportAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabadchainisvwmsorderprocessreportAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOrderProcessReportRequest is OrderProcessReportRequest Setter
// 请求入参
func (r *AlibabadchainisvwmsorderprocessreportAPIRequest) SetOrderProcessReportRequest(_orderProcessReportRequest *WmsOrderProcessReportRequest) error {
	r._orderProcessReportRequest = _orderProcessReportRequest
	r.Set("order_process_report_request", _orderProcessReportRequest)
	return nil
}

// GetOrderProcessReportRequest OrderProcessReportRequest Getter
func (r AlibabadchainisvwmsorderprocessreportAPIRequest) GetOrderProcessReportRequest() *WmsOrderProcessReportRequest {
	return r._orderProcessReportRequest
}
