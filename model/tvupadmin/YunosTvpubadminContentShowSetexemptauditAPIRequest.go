package tvupadmin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// YunostvpubadmincontentshowsetexemptauditAPIRequest 迎客松节目设置免审开关 API请求
// yunos.tvpubadmin.content.show.setexemptaudit
//
// 迎客松节目设置免审开关
type YunostvpubadmincontentshowsetexemptauditAPIRequest struct {
	model.Params
	// 节目longid
	_showLongId int64
	// 牌照id：1CIBN，2WASU
	_license int64
	// 牌照免审：1-开启节目免审，2-关闭节目免审
	_exemptAudit int64
}

// NewYunostvpubadmincontentshowsetexemptauditRequest 初始化YunostvpubadmincontentshowsetexemptauditAPIRequest对象
func NewYunostvpubadmincontentshowsetexemptauditRequest() *YunostvpubadmincontentshowsetexemptauditAPIRequest {
	return &YunostvpubadmincontentshowsetexemptauditAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunostvpubadmincontentshowsetexemptauditAPIRequest) GetApiMethodName() string {
	return "yunos.tvpubadmin.content.show.setexemptaudit"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunostvpubadmincontentshowsetexemptauditAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YunostvpubadmincontentshowsetexemptauditAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetShowLongId is ShowLongId Setter
// 节目longid
func (r *YunostvpubadmincontentshowsetexemptauditAPIRequest) SetShowLongId(_showLongId int64) error {
	r._showLongId = _showLongId
	r.Set("show_long_id", _showLongId)
	return nil
}

// GetShowLongId ShowLongId Getter
func (r YunostvpubadmincontentshowsetexemptauditAPIRequest) GetShowLongId() int64 {
	return r._showLongId
}

// SetLicense is License Setter
// 牌照id：1CIBN，2WASU
func (r *YunostvpubadmincontentshowsetexemptauditAPIRequest) SetLicense(_license int64) error {
	r._license = _license
	r.Set("license", _license)
	return nil
}

// GetLicense License Getter
func (r YunostvpubadmincontentshowsetexemptauditAPIRequest) GetLicense() int64 {
	return r._license
}

// SetExemptAudit is ExemptAudit Setter
// 牌照免审：1-开启节目免审，2-关闭节目免审
func (r *YunostvpubadmincontentshowsetexemptauditAPIRequest) SetExemptAudit(_exemptAudit int64) error {
	r._exemptAudit = _exemptAudit
	r.Set("exempt_audit", _exemptAudit)
	return nil
}

// GetExemptAudit ExemptAudit Getter
func (r YunostvpubadmincontentshowsetexemptauditAPIRequest) GetExemptAudit() int64 {
	return r._exemptAudit
}
