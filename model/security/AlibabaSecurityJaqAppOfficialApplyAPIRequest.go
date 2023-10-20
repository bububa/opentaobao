package security

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaSecurityJaqAppOfficialApplyAPIRequest 聚安全官方应用申请 API请求
// alibaba.security.jaq.app.official.apply
//
// 官方应用申请接口
type AlibabaSecurityJaqAppOfficialApplyAPIRequest struct {
	model.Params
	// 官方应用申请入参
	_officialAppApplyRequest *OfficialAppApplyRequest
}

// NewAlibabaSecurityJaqAppOfficialApplyRequest 初始化AlibabaSecurityJaqAppOfficialApplyAPIRequest对象
func NewAlibabaSecurityJaqAppOfficialApplyRequest() *AlibabaSecurityJaqAppOfficialApplyAPIRequest {
	return &AlibabaSecurityJaqAppOfficialApplyAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaSecurityJaqAppOfficialApplyAPIRequest) Reset() {
	r._officialAppApplyRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaSecurityJaqAppOfficialApplyAPIRequest) GetApiMethodName() string {
	return "alibaba.security.jaq.app.official.apply"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaSecurityJaqAppOfficialApplyAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaSecurityJaqAppOfficialApplyAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOfficialAppApplyRequest is OfficialAppApplyRequest Setter
// 官方应用申请入参
func (r *AlibabaSecurityJaqAppOfficialApplyAPIRequest) SetOfficialAppApplyRequest(_officialAppApplyRequest *OfficialAppApplyRequest) error {
	r._officialAppApplyRequest = _officialAppApplyRequest
	r.Set("official_app_apply_request", _officialAppApplyRequest)
	return nil
}

// GetOfficialAppApplyRequest OfficialAppApplyRequest Getter
func (r AlibabaSecurityJaqAppOfficialApplyAPIRequest) GetOfficialAppApplyRequest() *OfficialAppApplyRequest {
	return r._officialAppApplyRequest
}

var poolAlibabaSecurityJaqAppOfficialApplyAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaSecurityJaqAppOfficialApplyRequest()
	},
}

// GetAlibabaSecurityJaqAppOfficialApplyRequest 从 sync.Pool 获取 AlibabaSecurityJaqAppOfficialApplyAPIRequest
func GetAlibabaSecurityJaqAppOfficialApplyAPIRequest() *AlibabaSecurityJaqAppOfficialApplyAPIRequest {
	return poolAlibabaSecurityJaqAppOfficialApplyAPIRequest.Get().(*AlibabaSecurityJaqAppOfficialApplyAPIRequest)
}

// ReleaseAlibabaSecurityJaqAppOfficialApplyAPIRequest 将 AlibabaSecurityJaqAppOfficialApplyAPIRequest 放入 sync.Pool
func ReleaseAlibabaSecurityJaqAppOfficialApplyAPIRequest(v *AlibabaSecurityJaqAppOfficialApplyAPIRequest) {
	v.Reset()
	poolAlibabaSecurityJaqAppOfficialApplyAPIRequest.Put(v)
}
