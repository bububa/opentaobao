package tvupadmin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取停开服apk信息 APIRequest
yunos.tvpubadmin.device.apkinfo

获取停开服apk信息
*/
type YunosTvpubadminDeviceApkinfoRequest struct {
    model.Params

    // apkid
    id   int64 

}

func NewYunosTvpubadminDeviceApkinfoRequest() *YunosTvpubadminDeviceApkinfoRequest{
    return &YunosTvpubadminDeviceApkinfoRequest{
        Params: model.NewParams(),
    }
}

func (r YunosTvpubadminDeviceApkinfoRequest) GetApiMethodName() string {
    return "yunos.tvpubadmin.device.apkinfo"
}

func (r YunosTvpubadminDeviceApkinfoRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *YunosTvpubadminDeviceApkinfoRequest) SetId(id int64) error {
    r.id = id
    r.Set("id", id)
    return nil
}

func (r YunosTvpubadminDeviceApkinfoRequest) GetId() int64 {
    return r.id
}

