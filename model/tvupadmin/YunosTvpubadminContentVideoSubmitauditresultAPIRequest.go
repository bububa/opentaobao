package tvupadmin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// YunostvpubadmincontentvideosubmitauditresultAPIRequest 迎客松提交视频审核结果 API请求
// yunos.tvpubadmin.content.video.submitauditresult
//
// 迎客松提交视频审核结果
type YunostvpubadmincontentvideosubmitauditresultAPIRequest struct {
	model.Params
	// 审核结果，json格式
	_licenseAuditResult string
}

// NewYunostvpubadmincontentvideosubmitauditresultRequest 初始化YunostvpubadmincontentvideosubmitauditresultAPIRequest对象
func NewYunostvpubadmincontentvideosubmitauditresultRequest() *YunostvpubadmincontentvideosubmitauditresultAPIRequest {
	return &YunostvpubadmincontentvideosubmitauditresultAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunostvpubadmincontentvideosubmitauditresultAPIRequest) GetApiMethodName() string {
	return "yunos.tvpubadmin.content.video.submitauditresult"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunostvpubadmincontentvideosubmitauditresultAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YunostvpubadmincontentvideosubmitauditresultAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetLicenseAuditResult is LicenseAuditResult Setter
// 审核结果，json格式
func (r *YunostvpubadmincontentvideosubmitauditresultAPIRequest) SetLicenseAuditResult(_licenseAuditResult string) error {
	r._licenseAuditResult = _licenseAuditResult
	r.Set("license_audit_result", _licenseAuditResult)
	return nil
}

// GetLicenseAuditResult LicenseAuditResult Getter
func (r YunostvpubadmincontentvideosubmitauditresultAPIRequest) GetLicenseAuditResult() string {
	return r._licenseAuditResult
}
