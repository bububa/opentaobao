package lstspeacker

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
音箱播放配置 API请求
alibaba.lst.speaker.configure.setpaytime

音箱播放配置
*/
type AlibabaLstSpeakerConfigureSetpaytimeRequest struct {
    model.Params
    // 设备编码
    deviceCode   string
    // 开始时间
    playStartTime   string
    // 结束时间
    playEndTime   string
    // 是否播放广告
    isOnlyPlayAdvert   bool
    // 是否设置播放时间
    isSetPlayTime   bool
}

// 初始化AlibabaLstSpeakerConfigureSetpaytimeRequest对象
func NewAlibabaLstSpeakerConfigureSetpaytimeRequest() *AlibabaLstSpeakerConfigureSetpaytimeRequest{
    return &AlibabaLstSpeakerConfigureSetpaytimeRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaLstSpeakerConfigureSetpaytimeRequest) GetApiMethodName() string {
    return "alibaba.lst.speaker.configure.setpaytime"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaLstSpeakerConfigureSetpaytimeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// DeviceCode Setter
// 设备编码
func (r *AlibabaLstSpeakerConfigureSetpaytimeRequest) SetDeviceCode(deviceCode string) error {
    r.deviceCode = deviceCode
    r.Set("device_code", deviceCode)
    return nil
}

// DeviceCode Getter
func (r AlibabaLstSpeakerConfigureSetpaytimeRequest) GetDeviceCode() string {
    return r.deviceCode
}
// PlayStartTime Setter
// 开始时间
func (r *AlibabaLstSpeakerConfigureSetpaytimeRequest) SetPlayStartTime(playStartTime string) error {
    r.playStartTime = playStartTime
    r.Set("play_start_time", playStartTime)
    return nil
}

// PlayStartTime Getter
func (r AlibabaLstSpeakerConfigureSetpaytimeRequest) GetPlayStartTime() string {
    return r.playStartTime
}
// PlayEndTime Setter
// 结束时间
func (r *AlibabaLstSpeakerConfigureSetpaytimeRequest) SetPlayEndTime(playEndTime string) error {
    r.playEndTime = playEndTime
    r.Set("play_end_time", playEndTime)
    return nil
}

// PlayEndTime Getter
func (r AlibabaLstSpeakerConfigureSetpaytimeRequest) GetPlayEndTime() string {
    return r.playEndTime
}
// IsOnlyPlayAdvert Setter
// 是否播放广告
func (r *AlibabaLstSpeakerConfigureSetpaytimeRequest) SetIsOnlyPlayAdvert(isOnlyPlayAdvert bool) error {
    r.isOnlyPlayAdvert = isOnlyPlayAdvert
    r.Set("is_only_play_advert", isOnlyPlayAdvert)
    return nil
}

// IsOnlyPlayAdvert Getter
func (r AlibabaLstSpeakerConfigureSetpaytimeRequest) GetIsOnlyPlayAdvert() bool {
    return r.isOnlyPlayAdvert
}
// IsSetPlayTime Setter
// 是否设置播放时间
func (r *AlibabaLstSpeakerConfigureSetpaytimeRequest) SetIsSetPlayTime(isSetPlayTime bool) error {
    r.isSetPlayTime = isSetPlayTime
    r.Set("is_set_play_time", isSetPlayTime)
    return nil
}

// IsSetPlayTime Getter
func (r AlibabaLstSpeakerConfigureSetpaytimeRequest) GetIsSetPlayTime() bool {
    return r.isSetPlayTime
}
