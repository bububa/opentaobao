package tvupadmin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
迎客松节目设置免审开关 API请求
yunos.tvpubadmin.content.show.setexemptaudit

迎客松节目设置免审开关
*/
type YunosTvpubadminContentShowSetexemptauditRequest struct {
    model.Params
    // 节目longid
    _showLongId   int64
    // 牌照id：1CIBN，2WASU
    _license   int64
    // 牌照免审：1-开启节目免审，2-关闭节目免审
    _exemptAudit   int64
}

// 初始化YunosTvpubadminContentShowSetexemptauditRequest对象
func NewYunosTvpubadminContentShowSetexemptauditRequest() *YunosTvpubadminContentShowSetexemptauditRequest{
    return &YunosTvpubadminContentShowSetexemptauditRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r YunosTvpubadminContentShowSetexemptauditRequest) GetApiMethodName() string {
    return "yunos.tvpubadmin.content.show.setexemptaudit"
}

// IRequest interface 方法, 获取API参数
func (r YunosTvpubadminContentShowSetexemptauditRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ShowLongId Setter
// 节目longid
func (r *YunosTvpubadminContentShowSetexemptauditRequest) SetShowLongId(_showLongId int64) error {
    r._showLongId = _showLongId
    r.Set("show_long_id", _showLongId)
    return nil
}

// ShowLongId Getter
func (r YunosTvpubadminContentShowSetexemptauditRequest) GetShowLongId() int64 {
    return r._showLongId
}
// License Setter
// 牌照id：1CIBN，2WASU
func (r *YunosTvpubadminContentShowSetexemptauditRequest) SetLicense(_license int64) error {
    r._license = _license
    r.Set("license", _license)
    return nil
}

// License Getter
func (r YunosTvpubadminContentShowSetexemptauditRequest) GetLicense() int64 {
    return r._license
}
// ExemptAudit Setter
// 牌照免审：1-开启节目免审，2-关闭节目免审
func (r *YunosTvpubadminContentShowSetexemptauditRequest) SetExemptAudit(_exemptAudit int64) error {
    r._exemptAudit = _exemptAudit
    r.Set("exempt_audit", _exemptAudit)
    return nil
}

// ExemptAudit Getter
func (r YunosTvpubadminContentShowSetexemptauditRequest) GetExemptAudit() int64 {
    return r._exemptAudit
}
