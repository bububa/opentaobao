package taotv

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
TV桌面为你推荐接口 APIRequest
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

func NewYoukuTvDesktopToyouRecommendRequest() *YoukuTvDesktopToyouRecommendRequest{
    return &YoukuTvDesktopToyouRecommendRequest{
        Params: model.NewParams(),
    }
}

func (r YoukuTvDesktopToyouRecommendRequest) GetApiMethodName() string {
    return "youku.tv.desktop.toyou.recommend"
}

func (r YoukuTvDesktopToyouRecommendRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *YoukuTvDesktopToyouRecommendRequest) SetToken(token string) error {
    r.token = token
    r.Set("token", token)
    return nil
}

func (r YoukuTvDesktopToyouRecommendRequest) GetToken() string {
    return r.token
}

func (r *YoukuTvDesktopToyouRecommendRequest) SetBcp(bcp string) error {
    r.bcp = bcp
    r.Set("bcp", bcp)
    return nil
}

func (r YoukuTvDesktopToyouRecommendRequest) GetBcp() string {
    return r.bcp
}

func (r *YoukuTvDesktopToyouRecommendRequest) SetDeviceModel(deviceModel string) error {
    r.deviceModel = deviceModel
    r.Set("device_model", deviceModel)
    return nil
}

func (r YoukuTvDesktopToyouRecommendRequest) GetDeviceModel() string {
    return r.deviceModel
}

func (r *YoukuTvDesktopToyouRecommendRequest) SetVersionCode(versionCode int64) error {
    r.versionCode = versionCode
    r.Set("version_code", versionCode)
    return nil
}

func (r YoukuTvDesktopToyouRecommendRequest) GetVersionCode() int64 {
    return r.versionCode
}

func (r *YoukuTvDesktopToyouRecommendRequest) SetMac(mac string) error {
    r.mac = mac
    r.Set("mac", mac)
    return nil
}

func (r YoukuTvDesktopToyouRecommendRequest) GetMac() string {
    return r.mac
}

func (r *YoukuTvDesktopToyouRecommendRequest) SetUuid(uuid string) error {
    r.uuid = uuid
    r.Set("uuid", uuid)
    return nil
}

func (r YoukuTvDesktopToyouRecommendRequest) GetUuid() string {
    return r.uuid
}

func (r *YoukuTvDesktopToyouRecommendRequest) SetFrom(from string) error {
    r.from = from
    r.Set("from", from)
    return nil
}

func (r YoukuTvDesktopToyouRecommendRequest) GetFrom() string {
    return r.from
}

func (r *YoukuTvDesktopToyouRecommendRequest) SetChargeType(chargeType string) error {
    r.chargeType = chargeType
    r.Set("charge_type", chargeType)
    return nil
}

func (r YoukuTvDesktopToyouRecommendRequest) GetChargeType() string {
    return r.chargeType
}

func (r *YoukuTvDesktopToyouRecommendRequest) SetMaxSize(maxSize int64) error {
    r.maxSize = maxSize
    r.Set("max_size", maxSize)
    return nil
}

func (r YoukuTvDesktopToyouRecommendRequest) GetMaxSize() int64 {
    return r.maxSize
}

func (r *YoukuTvDesktopToyouRecommendRequest) SetSw(sw string) error {
    r.sw = sw
    r.Set("sw", sw)
    return nil
}

func (r YoukuTvDesktopToyouRecommendRequest) GetSw() string {
    return r.sw
}

func (r *YoukuTvDesktopToyouRecommendRequest) SetDeviceMedia(deviceMedia string) error {
    r.deviceMedia = deviceMedia
    r.Set("device_media", deviceMedia)
    return nil
}

func (r YoukuTvDesktopToyouRecommendRequest) GetDeviceMedia() string {
    return r.deviceMedia
}

func (r *YoukuTvDesktopToyouRecommendRequest) SetIp(ip string) error {
    r.ip = ip
    r.Set("ip", ip)
    return nil
}

func (r YoukuTvDesktopToyouRecommendRequest) GetIp() string {
    return r.ip
}

