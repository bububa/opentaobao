package alihealthpw

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthPwApplynodeUpdateAPIRequest 申请节点变更回调 API请求
// alibaba.alihealth.pw.applynode.update
//
// 基金会回调阿里健康更新药品援助申请单的状态
type AlibabaAlihealthPwApplynodeUpdateAPIRequest struct {
	model.Params
	// 回调入参
	_body *AuditRollbackRo
}

// NewAlibabaAlihealthPwApplynodeUpdateRequest 初始化AlibabaAlihealthPwApplynodeUpdateAPIRequest对象
func NewAlibabaAlihealthPwApplynodeUpdateRequest() *AlibabaAlihealthPwApplynodeUpdateAPIRequest {
	return &AlibabaAlihealthPwApplynodeUpdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthPwApplynodeUpdateAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.pw.applynode.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthPwApplynodeUpdateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Body Setter
// 回调入参
func (r *AlibabaAlihealthPwApplynodeUpdateAPIRequest) SetBody(_body *AuditRollbackRo) error {
	r._body = _body
	r.Set("body", _body)
	return nil
}

// Get Body Getter
func (r AlibabaAlihealthPwApplynodeUpdateAPIRequest) GetBody() *AuditRollbackRo {
	return r._body
}
