package tvupadmin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取应用升级详情 APIRequest
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

func NewYunosTvpubadminDeviceAppupgradedetailRequest() *YunosTvpubadminDeviceAppupgradedetailRequest{
    return &YunosTvpubadminDeviceAppupgradedetailRequest{
        Params: model.NewParams(),
    }
}

func (r YunosTvpubadminDeviceAppupgradedetailRequest) GetApiMethodName() string {
    return "yunos.tvpubadmin.device.appupgradedetail"
}

func (r YunosTvpubadminDeviceAppupgradedetailRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *YunosTvpubadminDeviceAppupgradedetailRequest) SetVersionId(versionId int64) error {
    r.versionId = versionId
    r.Set("version_id", versionId)
    return nil
}

func (r YunosTvpubadminDeviceAppupgradedetailRequest) GetVersionId() int64 {
    return r.versionId
}

func (r *YunosTvpubadminDeviceAppupgradedetailRequest) SetLicense(license int64) error {
    r.license = license
    r.Set("license", license)
    return nil
}

func (r YunosTvpubadminDeviceAppupgradedetailRequest) GetLicense() int64 {
    return r.license
}

