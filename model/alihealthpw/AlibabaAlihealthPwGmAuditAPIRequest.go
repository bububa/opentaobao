package alihealthpw

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthPwGmAuditAPIRequest 同情用药审核接口 API请求
// alibaba.alihealth.pw.gm.audit
//
// 同情用药审核接口，提供给合作方审核申请单
type AlibabaAlihealthPwGmAuditAPIRequest struct {
	model.Params
	// 入参
	_body *AuditReq
}

// NewAlibabaAlihealthPwGmAuditRequest 初始化AlibabaAlihealthPwGmAuditAPIRequest对象
func NewAlibabaAlihealthPwGmAuditRequest() *AlibabaAlihealthPwGmAuditAPIRequest {
	return &AlibabaAlihealthPwGmAuditAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthPwGmAuditAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.pw.gm.audit"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthPwGmAuditAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetBody is Body Setter
// 入参
func (r *AlibabaAlihealthPwGmAuditAPIRequest) SetBody(_body *AuditReq) error {
	r._body = _body
	r.Set("body", _body)
	return nil
}

// GetBody Body Getter
func (r AlibabaAlihealthPwGmAuditAPIRequest) GetBody() *AuditReq {
	return r._body
}
