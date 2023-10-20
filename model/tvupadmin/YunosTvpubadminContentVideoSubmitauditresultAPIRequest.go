package tvupadmin

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// YunosTvpubadminContentVideoSubmitauditresultAPIRequest 迎客松提交视频审核结果 API请求
// yunos.tvpubadmin.content.video.submitauditresult
//
// 迎客松提交视频审核结果
type YunosTvpubadminContentVideoSubmitauditresultAPIRequest struct {
	model.Params
	// 审核结果，json格式
	_licenseAuditResult string
}

// NewYunosTvpubadminContentVideoSubmitauditresultRequest 初始化YunosTvpubadminContentVideoSubmitauditresultAPIRequest对象
func NewYunosTvpubadminContentVideoSubmitauditresultRequest() *YunosTvpubadminContentVideoSubmitauditresultAPIRequest {
	return &YunosTvpubadminContentVideoSubmitauditresultAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *YunosTvpubadminContentVideoSubmitauditresultAPIRequest) Reset() {
	r._licenseAuditResult = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunosTvpubadminContentVideoSubmitauditresultAPIRequest) GetApiMethodName() string {
	return "yunos.tvpubadmin.content.video.submitauditresult"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunosTvpubadminContentVideoSubmitauditresultAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YunosTvpubadminContentVideoSubmitauditresultAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetLicenseAuditResult is LicenseAuditResult Setter
// 审核结果，json格式
func (r *YunosTvpubadminContentVideoSubmitauditresultAPIRequest) SetLicenseAuditResult(_licenseAuditResult string) error {
	r._licenseAuditResult = _licenseAuditResult
	r.Set("license_audit_result", _licenseAuditResult)
	return nil
}

// GetLicenseAuditResult LicenseAuditResult Getter
func (r YunosTvpubadminContentVideoSubmitauditresultAPIRequest) GetLicenseAuditResult() string {
	return r._licenseAuditResult
}

var poolYunosTvpubadminContentVideoSubmitauditresultAPIRequest = sync.Pool{
	New: func() any {
		return NewYunosTvpubadminContentVideoSubmitauditresultRequest()
	},
}

// GetYunosTvpubadminContentVideoSubmitauditresultRequest 从 sync.Pool 获取 YunosTvpubadminContentVideoSubmitauditresultAPIRequest
func GetYunosTvpubadminContentVideoSubmitauditresultAPIRequest() *YunosTvpubadminContentVideoSubmitauditresultAPIRequest {
	return poolYunosTvpubadminContentVideoSubmitauditresultAPIRequest.Get().(*YunosTvpubadminContentVideoSubmitauditresultAPIRequest)
}

// ReleaseYunosTvpubadminContentVideoSubmitauditresultAPIRequest 将 YunosTvpubadminContentVideoSubmitauditresultAPIRequest 放入 sync.Pool
func ReleaseYunosTvpubadminContentVideoSubmitauditresultAPIRequest(v *YunosTvpubadminContentVideoSubmitauditresultAPIRequest) {
	v.Reset()
	poolYunosTvpubadminContentVideoSubmitauditresultAPIRequest.Put(v)
}
