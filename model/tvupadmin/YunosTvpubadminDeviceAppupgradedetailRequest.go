package tvupadmin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取应用升级详情 API请求
yunos.tvpubadmin.device.appupgradedetail

获取应用升级详情
*/
type YunosTvpubadminDeviceAppupgradedetailRequest struct {
    model.Params
    // 应用升级的ID
    versionId   int64
    // 牌照方
    license   int64
}

// 初始化YunosTvpubadminDeviceAppupgradedetailRequest对象
func NewYunosTvpubadminDeviceAppupgradedetailRequest() *YunosTvpubadminDeviceAppupgradedetailRequest{
    return &YunosTvpubadminDeviceAppupgradedetailRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r YunosTvpubadminDeviceAppupgradedetailRequest) GetApiMethodName() string {
    return "yunos.tvpubadmin.device.appupgradedetail"
}

// IRequest interface 方法, 获取API参数
func (r YunosTvpubadminDeviceAppupgradedetailRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// VersionId Setter
// 应用升级的ID
func (r *YunosTvpubadminDeviceAppupgradedetailRequest) SetVersionId(versionId int64) error {
    r.versionId = versionId
    r.Set("version_id", versionId)
    return nil
}

// VersionId Getter
func (r YunosTvpubadminDeviceAppupgradedetailRequest) GetVersionId() int64 {
    return r.versionId
}
// License Setter
// 牌照方
func (r *YunosTvpubadminDeviceAppupgradedetailRequest) SetLicense(license int64) error {
    r.license = license
    r.Set("license", license)
    return nil
}

// License Getter
func (r YunosTvpubadminDeviceAppupgradedetailRequest) GetLicense() int64 {
    return r.license
}
