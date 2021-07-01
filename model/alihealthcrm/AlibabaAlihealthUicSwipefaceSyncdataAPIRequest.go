package alihealthcrm

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
刷脸测睡眠数据同步 API请求
alibaba.alihealth.uic.swipeface.syncdata

刷脸测睡眠数据同步，三方数据回传
*/
type AlibabaAlihealthUicSwipefaceSyncdataAPIRequest struct {
    model.Params
    // 用户ID
    _tpUserId   int64
    // 缺觉小时数
    _lackSleepHours   int64
    // 黑圆圈级别
    _blackEyeLevel   int64
    // 眼袋级别
    _eyeBagSwollenLevel   int64
    // 鱼尾纹数
    _fishTail   int64
    // 嘴唇颜色
    _lipColor   string
}

// 初始化AlibabaAlihealthUicSwipefaceSyncdataAPIRequest对象
func NewAlibabaAlihealthUicSwipefaceSyncdataRequest() *AlibabaAlihealthUicSwipefaceSyncdataAPIRequest{
    return &AlibabaAlihealthUicSwipefaceSyncdataAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthUicSwipefaceSyncdataAPIRequest) GetApiMethodName() string {
    return "alibaba.alihealth.uic.swipeface.syncdata"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthUicSwipefaceSyncdataAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// TpUserId Setter
// 用户ID
func (r *AlibabaAlihealthUicSwipefaceSyncdataAPIRequest) SetTpUserId(_tpUserId int64) error {
    r._tpUserId = _tpUserId
    r.Set("tp_user_id", _tpUserId)
    return nil
}

// TpUserId Getter
func (r AlibabaAlihealthUicSwipefaceSyncdataAPIRequest) GetTpUserId() int64 {
    return r._tpUserId
}
// LackSleepHours Setter
// 缺觉小时数
func (r *AlibabaAlihealthUicSwipefaceSyncdataAPIRequest) SetLackSleepHours(_lackSleepHours int64) error {
    r._lackSleepHours = _lackSleepHours
    r.Set("lack_sleep_hours", _lackSleepHours)
    return nil
}

// LackSleepHours Getter
func (r AlibabaAlihealthUicSwipefaceSyncdataAPIRequest) GetLackSleepHours() int64 {
    return r._lackSleepHours
}
// BlackEyeLevel Setter
// 黑圆圈级别
func (r *AlibabaAlihealthUicSwipefaceSyncdataAPIRequest) SetBlackEyeLevel(_blackEyeLevel int64) error {
    r._blackEyeLevel = _blackEyeLevel
    r.Set("black_eye_level", _blackEyeLevel)
    return nil
}

// BlackEyeLevel Getter
func (r AlibabaAlihealthUicSwipefaceSyncdataAPIRequest) GetBlackEyeLevel() int64 {
    return r._blackEyeLevel
}
// EyeBagSwollenLevel Setter
// 眼袋级别
func (r *AlibabaAlihealthUicSwipefaceSyncdataAPIRequest) SetEyeBagSwollenLevel(_eyeBagSwollenLevel int64) error {
    r._eyeBagSwollenLevel = _eyeBagSwollenLevel
    r.Set("eye_bag_swollen_level", _eyeBagSwollenLevel)
    return nil
}

// EyeBagSwollenLevel Getter
func (r AlibabaAlihealthUicSwipefaceSyncdataAPIRequest) GetEyeBagSwollenLevel() int64 {
    return r._eyeBagSwollenLevel
}
// FishTail Setter
// 鱼尾纹数
func (r *AlibabaAlihealthUicSwipefaceSyncdataAPIRequest) SetFishTail(_fishTail int64) error {
    r._fishTail = _fishTail
    r.Set("fish_tail", _fishTail)
    return nil
}

// FishTail Getter
func (r AlibabaAlihealthUicSwipefaceSyncdataAPIRequest) GetFishTail() int64 {
    return r._fishTail
}
// LipColor Setter
// 嘴唇颜色
func (r *AlibabaAlihealthUicSwipefaceSyncdataAPIRequest) SetLipColor(_lipColor string) error {
    r._lipColor = _lipColor
    r.Set("lip_color", _lipColor)
    return nil
}

// LipColor Getter
func (r AlibabaAlihealthUicSwipefaceSyncdataAPIRequest) GetLipColor() string {
    return r._lipColor
}
