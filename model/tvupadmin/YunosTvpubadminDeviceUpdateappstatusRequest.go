package tvupadmin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
更新应用版本审核状态 API请求
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

// 初始化YunosTvpubadminDeviceUpdateappstatusRequest对象
func NewYunosTvpubadminDeviceUpdateappstatusRequest() *YunosTvpubadminDeviceUpdateappstatusRequest{
    return &YunosTvpubadminDeviceUpdateappstatusRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r YunosTvpubadminDeviceUpdateappstatusRequest) GetApiMethodName() string {
    return "yunos.tvpubadmin.device.updateappstatus"
}

// IRequest interface 方法, 获取API参数
func (r YunosTvpubadminDeviceUpdateappstatusRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// VersionId Setter
// 应用ID
func (r *YunosTvpubadminDeviceUpdateappstatusRequest) SetVersionId(versionId int64) error {
    r.versionId = versionId
    r.Set("version_id", versionId)
    return nil
}

// VersionId Getter
func (r YunosTvpubadminDeviceUpdateappstatusRequest) GetVersionId() int64 {
    return r.versionId
}
// License Setter
// 牌照方
func (r *YunosTvpubadminDeviceUpdateappstatusRequest) SetLicense(license int64) error {
    r.license = license
    r.Set("license", license)
    return nil
}

// License Getter
func (r YunosTvpubadminDeviceUpdateappstatusRequest) GetLicense() int64 {
    return r.license
}
// Status Setter
// 审核状态
func (r *YunosTvpubadminDeviceUpdateappstatusRequest) SetStatus(status string) error {
    r.status = status
    r.Set("status", status)
    return nil
}

// Status Getter
func (r YunosTvpubadminDeviceUpdateappstatusRequest) GetStatus() string {
    return r.status
}
// AuditComment Setter
// 审核意见
func (r *YunosTvpubadminDeviceUpdateappstatusRequest) SetAuditComment(auditComment string) error {
    r.auditComment = auditComment
    r.Set("audit_comment", auditComment)
    return nil
}

// AuditComment Getter
func (r YunosTvpubadminDeviceUpdateappstatusRequest) GetAuditComment() string {
    return r.auditComment
}
