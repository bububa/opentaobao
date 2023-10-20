package qimen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoQimenOrderexceptionReportAPIRequest 订单异常通知接口 API请求
// taobao.qimen.orderexception.report
//
// WMS调用奇门的接口,当WMS接收到ERP的发货指令时，由于种种原因（5.1.5说明了各种异常场景）可能无法完成发货。WMS通过调用此接口，通知ERP具体异常情况
type TaobaoQimenOrderexceptionReportAPIRequest struct {
	model.Params
	//
	_request *TaobaoQimenOrderexceptionReportRequest
}

// NewTaobaoQimenOrderexceptionReportRequest 初始化TaobaoQimenOrderexceptionReportAPIRequest对象
func NewTaobaoQimenOrderexceptionReportRequest() *TaobaoQimenOrderexceptionReportAPIRequest {
	return &TaobaoQimenOrderexceptionReportAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoQimenOrderexceptionReportAPIRequest) GetApiMethodName() string {
	return "taobao.qimen.orderexception.report"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoQimenOrderexceptionReportAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoQimenOrderexceptionReportAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRequest is Request Setter
func (r *TaobaoQimenOrderexceptionReportAPIRequest) SetRequest(_request *TaobaoQimenOrderexceptionReportRequest) error {
	r._request = _request
	r.Set("request", _request)
	return nil
}

// GetRequest Request Getter
func (r TaobaoQimenOrderexceptionReportAPIRequest) GetRequest() *TaobaoQimenOrderexceptionReportRequest {
	return r._request
}
