package alihealthpw

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthpwgmdetailAPIRequest 同情用药申请单详情接口 API请求
// alibaba.alihealth.pw.gm.detail
//
// 同情用药申请单详情接口，提供给合作方查询申请单详情
type AlibabaalihealthpwgmdetailAPIRequest struct {
	model.Params
	// 入参
	_body *DetailForBreq
}

// NewAlibabaalihealthpwgmdetailRequest 初始化AlibabaalihealthpwgmdetailAPIRequest对象
func NewAlibabaalihealthpwgmdetailRequest() *AlibabaalihealthpwgmdetailAPIRequest {
	return &AlibabaalihealthpwgmdetailAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthpwgmdetailAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.pw.gm.detail"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthpwgmdetailAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthpwgmdetailAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBody is Body Setter
// 入参
func (r *AlibabaalihealthpwgmdetailAPIRequest) SetBody(_body *DetailForBreq) error {
	r._body = _body
	r.Set("body", _body)
	return nil
}

// GetBody Body Getter
func (r AlibabaalihealthpwgmdetailAPIRequest) GetBody() *DetailForBreq {
	return r._body
}
