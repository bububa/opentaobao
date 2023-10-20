package qimen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoQimenOrderprocessReportAPIRequest 订单流水通知接口 API请求
// taobao.qimen.orderprocess.report
//
// taobao.qimen.orderprocess.report
type TaobaoQimenOrderprocessReportAPIRequest struct {
	model.Params
	//
	_request *OrderProcessReportRequest
}

// NewTaobaoQimenOrderprocessReportRequest 初始化TaobaoQimenOrderprocessReportAPIRequest对象
func NewTaobaoQimenOrderprocessReportRequest() *TaobaoQimenOrderprocessReportAPIRequest {
	return &TaobaoQimenOrderprocessReportAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoQimenOrderprocessReportAPIRequest) GetApiMethodName() string {
	return "taobao.qimen.orderprocess.report"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoQimenOrderprocessReportAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoQimenOrderprocessReportAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRequest is Request Setter
func (r *TaobaoQimenOrderprocessReportAPIRequest) SetRequest(_request *OrderProcessReportRequest) error {
	r._request = _request
	r.Set("request", _request)
	return nil
}

// GetRequest Request Getter
func (r TaobaoQimenOrderprocessReportAPIRequest) GetRequest() *OrderProcessReportRequest {
	return r._request
}
