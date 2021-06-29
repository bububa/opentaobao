package tvupadmin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
更新应用版本审核状态 APIRequest
yunos.tvpubadmin.device.updateappstatus

更新应用版本审核状态
*/
type YunosTvpubadminDeviceUpdateappstatusRequest struct {
    model.Params

    // 应用ID
    versionId   int64 

    // 牌照方
    license   int64 

    // 审核状态
    status   string 

    // 审核意见
    auditComment   string 

}

func NewYunosTvpubadminDeviceUpdateappstatusRequest() *YunosTvpubadminDeviceUpdateappstatusRequest{
    return &YunosTvpubadminDeviceUpdateappstatusRequest{
        Params: model.NewParams(),
    }
}

func (r YunosTvpubadminDeviceUpdateappstatusRequest) GetApiMethodName() string {
    return "yunos.tvpubadmin.device.updateappstatus"
}

func (r YunosTvpubadminDeviceUpdateappstatusRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *YunosTvpubadminDeviceUpdateappstatusRequest) SetVersionId(versionId int64) error {
    r.versionId = versionId
    r.Set("version_id", versionId)
    return nil
}

func (r YunosTvpubadminDeviceUpdateappstatusRequest) GetVersionId() int64 {
    return r.versionId
}

func (r *YunosTvpubadminDeviceUpdateappstatusRequest) SetLicense(license int64) error {
    r.license = license
    r.Set("license", license)
    return nil
}

func (r YunosTvpubadminDeviceUpdateappstatusRequest) GetLicense() int64 {
    return r.license
}

func (r *YunosTvpubadminDeviceUpdateappstatusRequest) SetStatus(status string) error {
    r.status = status
    r.Set("status", status)
    return nil
}

func (r YunosTvpubadminDeviceUpdateappstatusRequest) GetStatus() string {
    return r.status
}

func (r *YunosTvpubadminDeviceUpdateappstatusRequest) SetAuditComment(auditComment string) error {
    r.auditComment = auditComment
    r.Set("audit_comment", auditComment)
    return nil
}

func (r YunosTvpubadminDeviceUpdateappstatusRequest) GetAuditComment() string {
    return r.auditComment
}

