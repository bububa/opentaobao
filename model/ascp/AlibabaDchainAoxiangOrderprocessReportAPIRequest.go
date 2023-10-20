package ascp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabadchainaoxiangorderprocessreportAPIRequest 回传仓内作业节点 API请求
// alibaba.dchain.aoxiang.orderprocess.report
//
// 回传仓内作业节点
type AlibabadchainaoxiangorderprocessreportAPIRequest struct {
	model.Params
	// 单据作业状态回传参数
	_orderprocessReportRequest *OrderProcessReportRequest
}

// NewAlibabadchainaoxiangorderprocessreportRequest 初始化AlibabadchainaoxiangorderprocessreportAPIRequest对象
func NewAlibabadchainaoxiangorderprocessreportRequest() *AlibabadchainaoxiangorderprocessreportAPIRequest {
	return &AlibabadchainaoxiangorderprocessreportAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabadchainaoxiangorderprocessreportAPIRequest) GetApiMethodName() string {
	return "alibaba.dchain.aoxiang.orderprocess.report"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabadchainaoxiangorderprocessreportAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabadchainaoxiangorderprocessreportAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOrderprocessReportRequest is OrderprocessReportRequest Setter
// 单据作业状态回传参数
func (r *AlibabadchainaoxiangorderprocessreportAPIRequest) SetOrderprocessReportRequest(_orderprocessReportRequest *OrderProcessReportRequest) error {
	r._orderprocessReportRequest = _orderprocessReportRequest
	r.Set("orderprocess_report_request", _orderprocessReportRequest)
	return nil
}

// GetOrderprocessReportRequest OrderprocessReportRequest Getter
func (r AlibabadchainaoxiangorderprocessreportAPIRequest) GetOrderprocessReportRequest() *OrderProcessReportRequest {
	return r._orderprocessReportRequest
}
