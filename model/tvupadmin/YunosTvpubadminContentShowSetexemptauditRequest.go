package tvupadmin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
迎客松节目设置免审开关 APIRequest
yunos.tvpubadmin.content.show.setexemptaudit

迎客松节目设置免审开关
*/
type YunosTvpubadminContentShowSetexemptauditRequest struct {
    model.Params

    // 节目longid
    showLongId   int64 

    // 牌照id：1CIBN，2WASU
    license   int64 

    // 牌照免审：1-开启节目免审，2-关闭节目免审
    exemptAudit   int64 

}

func NewYunosTvpubadminContentShowSetexemptauditRequest() *YunosTvpubadminContentShowSetexemptauditRequest{
    return &YunosTvpubadminContentShowSetexemptauditRequest{
        Params: model.NewParams(),
    }
}

func (r YunosTvpubadminContentShowSetexemptauditRequest) GetApiMethodName() string {
    return "yunos.tvpubadmin.content.show.setexemptaudit"
}

func (r YunosTvpubadminContentShowSetexemptauditRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *YunosTvpubadminContentShowSetexemptauditRequest) SetShowLongId(showLongId int64) error {
    r.showLongId = showLongId
    r.Set("show_long_id", showLongId)
    return nil
}

func (r YunosTvpubadminContentShowSetexemptauditRequest) GetShowLongId() int64 {
    return r.showLongId
}

func (r *YunosTvpubadminContentShowSetexemptauditRequest) SetLicense(license int64) error {
    r.license = license
    r.Set("license", license)
    return nil
}

func (r YunosTvpubadminContentShowSetexemptauditRequest) GetLicense() int64 {
    return r.license
}

func (r *YunosTvpubadminContentShowSetexemptauditRequest) SetExemptAudit(exemptAudit int64) error {
    r.exemptAudit = exemptAudit
    r.Set("exempt_audit", exemptAudit)
    return nil
}

func (r YunosTvpubadminContentShowSetexemptauditRequest) GetExemptAudit() int64 {
    return r.exemptAudit
}

