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
type YunosTvpubadminDeviceApkinfoRequest struct {
    model.Params
    // apkid
    _id   int64
}

// 初始化YunosTvpubadminDeviceApkinfoRequest对象
func NewYunosTvpubadminDeviceApkinfoRequest() *YunosTvpubadminDeviceApkinfoRequest{
    return &YunosTvpubadminDeviceApkinfoRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r YunosTvpubadminDeviceApkinfoRequest) GetApiMethodName() string {
    return "yunos.tvpubadmin.device.apkinfo"
}

// IRequest interface 方法, 获取API参数
func (r YunosTvpubadminDeviceApkinfoRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Id Setter
// apkid
func (r *YunosTvpubadminDeviceApkinfoRequest) SetId(_id int64) error {
    r._id = _id
    r.Set("id", _id)
    return nil
}

// Id Getter
func (r YunosTvpubadminDeviceApkinfoRequest) GetId() int64 {
    return r._id
}
