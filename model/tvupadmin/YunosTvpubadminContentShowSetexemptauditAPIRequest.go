package tvupadmin

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// YunosTvpubadminContentShowSetexemptauditAPIRequest 迎客松节目设置免审开关 API请求
// yunos.tvpubadmin.content.show.setexemptaudit
//
// 迎客松节目设置免审开关
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
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *YunosTvpubadminContentShowSetexemptauditAPIRequest) Reset() {
	r._showLongId = 0
	r._license = 0
	r._exemptAudit = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunosTvpubadminContentShowSetexemptauditAPIRequest) GetApiMethodName() string {
	return "yunos.tvpubadmin.content.show.setexemptaudit"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunosTvpubadminContentShowSetexemptauditAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YunosTvpubadminContentShowSetexemptauditAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetShowLongId is ShowLongId Setter
// 节目longid
func (r *YunosTvpubadminContentShowSetexemptauditAPIRequest) SetShowLongId(_showLongId int64) error {
	r._showLongId = _showLongId
	r.Set("show_long_id", _showLongId)
	return nil
}

// GetShowLongId ShowLongId Getter
func (r YunosTvpubadminContentShowSetexemptauditAPIRequest) GetShowLongId() int64 {
	return r._showLongId
}

// SetLicense is License Setter
// 牌照id：1CIBN，2WASU
func (r *YunosTvpubadminContentShowSetexemptauditAPIRequest) SetLicense(_license int64) error {
	r._license = _license
	r.Set("license", _license)
	return nil
}

// GetLicense License Getter
func (r YunosTvpubadminContentShowSetexemptauditAPIRequest) GetLicense() int64 {
	return r._license
}

// SetExemptAudit is ExemptAudit Setter
// 牌照免审：1-开启节目免审，2-关闭节目免审
func (r *YunosTvpubadminContentShowSetexemptauditAPIRequest) SetExemptAudit(_exemptAudit int64) error {
	r._exemptAudit = _exemptAudit
	r.Set("exempt_audit", _exemptAudit)
	return nil
}

// GetExemptAudit ExemptAudit Getter
func (r YunosTvpubadminContentShowSetexemptauditAPIRequest) GetExemptAudit() int64 {
	return r._exemptAudit
}

var poolYunosTvpubadminContentShowSetexemptauditAPIRequest = sync.Pool{
	New: func() any {
		return NewYunosTvpubadminContentShowSetexemptauditRequest()
	},
}

// GetYunosTvpubadminContentShowSetexemptauditRequest 从 sync.Pool 获取 YunosTvpubadminContentShowSetexemptauditAPIRequest
func GetYunosTvpubadminContentShowSetexemptauditAPIRequest() *YunosTvpubadminContentShowSetexemptauditAPIRequest {
	return poolYunosTvpubadminContentShowSetexemptauditAPIRequest.Get().(*YunosTvpubadminContentShowSetexemptauditAPIRequest)
}

// ReleaseYunosTvpubadminContentShowSetexemptauditAPIRequest 将 YunosTvpubadminContentShowSetexemptauditAPIRequest 放入 sync.Pool
func ReleaseYunosTvpubadminContentShowSetexemptauditAPIRequest(v *YunosTvpubadminContentShowSetexemptauditAPIRequest) {
	v.Reset()
	poolYunosTvpubadminContentShowSetexemptauditAPIRequest.Put(v)
}
