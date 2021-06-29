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
    _token   string
    // 播控方，1:华数 7:CIBN
    _bcp   string
    // 终端设备型号
    _deviceModel   string
    // 桌面版本号
    _versionCode   int64
    // 终端设备mac
    _mac   string
    // 终端设备uuid
    _uuid   string
    // 影视来源，0-淘tv 7-优酷免费 9-优酷付费 多项用“,”隔开
    _from   string
    // 支持收费方式，0-免费 1-限免 2-单点 3-包月 4-红包 5-单包,多项用“,”隔开
    _chargeType   string
    // 获取的最大节目数
    _maxSize   int64
    // 分辨率，sw720 sw1080
    _sw   string
    // 支持媒体类型,dts,drm,dolby,h265
    _deviceMedia   string
    // 请求IP
    _ip   string
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
func (r *YoukuTvDesktopToyouRecommendRequest) SetToken(_token string) error {
    r._token = _token
    r.Set("token", _token)
    return nil
}

// Token Getter
func (r YoukuTvDesktopToyouRecommendRequest) GetToken() string {
    return r._token
}
// Bcp Setter
// 播控方，1:华数 7:CIBN
func (r *YoukuTvDesktopToyouRecommendRequest) SetBcp(_bcp string) error {
    r._bcp = _bcp
    r.Set("bcp", _bcp)
    return nil
}

// Bcp Getter
func (r YoukuTvDesktopToyouRecommendRequest) GetBcp() string {
    return r._bcp
}
// DeviceModel Setter
// 终端设备型号
func (r *YoukuTvDesktopToyouRecommendRequest) SetDeviceModel(_deviceModel string) error {
    r._deviceModel = _deviceModel
    r.Set("device_model", _deviceModel)
    return nil
}

// DeviceModel Getter
func (r YoukuTvDesktopToyouRecommendRequest) GetDeviceModel() string {
    return r._deviceModel
}
// VersionCode Setter
// 桌面版本号
func (r *YoukuTvDesktopToyouRecommendRequest) SetVersionCode(_versionCode int64) error {
    r._versionCode = _versionCode
    r.Set("version_code", _versionCode)
    return nil
}

// VersionCode Getter
func (r YoukuTvDesktopToyouRecommendRequest) GetVersionCode() int64 {
    return r._versionCode
}
// Mac Setter
// 终端设备mac
func (r *YoukuTvDesktopToyouRecommendRequest) SetMac(_mac string) error {
    r._mac = _mac
    r.Set("mac", _mac)
    return nil
}

// Mac Getter
func (r YoukuTvDesktopToyouRecommendRequest) GetMac() string {
    return r._mac
}
// Uuid Setter
// 终端设备uuid
func (r *YoukuTvDesktopToyouRecommendRequest) SetUuid(_uuid string) error {
    r._uuid = _uuid
    r.Set("uuid", _uuid)
    return nil
}

// Uuid Getter
func (r YoukuTvDesktopToyouRecommendRequest) GetUuid() string {
    return r._uuid
}
// From Setter
// 影视来源，0-淘tv 7-优酷免费 9-优酷付费 多项用“,”隔开
func (r *YoukuTvDesktopToyouRecommendRequest) SetFrom(_from string) error {
    r._from = _from
    r.Set("from", _from)
    return nil
}

// From Getter
func (r YoukuTvDesktopToyouRecommendRequest) GetFrom() string {
    return r._from
}
// ChargeType Setter
// 支持收费方式，0-免费 1-限免 2-单点 3-包月 4-红包 5-单包,多项用“,”隔开
func (r *YoukuTvDesktopToyouRecommendRequest) SetChargeType(_chargeType string) error {
    r._chargeType = _chargeType
    r.Set("charge_type", _chargeType)
    return nil
}

// ChargeType Getter
func (r YoukuTvDesktopToyouRecommendRequest) GetChargeType() string {
    return r._chargeType
}
// MaxSize Setter
// 获取的最大节目数
func (r *YoukuTvDesktopToyouRecommendRequest) SetMaxSize(_maxSize int64) error {
    r._maxSize = _maxSize
    r.Set("max_size", _maxSize)
    return nil
}

// MaxSize Getter
func (r YoukuTvDesktopToyouRecommendRequest) GetMaxSize() int64 {
    return r._maxSize
}
// Sw Setter
// 分辨率，sw720 sw1080
func (r *YoukuTvDesktopToyouRecommendRequest) SetSw(_sw string) error {
    r._sw = _sw
    r.Set("sw", _sw)
    return nil
}

// Sw Getter
func (r YoukuTvDesktopToyouRecommendRequest) GetSw() string {
    return r._sw
}
// DeviceMedia Setter
// 支持媒体类型,dts,drm,dolby,h265
func (r *YoukuTvDesktopToyouRecommendRequest) SetDeviceMedia(_deviceMedia string) error {
    r._deviceMedia = _deviceMedia
    r.Set("device_media", _deviceMedia)
    return nil
}

// DeviceMedia Getter
func (r YoukuTvDesktopToyouRecommendRequest) GetDeviceMedia() string {
    return r._deviceMedia
}
// Ip Setter
// 请求IP
func (r *YoukuTvDesktopToyouRecommendRequest) SetIp(_ip string) error {
    r._ip = _ip
    r.Set("ip", _ip)
    return nil
}

// Ip Getter
func (r YoukuTvDesktopToyouRecommendRequest) GetIp() string {
    return r._ip
}
