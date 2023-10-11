package tmallservice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallSscWorkcardAcceptAPIRequest 服务商接单完成 API请求
// tmall.ssc.workcard.accept
//
// 工单完结
type TmallSscWorkcardAcceptAPIRequest struct {
	model.Params
	// 服务商接单入参
	_acceptWorkcardRequest *AcceptWorkcardRequest
}

// NewTmallSscWorkcardAcceptRequest 初始化TmallSscWorkcardAcceptAPIRequest对象
func NewTmallSscWorkcardAcceptRequest() *TmallSscWorkcardAcceptAPIRequest {
	return &TmallSscWorkcardAcceptAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallSscWorkcardAcceptAPIRequest) GetApiMethodName() string {
	return "tmall.ssc.workcard.accept"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallSscWorkcardAcceptAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallSscWorkcardAcceptAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAcceptWorkcardRequest is AcceptWorkcardRequest Setter
// 服务商接单入参
func (r *TmallSscWorkcardAcceptAPIRequest) SetAcceptWorkcardRequest(_acceptWorkcardRequest *AcceptWorkcardRequest) error {
	r._acceptWorkcardRequest = _acceptWorkcardRequest
	r.Set("accept_workcard_request", _acceptWorkcardRequest)
	return nil
}

// GetAcceptWorkcardRequest AcceptWorkcardRequest Getter
func (r TmallSscWorkcardAcceptAPIRequest) GetAcceptWorkcardRequest() *AcceptWorkcardRequest {
	return r._acceptWorkcardRequest
}
