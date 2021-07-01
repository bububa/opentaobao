package tvupadmin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* YunosTvpubadminContentShowSetexemptauditAPIRequest
迎客松节目设置免审开关 API请求
yunos.tvpubadmin.content.show.setexemptaudit

迎客松节目设置免审开关 */
type YunosTvpubadminContentShowSetexemptauditAPIRequest struct {
	model.Params
	// 节目longid
	_showLongId int64
	// 牌照id：1CIBN，2WASU
	_license int64
	// 牌照免审：1-开启节目免审，2-关闭节目免审
	_exemptAudit int64
}

// NewYunosTvpubadminContentShowSetexemptauditRequest 初始化YunosTvpubadminContentShowSetexemptauditAPIRequest对象
func NewYunosTvpubadminContentShowSetexemptauditRequest() *YunosTvpubadminContentShowSetexemptauditAPIRequest {
	return &YunosTvpubadminContentShowSetexemptauditAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunosTvpubadminContentShowSetexemptauditAPIRequest) GetApiMethodName() string {
	return "yunos.tvpubadmin.content.show.setexemptaudit"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunosTvpubadminContentShowSetexemptauditAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is ShowLongId Setter
// 节目longid
func (r *YunosTvpubadminContentShowSetexemptauditAPIRequest) SetShowLongId(_showLongId int64) error {
	r._showLongId = _showLongId
	r.Set("show_long_id", _showLongId)
	return nil
}

// Get ShowLongId Getter
func (r YunosTvpubadminContentShowSetexemptauditAPIRequest) GetShowLongId() int64 {
	return r._showLongId
}

// Set is License Setter
// 牌照id：1CIBN，2WASU
func (r *YunosTvpubadminContentShowSetexemptauditAPIRequest) SetLicense(_license int64) error {
	r._license = _license
	r.Set("license", _license)
	return nil
}

// Get License Getter
func (r YunosTvpubadminContentShowSetexemptauditAPIRequest) GetLicense() int64 {
	return r._license
}

// Set is ExemptAudit Setter
// 牌照免审：1-开启节目免审，2-关闭节目免审
func (r *YunosTvpubadminContentShowSetexemptauditAPIRequest) SetExemptAudit(_exemptAudit int64) error {
	r._exemptAudit = _exemptAudit
	r.Set("exempt_audit", _exemptAudit)
	return nil
}

// Get ExemptAudit Getter
func (r YunosTvpubadminContentShowSetexemptauditAPIRequest) GetExemptAudit() int64 {
	return r._exemptAudit
}
