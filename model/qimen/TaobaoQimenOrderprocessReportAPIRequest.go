package qimen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoqimenorderprocessreportAPIRequest 订单流水通知接口 API请求
// taobao.qimen.orderprocess.report
//
// taobao.qimen.orderprocess.report
type TaobaoqimenorderprocessreportAPIRequest struct {
	model.Params
	//
	_request *OrderProcessReportRequest
}

// NewTaobaoqimenorderprocessreportRequest 初始化TaobaoqimenorderprocessreportAPIRequest对象
func NewTaobaoqimenorderprocessreportRequest() *TaobaoqimenorderprocessreportAPIRequest {
	return &TaobaoqimenorderprocessreportAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoqimenorderprocessreportAPIRequest) GetApiMethodName() string {
	return "taobao.qimen.orderprocess.report"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoqimenorderprocessreportAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoqimenorderprocessreportAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRequest is Request Setter
func (r *TaobaoqimenorderprocessreportAPIRequest) SetRequest(_request *OrderProcessReportRequest) error {
	r._request = _request
	r.Set("request", _request)
	return nil
}

// GetRequest Request Getter
func (r TaobaoqimenorderprocessreportAPIRequest) GetRequest() *OrderProcessReportRequest {
	return r._request
}
