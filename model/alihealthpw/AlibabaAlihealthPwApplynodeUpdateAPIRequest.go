package alihealthpw

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthpwapplynodeupdateAPIRequest 申请节点变更回调 API请求
// alibaba.alihealth.pw.applynode.update
//
// 基金会回调阿里健康更新药品援助申请单的状态
type AlibabaalihealthpwapplynodeupdateAPIRequest struct {
	model.Params
	// 回调入参
	_body *AuditRollbackRo
}

// NewAlibabaalihealthpwapplynodeupdateRequest 初始化AlibabaalihealthpwapplynodeupdateAPIRequest对象
func NewAlibabaalihealthpwapplynodeupdateRequest() *AlibabaalihealthpwapplynodeupdateAPIRequest {
	return &AlibabaalihealthpwapplynodeupdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthpwapplynodeupdateAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.pw.applynode.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthpwapplynodeupdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthpwapplynodeupdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBody is Body Setter
// 回调入参
func (r *AlibabaalihealthpwapplynodeupdateAPIRequest) SetBody(_body *AuditRollbackRo) error {
	r._body = _body
	r.Set("body", _body)
	return nil
}

// GetBody Body Getter
func (r AlibabaalihealthpwapplynodeupdateAPIRequest) GetBody() *AuditRollbackRo {
	return r._body
}
