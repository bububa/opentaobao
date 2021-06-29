package tvupadmin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
更新系统版本审核状态 APIRequest
yunos.tvpubadmin.device.updateosstatus

更新系统版本审核状态
*/
type YunosTvpubadminDeviceUpdateosstatusRequest struct {
    model.Params

    // 升级ID
    versionId   int64 

    // 牌照方
    license   int64 

    // 审核状态
    status   string 

    // 审核意见
    auditComment   string 

}

func NewYunosTvpubadminDeviceUpdateosstatusRequest() *YunosTvpubadminDeviceUpdateosstatusRequest{
    return &YunosTvpubadminDeviceUpdateosstatusRequest{
        Params: model.NewParams(),
    }
}

func (r YunosTvpubadminDeviceUpdateosstatusRequest) GetApiMethodName() string {
    return "yunos.tvpubadmin.device.updateosstatus"
}

func (r YunosTvpubadminDeviceUpdateosstatusRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *YunosTvpubadminDeviceUpdateosstatusRequest) SetVersionId(versionId int64) error {
    r.versionId = versionId
    r.Set("version_id", versionId)
    return nil
}

func (r YunosTvpubadminDeviceUpdateosstatusRequest) GetVersionId() int64 {
    return r.versionId
}

func (r *YunosTvpubadminDeviceUpdateosstatusRequest) SetLicense(license int64) error {
    r.license = license
    r.Set("license", license)
    return nil
}

func (r YunosTvpubadminDeviceUpdateosstatusRequest) GetLicense() int64 {
    return r.license
}

func (r *YunosTvpubadminDeviceUpdateosstatusRequest) SetStatus(status string) error {
    r.status = status
    r.Set("status", status)
    return nil
}

func (r YunosTvpubadminDeviceUpdateosstatusRequest) GetStatus() string {
    return r.status
}

func (r *YunosTvpubadminDeviceUpdateosstatusRequest) SetAuditComment(auditComment string) error {
    r.auditComment = auditComment
    r.Set("audit_comment", auditComment)
    return nil
}

func (r YunosTvpubadminDeviceUpdateosstatusRequest) GetAuditComment() string {
    return r.auditComment
}

