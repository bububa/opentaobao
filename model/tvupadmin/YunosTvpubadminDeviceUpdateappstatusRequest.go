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
    _versionId   int64
    // 牌照方
    _license   int64
    // 审核状态
    _status   string
    // 审核意见
    _auditComment   string
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
func (r *YunosTvpubadminDeviceUpdateappstatusRequest) SetVersionId(_versionId int64) error {
    r._versionId = _versionId
    r.Set("version_id", _versionId)
    return nil
}

// VersionId Getter
func (r YunosTvpubadminDeviceUpdateappstatusRequest) GetVersionId() int64 {
    return r._versionId
}
// License Setter
// 牌照方
func (r *YunosTvpubadminDeviceUpdateappstatusRequest) SetLicense(_license int64) error {
    r._license = _license
    r.Set("license", _license)
    return nil
}

// License Getter
func (r YunosTvpubadminDeviceUpdateappstatusRequest) GetLicense() int64 {
    return r._license
}
// Status Setter
// 审核状态
func (r *YunosTvpubadminDeviceUpdateappstatusRequest) SetStatus(_status string) error {
    r._status = _status
    r.Set("status", _status)
    return nil
}

// Status Getter
func (r YunosTvpubadminDeviceUpdateappstatusRequest) GetStatus() string {
    return r._status
}
// AuditComment Setter
// 审核意见
func (r *YunosTvpubadminDeviceUpdateappstatusRequest) SetAuditComment(_auditComment string) error {
    r._auditComment = _auditComment
    r.Set("audit_comment", _auditComment)
    return nil
}

// AuditComment Getter
func (r YunosTvpubadminDeviceUpdateappstatusRequest) GetAuditComment() string {
    return r._auditComment
}
