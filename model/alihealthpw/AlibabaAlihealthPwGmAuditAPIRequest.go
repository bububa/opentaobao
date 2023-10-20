package alihealthpw

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthpwgmauditAPIRequest 同情用药审核接口 API请求
// alibaba.alihealth.pw.gm.audit
//
// 同情用药审核接口，提供给合作方审核申请单
type AlibabaalihealthpwgmauditAPIRequest struct {
	model.Params
	// 入参
	_body *AuditReq
}

// NewAlibabaalihealthpwgmauditRequest 初始化AlibabaalihealthpwgmauditAPIRequest对象
func NewAlibabaalihealthpwgmauditRequest() *AlibabaalihealthpwgmauditAPIRequest {
	return &AlibabaalihealthpwgmauditAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthpwgmauditAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.pw.gm.audit"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthpwgmauditAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthpwgmauditAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBody is Body Setter
// 入参
func (r *AlibabaalihealthpwgmauditAPIRequest) SetBody(_body *AuditReq) error {
	r._body = _body
	r.Set("body", _body)
	return nil
}

// GetBody Body Getter
func (r AlibabaalihealthpwgmauditAPIRequest) GetBody() *AuditReq {
	return r._body
}
