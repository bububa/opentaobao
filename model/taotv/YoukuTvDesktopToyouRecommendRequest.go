package taotv

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
TV桌面为你推荐接口 API请求
youku.tv.desktop.toyou.recommend

提供为你推荐数据
*/
type YoukuTvDesktopToyouRecommendRequest struct {
    model.Params
    // 用户登陆token
    token   string
    // 播控方，1:华数 7:CIBN
    bcp   string
    // 终端设备型号
    deviceModel   string
    // 桌面版本号
    versionCode   int64
    // 终端设备mac
    mac   string
    // 终端设备uuid
    uuid   string
    // 影视来源，0-淘tv 7-优酷免费 9-优酷付费 多项用“,”隔开
    from   string
    // 支持收费方式，0-免费 1-限免 2-单点 3-包月 4-红包 5-单包,多项用“,”隔开
    chargeType   string
    // 获取的最大节目数
    maxSize   int64
    // 分辨率，sw720 sw1080
    sw   string
    // 支持媒体类型,dts,drm,dolby,h265
    deviceMedia   string
    // 请求IP
    ip   string
}

// 初始化YoukuTvDesktopToyouRecommendRequest对象
func NewYoukuTvDesktopToyouRecommendRequest() *YoukuTvDesktopToyouRecommendRequest{
    return &YoukuTvDesktopToyouRecommendRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r YoukuTvDesktopToyouRecommendRequest) GetApiMethodName() string {
    return "youku.tv.desktop.toyou.recommend"
}

// IRequest interface 方法, 获取API参数
func (r YoukuTvDesktopToyouRecommendRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Token Setter
// 用户登陆token
func (r *YoukuTvDesktopToyouRecommendRequest) SetToken(token string) error {
    r.token = token
    r.Set("token", token)
    return nil
}

// Token Getter
func (r YoukuTvDesktopToyouRecommendRequest) GetToken() string {
    return r.token
}
// Bcp Setter
// 播控方，1:华数 7:CIBN
func (r *YoukuTvDesktopToyouRecommendRequest) SetBcp(bcp string) error {
    r.bcp = bcp
    r.Set("bcp", bcp)
    return nil
}

// Bcp Getter
func (r YoukuTvDesktopToyouRecommendRequest) GetBcp() string {
    return r.bcp
}
// DeviceModel Setter
// 终端设备型号
func (r *YoukuTvDesktopToyouRecommendRequest) SetDeviceModel(deviceModel string) error {
    r.deviceModel = deviceModel
    r.Set("device_model", deviceModel)
    return nil
}

// DeviceModel Getter
func (r YoukuTvDesktopToyouRecommendRequest) GetDeviceModel() string {
    return r.deviceModel
}
// VersionCode Setter
// 桌面版本号
func (r *YoukuTvDesktopToyouRecommendRequest) SetVersionCode(versionCode int64) error {
    r.versionCode = versionCode
    r.Set("version_code", versionCode)
    return nil
}

// VersionCode Getter
func (r YoukuTvDesktopToyouRecommendRequest) GetVersionCode() int64 {
    return r.versionCode
}
// Mac Setter
// 终端设备mac
func (r *YoukuTvDesktopToyouRecommendRequest) SetMac(mac string) error {
    r.mac = mac
    r.Set("mac", mac)
    return nil
}

// Mac Getter
func (r YoukuTvDesktopToyouRecommendRequest) GetMac() string {
    return r.mac
}
// Uuid Setter
// 终端设备uuid
func (r *YoukuTvDesktopToyouRecommendRequest) SetUuid(uuid string) error {
    r.uuid = uuid
    r.Set("uuid", uuid)
    return nil
}

// Uuid Getter
func (r YoukuTvDesktopToyouRecommendRequest) GetUuid() string {
    return r.uuid
}
// From Setter
// 影视来源，0-淘tv 7-优酷免费 9-优酷付费 多项用“,”隔开
func (r *YoukuTvDesktopToyouRecommendRequest) SetFrom(from string) error {
    r.from = from
    r.Set("from", from)
    return nil
}

// From Getter
func (r YoukuTvDesktopToyouRecommendRequest) GetFrom() string {
    return r.from
}
// ChargeType Setter
// 支持收费方式，0-免费 1-限免 2-单点 3-包月 4-红包 5-单包,多项用“,”隔开
func (r *YoukuTvDesktopToyouRecommendRequest) SetChargeType(chargeType string) error {
    r.chargeType = chargeType
    r.Set("charge_type", chargeType)
    return nil
}

// ChargeType Getter
func (r YoukuTvDesktopToyouRecommendRequest) GetChargeType() string {
    return r.chargeType
}
// MaxSize Setter
// 获取的最大节目数
func (r *YoukuTvDesktopToyouRecommendRequest) SetMaxSize(maxSize int64) error {
    r.maxSize = maxSize
    r.Set("max_size", maxSize)
    return nil
}

// MaxSize Getter
func (r YoukuTvDesktopToyouRecommendRequest) GetMaxSize() int64 {
    return r.maxSize
}
// Sw Setter
// 分辨率，sw720 sw1080
func (r *YoukuTvDesktopToyouRecommendRequest) SetSw(sw string) error {
    r.sw = sw
    r.Set("sw", sw)
    return nil
}

// Sw Getter
func (r YoukuTvDesktopToyouRecommendRequest) GetSw() string {
    return r.sw
}
// DeviceMedia Setter
// 支持媒体类型,dts,drm,dolby,h265
func (r *YoukuTvDesktopToyouRecommendRequest) SetDeviceMedia(deviceMedia string) error {
    r.deviceMedia = deviceMedia
    r.Set("device_media", deviceMedia)
    return nil
}

// DeviceMedia Getter
func (r YoukuTvDesktopToyouRecommendRequest) GetDeviceMedia() string {
    return r.deviceMedia
}
// Ip Setter
// 请求IP
func (r *YoukuTvDesktopToyouRecommendRequest) SetIp(ip string) error {
    r.ip = ip
    r.Set("ip", ip)
    return nil
}

// Ip Getter
func (r YoukuTvDesktopToyouRecommendRequest) GetIp() string {
    return r.ip
}
