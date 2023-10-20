package alihealthpw

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihealthPwApplynodeUpdateAPIRequest) Reset() {
	r._body = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthPwApplynodeUpdateAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.pw.applynode.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthPwApplynodeUpdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthPwApplynodeUpdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBody is Body Setter
// 回调入参
func (r *AlibabaAlihealthPwApplynodeUpdateAPIRequest) SetBody(_body *AuditRollbackRo) error {
	r._body = _body
	r.Set("body", _body)
	return nil
}

// GetBody Body Getter
func (r AlibabaAlihealthPwApplynodeUpdateAPIRequest) GetBody() *AuditRollbackRo {
	return r._body
}

var poolAlibabaAlihealthPwApplynodeUpdateAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihealthPwApplynodeUpdateRequest()
	},
}

// GetAlibabaAlihealthPwApplynodeUpdateRequest 从 sync.Pool 获取 AlibabaAlihealthPwApplynodeUpdateAPIRequest
func GetAlibabaAlihealthPwApplynodeUpdateAPIRequest() *AlibabaAlihealthPwApplynodeUpdateAPIRequest {
	return poolAlibabaAlihealthPwApplynodeUpdateAPIRequest.Get().(*AlibabaAlihealthPwApplynodeUpdateAPIRequest)
}

// ReleaseAlibabaAlihealthPwApplynodeUpdateAPIRequest 将 AlibabaAlihealthPwApplynodeUpdateAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihealthPwApplynodeUpdateAPIRequest(v *AlibabaAlihealthPwApplynodeUpdateAPIRequest) {
	v.Reset()
	poolAlibabaAlihealthPwApplynodeUpdateAPIRequest.Put(v)
}
