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
    _deviceCode   string
    // 开始时间
    _playStartTime   string
    // 结束时间
    _playEndTime   string
    // 是否播放广告
    _isOnlyPlayAdvert   bool
    // 是否设置播放时间
    _isSetPlayTime   bool
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
func (r *AlibabaLstSpeakerConfigureSetpaytimeRequest) SetDeviceCode(_deviceCode string) error {
    r._deviceCode = _deviceCode
    r.Set("device_code", _deviceCode)
    return nil
}

// DeviceCode Getter
func (r AlibabaLstSpeakerConfigureSetpaytimeRequest) GetDeviceCode() string {
    return r._deviceCode
}
// PlayStartTime Setter
// 开始时间
func (r *AlibabaLstSpeakerConfigureSetpaytimeRequest) SetPlayStartTime(_playStartTime string) error {
    r._playStartTime = _playStartTime
    r.Set("play_start_time", _playStartTime)
    return nil
}

// PlayStartTime Getter
func (r AlibabaLstSpeakerConfigureSetpaytimeRequest) GetPlayStartTime() string {
    return r._playStartTime
}
// PlayEndTime Setter
// 结束时间
func (r *AlibabaLstSpeakerConfigureSetpaytimeRequest) SetPlayEndTime(_playEndTime string) error {
    r._playEndTime = _playEndTime
    r.Set("play_end_time", _playEndTime)
    return nil
}

// PlayEndTime Getter
func (r AlibabaLstSpeakerConfigureSetpaytimeRequest) GetPlayEndTime() string {
    return r._playEndTime
}
// IsOnlyPlayAdvert Setter
// 是否播放广告
func (r *AlibabaLstSpeakerConfigureSetpaytimeRequest) SetIsOnlyPlayAdvert(_isOnlyPlayAdvert bool) error {
    r._isOnlyPlayAdvert = _isOnlyPlayAdvert
    r.Set("is_only_play_advert", _isOnlyPlayAdvert)
    return nil
}

// IsOnlyPlayAdvert Getter
func (r AlibabaLstSpeakerConfigureSetpaytimeRequest) GetIsOnlyPlayAdvert() bool {
    return r._isOnlyPlayAdvert
}
// IsSetPlayTime Setter
// 是否设置播放时间
func (r *AlibabaLstSpeakerConfigureSetpaytimeRequest) SetIsSetPlayTime(_isSetPlayTime bool) error {
    r._isSetPlayTime = _isSetPlayTime
    r.Set("is_set_play_time", _isSetPlayTime)
    return nil
}

// IsSetPlayTime Getter
func (r AlibabaLstSpeakerConfigureSetpaytimeRequest) GetIsSetPlayTime() bool {
    return r._isSetPlayTime
}