package tvupadmin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取停开服apk列表 APIRequest
yunos.tvpubadmin.device.apks

获取停开服apk列表
*/
type YunosTvpubadminDeviceApksRequest struct {
    model.Params

    // 牌照
    license   int64 

}

func NewYunosTvpubadminDeviceApksRequest() *YunosTvpubadminDeviceApksRequest{
    return &YunosTvpubadminDeviceApksRequest{
        Params: model.NewParams(),
    }
}

func (r YunosTvpubadminDeviceApksRequest) GetApiMethodName() string {
    return "yunos.tvpubadmin.device.apks"
}

func (r YunosTvpubadminDeviceApksRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *YunosTvpubadminDeviceApksRequest) SetLicense(license int64) error {
    r.license = license
    r.Set("license", license)
    return nil
}

func (r YunosTvpubadminDeviceApksRequest) GetLicense() int64 {
    return r.license
}

