package baodian

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取宝点SDK的配置项（已迁移） API请求
taobao.baodian.server.sdk.config.get

获取SDK各种配置项（已迁移）
*/
type TaobaoBaodianServerSdkConfigGetRequest struct {
    model.Params
    // appKey
    appkey   string
    // 渠道
    channel   string
    // sdk版本号
    sdkVer   string
    // 与后端约定
    type   int64
}

// 初始化TaobaoBaodianServerSdkConfigGetRequest对象
func NewTaobaoBaodianServerSdkConfigGetRequest() *TaobaoBaodianServerSdkConfigGetRequest{
    return &TaobaoBaodianServerSdkConfigGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoBaodianServerSdkConfigGetRequest) GetApiMethodName() string {
    return "taobao.baodian.server.sdk.config.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoBaodianServerSdkConfigGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Appkey Setter
// appKey
func (r *TaobaoBaodianServerSdkConfigGetRequest) SetAppkey(appkey string) error {
    r.appkey = appkey
    r.Set("appkey", appkey)
    return nil
}

// Appkey Getter
func (r TaobaoBaodianServerSdkConfigGetRequest) GetAppkey() string {
    return r.appkey
}
// Channel Setter
// 渠道
func (r *TaobaoBaodianServerSdkConfigGetRequest) SetChannel(channel string) error {
    r.channel = channel
    r.Set("channel", channel)
    return nil
}

// Channel Getter
func (r TaobaoBaodianServerSdkConfigGetRequest) GetChannel() string {
    return r.channel
}
// SdkVer Setter
// sdk版本号
func (r *TaobaoBaodianServerSdkConfigGetRequest) SetSdkVer(sdkVer string) error {
    r.sdkVer = sdkVer
    r.Set("sdk_ver", sdkVer)
    return nil
}

// SdkVer Getter
func (r TaobaoBaodianServerSdkConfigGetRequest) GetSdkVer() string {
    return r.sdkVer
}
// Type Setter
// 与后端约定
func (r *TaobaoBaodianServerSdkConfigGetRequest) SetType(type int64) error {
    r.type = type
    r.Set("type", type)
    return nil
}

// Type Getter
func (r TaobaoBaodianServerSdkConfigGetRequest) GetType() int64 {
    return r.type
}
