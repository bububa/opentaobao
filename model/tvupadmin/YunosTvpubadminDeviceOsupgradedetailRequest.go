package tvupadmin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取系统升级详情 API请求
yunos.tvpubadmin.device.osupgradedetail

获取系统升级详情
*/
type YunosTvpubadminDeviceOsupgradedetailRequest struct {
    model.Params
    // 版本ID
    versionId   int64
    // 牌照方
    license   int64
}

// 初始化YunosTvpubadminDeviceOsupgradedetailRequest对象
func NewYunosTvpubadminDeviceOsupgradedetailRequest() *YunosTvpubadminDeviceOsupgradedetailRequest{
    return &YunosTvpubadminDeviceOsupgradedetailRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r YunosTvpubadminDeviceOsupgradedetailRequest) GetApiMethodName() string {
    return "yunos.tvpubadmin.device.osupgradedetail"
}

// IRequest interface 方法, 获取API参数
func (r YunosTvpubadminDeviceOsupgradedetailRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// VersionId Setter
// 版本ID
func (r *YunosTvpubadminDeviceOsupgradedetailRequest) SetVersionId(versionId int64) error {
    r.versionId = versionId
    r.Set("version_id", versionId)
    return nil
}

// VersionId Getter
func (r YunosTvpubadminDeviceOsupgradedetailRequest) GetVersionId() int64 {
    return r.versionId
}
// License Setter
// 牌照方
func (r *YunosTvpubadminDeviceOsupgradedetailRequest) SetLicense(license int64) error {
    r.license = license
    r.Set("license", license)
    return nil
}

// License Getter
func (r YunosTvpubadminDeviceOsupgradedetailRequest) GetLicense() int64 {
    return r.license
}
