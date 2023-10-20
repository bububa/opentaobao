package tmallservice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallsscworkcardacceptAPIRequest 服务商接单完成 API请求
// tmall.ssc.workcard.accept
//
// 工单完结
type TmallsscworkcardacceptAPIRequest struct {
	model.Params
	// 服务商接单入参
	_acceptWorkcardRequest *AcceptWorkcardRequest
}

// NewTmallsscworkcardacceptRequest 初始化TmallsscworkcardacceptAPIRequest对象
func NewTmallsscworkcardacceptRequest() *TmallsscworkcardacceptAPIRequest {
	return &TmallsscworkcardacceptAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallsscworkcardacceptAPIRequest) GetApiMethodName() string {
	return "tmall.ssc.workcard.accept"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallsscworkcardacceptAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallsscworkcardacceptAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAcceptWorkcardRequest is AcceptWorkcardRequest Setter
// 服务商接单入参
func (r *TmallsscworkcardacceptAPIRequest) SetAcceptWorkcardRequest(_acceptWorkcardRequest *AcceptWorkcardRequest) error {
	r._acceptWorkcardRequest = _acceptWorkcardRequest
	r.Set("accept_workcard_request", _acceptWorkcardRequest)
	return nil
}

// GetAcceptWorkcardRequest AcceptWorkcardRequest Getter
func (r TmallsscworkcardacceptAPIRequest) GetAcceptWorkcardRequest() *AcceptWorkcardRequest {
	return r._acceptWorkcardRequest
}
