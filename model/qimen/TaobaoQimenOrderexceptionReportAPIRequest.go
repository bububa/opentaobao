package qimen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoqimenorderexceptionreportAPIRequest 订单异常通知接口 API请求
// taobao.qimen.orderexception.report
//
// WMS调用奇门的接口,当WMS接收到ERP的发货指令时，由于种种原因（5.1.5说明了各种异常场景）可能无法完成发货。WMS通过调用此接口，通知ERP具体异常情况
type TaobaoqimenorderexceptionreportAPIRequest struct {
	model.Params
	//
	_request *TaobaoqimenorderexceptionreportRequest
}

// NewTaobaoqimenorderexceptionreportRequest 初始化TaobaoqimenorderexceptionreportAPIRequest对象
func NewTaobaoqimenorderexceptionreportRequest() *TaobaoqimenorderexceptionreportAPIRequest {
	return &TaobaoqimenorderexceptionreportAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoqimenorderexceptionreportAPIRequest) GetApiMethodName() string {
	return "taobao.qimen.orderexception.report"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoqimenorderexceptionreportAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoqimenorderexceptionreportAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRequest is Request Setter
func (r *TaobaoqimenorderexceptionreportAPIRequest) SetRequest(_request *TaobaoqimenorderexceptionreportRequest) error {
	r._request = _request
	r.Set("request", _request)
	return nil
}

// GetRequest Request Getter
func (r TaobaoqimenorderexceptionreportAPIRequest) GetRequest() *TaobaoqimenorderexceptionreportRequest {
	return r._request
}
