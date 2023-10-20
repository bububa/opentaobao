package tmallgenie

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmalltmjlappsapserviceordercancelAPIRequest 取消售后服务单 API请求
// tmall.tmjlapp.sap.serviceorder.cancel
//
// SAP跟天猫精灵app接口对接，用户在app取消sap售后服务工单
type TmalltmjlappsapserviceordercancelAPIRequest struct {
	model.Params
	// 取消服务单请求
	_cancelRequest *Dtcancelrequest
}

// NewTmalltmjlappsapserviceordercancelRequest 初始化TmalltmjlappsapserviceordercancelAPIRequest对象
func NewTmalltmjlappsapserviceordercancelRequest() *TmalltmjlappsapserviceordercancelAPIRequest {
	return &TmalltmjlappsapserviceordercancelAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmalltmjlappsapserviceordercancelAPIRequest) GetApiMethodName() string {
	return "tmall.tmjlapp.sap.serviceorder.cancel"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmalltmjlappsapserviceordercancelAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmalltmjlappsapserviceordercancelAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCancelRequest is CancelRequest Setter
// 取消服务单请求
func (r *TmalltmjlappsapserviceordercancelAPIRequest) SetCancelRequest(_cancelRequest *Dtcancelrequest) error {
	r._cancelRequest = _cancelRequest
	r.Set("cancel_request", _cancelRequest)
	return nil
}

// GetCancelRequest CancelRequest Getter
func (r TmalltmjlappsapserviceordercancelAPIRequest) GetCancelRequest() *Dtcancelrequest {
	return r._cancelRequest
}
