package tvupadmin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
更新系统版本审核状态 API请求
yunos.tvpubadmin.device.updateosstatus

更新系统版本审核状态
*/
type YunosTvpubadminDeviceUpdateosstatusRequest struct {
    model.Params
    // 升级ID
    _versionId   int64
    // 牌照方
    _license   int64
    // 审核状态
    _status   string
    // 审核意见
    _auditComment   string
}

// 初始化YunosTvpubadminDeviceUpdateosstatusRequest对象
func NewYunosTvpubadminDeviceUpdateosstatusRequest() *YunosTvpubadminDeviceUpdateosstatusRequest{
    return &YunosTvpubadminDeviceUpdateosstatusRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r YunosTvpubadminDeviceUpdateosstatusRequest) GetApiMethodName() string {
    return "yunos.tvpubadmin.device.updateosstatus"
}

// IRequest interface 方法, 获取API参数
func (r YunosTvpubadminDeviceUpdateosstatusRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// VersionId Setter
// 升级ID
func (r *YunosTvpubadminDeviceUpdateosstatusRequest) SetVersionId(_versionId int64) error {
    r._versionId = _versionId
    r.Set("version_id", _versionId)
    return nil
}

// VersionId Getter
func (r YunosTvpubadminDeviceUpdateosstatusRequest) GetVersionId() int64 {
    return r._versionId
}
// License Setter
// 牌照方
func (r *YunosTvpubadminDeviceUpdateosstatusRequest) SetLicense(_license int64) error {
    r._license = _license
    r.Set("license", _license)
    return nil
}

// License Getter
func (r YunosTvpubadminDeviceUpdateosstatusRequest) GetLicense() int64 {
    return r._license
}
// Status Setter
// 审核状态
func (r *YunosTvpubadminDeviceUpdateosstatusRequest) SetStatus(_status string) error {
    r._status = _status
    r.Set("status", _status)
    return nil
}

// Status Getter
func (r YunosTvpubadminDeviceUpdateosstatusRequest) GetStatus() string {
    return r._status
}
// AuditComment Setter
// 审核意见
func (r *YunosTvpubadminDeviceUpdateosstatusRequest) SetAuditComment(_auditComment string) error {
    r._auditComment = _auditComment
    r.Set("audit_comment", _auditComment)
    return nil
}

// AuditComment Getter
func (r YunosTvpubadminDeviceUpdateosstatusRequest) GetAuditComment() string {
    return r._auditComment
}
