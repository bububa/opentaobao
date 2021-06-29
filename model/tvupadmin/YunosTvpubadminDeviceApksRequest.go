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
type YunosTvpubadminDeviceApksRequest struct {
    model.Params
    // 牌照
    _license   int64
}

// 初始化YunosTvpubadminDeviceApksRequest对象
func NewYunosTvpubadminDeviceApksRequest() *YunosTvpubadminDeviceApksRequest{
    return &YunosTvpubadminDeviceApksRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r YunosTvpubadminDeviceApksRequest) GetApiMethodName() string {
    return "yunos.tvpubadmin.device.apks"
}

// IRequest interface 方法, 获取API参数
func (r YunosTvpubadminDeviceApksRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// License Setter
// 牌照
func (r *YunosTvpubadminDeviceApksRequest) SetLicense(_license int64) error {
    r._license = _license
    r.Set("license", _license)
    return nil
}

// License Getter
func (r YunosTvpubadminDeviceApksRequest) GetLicense() int64 {
    return r._license
}
