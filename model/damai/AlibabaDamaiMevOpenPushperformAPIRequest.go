package damai

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabadamaimevopenpushperformAPIRequest 大麦换验平台-第三方对外开放-场次接口pushPerform API请求
// alibaba.damai.mev.open.pushperform
//
// pushPerform
type AlibabadamaimevopenpushperformAPIRequest struct {
	model.Params
	// 入参pushPerformParam
	_pushPerformParam *ThirdPerformPushOpenParam
}

// NewAlibabadamaimevopenpushperformRequest 初始化AlibabadamaimevopenpushperformAPIRequest对象
func NewAlibabadamaimevopenpushperformRequest() *AlibabadamaimevopenpushperformAPIRequest {
	return &AlibabadamaimevopenpushperformAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabadamaimevopenpushperformAPIRequest) GetApiMethodName() string {
	return "alibaba.damai.mev.open.pushperform"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabadamaimevopenpushperformAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabadamaimevopenpushperformAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPushPerformParam is PushPerformParam Setter
// 入参pushPerformParam
func (r *AlibabadamaimevopenpushperformAPIRequest) SetPushPerformParam(_pushPerformParam *ThirdPerformPushOpenParam) error {
	r._pushPerformParam = _pushPerformParam
	r.Set("push_perform_param", _pushPerformParam)
	return nil
}

// GetPushPerformParam PushPerformParam Getter
func (r AlibabadamaimevopenpushperformAPIRequest) GetPushPerformParam() *ThirdPerformPushOpenParam {
	return r._pushPerformParam
}
