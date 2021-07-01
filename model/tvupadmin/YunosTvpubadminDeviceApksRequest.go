package tvupadmin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取停开服apk列表 API请求
yunos.tvpubadmin.device.apks

获取停开服apk列表
*/
type YunosTvpubadminDeviceApksAPIRequest struct {
    model.Params
    // 牌照
    _license   int64
}

// 初始化YunosTvpubadminDeviceApksAPIRequest对象
func NewYunosTvpubadminDeviceApksRequest() *YunosTvpubadminDeviceApksAPIRequest{
    return &YunosTvpubadminDeviceApksAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r YunosTvpubadminDeviceApksAPIRequest) GetApiMethodName() string {
    return "yunos.tvpubadmin.device.apks"
}

// IRequest interface 方法, 获取API参数
func (r YunosTvpubadminDeviceApksAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// License Setter
// 牌照
func (r *YunosTvpubadminDeviceApksAPIRequest) SetLicense(_license int64) error {
    r._license = _license
    r.Set("license", _license)
    return nil
}

// License Getter
func (r YunosTvpubadminDeviceApksAPIRequest) GetLicense() int64 {
    return r._license
}
