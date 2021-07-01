package tvupadmin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取停开服apk信息 API请求
yunos.tvpubadmin.device.apkinfo

获取停开服apk信息
*/
type YunosTvpubadminDeviceApkinfoAPIRequest struct {
    model.Params
    // apkid
    _id   int64
}

// 初始化YunosTvpubadminDeviceApkinfoAPIRequest对象
func NewYunosTvpubadminDeviceApkinfoRequest() *YunosTvpubadminDeviceApkinfoAPIRequest{
    return &YunosTvpubadminDeviceApkinfoAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r YunosTvpubadminDeviceApkinfoAPIRequest) GetApiMethodName() string {
    return "yunos.tvpubadmin.device.apkinfo"
}

// IRequest interface 方法, 获取API参数
func (r YunosTvpubadminDeviceApkinfoAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Id Setter
// apkid
func (r *YunosTvpubadminDeviceApkinfoAPIRequest) SetId(_id int64) error {
    r._id = _id
    r.Set("id", _id)
    return nil
}

// Id Getter
func (r YunosTvpubadminDeviceApkinfoAPIRequest) GetId() int64 {
    return r._id
}
