package alihealthcrm

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
刷脸测睡眠数据同步 APIRequest
alibaba.alihealth.uic.swipeface.syncdata

刷脸测睡眠数据同步，三方数据回传
*/
type AlibabaAlihealthUicSwipefaceSyncdataRequest struct {
    model.Params

    // 用户ID
    tpUserId   int64 

    // 缺觉小时数
    lackSleepHours   int64 

    // 黑圆圈级别
    blackEyeLevel   int64 

    // 眼袋级别
    eyeBagSwollenLevel   int64 

    // 鱼尾纹数
    fishTail   int64 

    // 嘴唇颜色
    lipColor   string 

}

func NewAlibabaAlihealthUicSwipefaceSyncdataRequest() *AlibabaAlihealthUicSwipefaceSyncdataRequest{
    return &AlibabaAlihealthUicSwipefaceSyncdataRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihealthUicSwipefaceSyncdataRequest) GetApiMethodName() string {
    return "alibaba.alihealth.uic.swipeface.syncdata"
}

func (r AlibabaAlihealthUicSwipefaceSyncdataRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihealthUicSwipefaceSyncdataRequest) SetTpUserId(tpUserId int64) error {
    r.tpUserId = tpUserId
    r.Set("tp_user_id", tpUserId)
    return nil
}

func (r AlibabaAlihealthUicSwipefaceSyncdataRequest) GetTpUserId() int64 {
    return r.tpUserId
}

func (r *AlibabaAlihealthUicSwipefaceSyncdataRequest) SetLackSleepHours(lackSleepHours int64) error {
    r.lackSleepHours = lackSleepHours
    r.Set("lack_sleep_hours", lackSleepHours)
    return nil
}

func (r AlibabaAlihealthUicSwipefaceSyncdataRequest) GetLackSleepHours() int64 {
    return r.lackSleepHours
}

func (r *AlibabaAlihealthUicSwipefaceSyncdataRequest) SetBlackEyeLevel(blackEyeLevel int64) error {
    r.blackEyeLevel = blackEyeLevel
    r.Set("black_eye_level", blackEyeLevel)
    return nil
}

func (r AlibabaAlihealthUicSwipefaceSyncdataRequest) GetBlackEyeLevel() int64 {
    return r.blackEyeLevel
}

func (r *AlibabaAlihealthUicSwipefaceSyncdataRequest) SetEyeBagSwollenLevel(eyeBagSwollenLevel int64) error {
    r.eyeBagSwollenLevel = eyeBagSwollenLevel
    r.Set("eye_bag_swollen_level", eyeBagSwollenLevel)
    return nil
}

func (r AlibabaAlihealthUicSwipefaceSyncdataRequest) GetEyeBagSwollenLevel() int64 {
    return r.eyeBagSwollenLevel
}

func (r *AlibabaAlihealthUicSwipefaceSyncdataRequest) SetFishTail(fishTail int64) error {
    r.fishTail = fishTail
    r.Set("fish_tail", fishTail)
    return nil
}

func (r AlibabaAlihealthUicSwipefaceSyncdataRequest) GetFishTail() int64 {
    return r.fishTail
}

func (r *AlibabaAlihealthUicSwipefaceSyncdataRequest) SetLipColor(lipColor string) error {
    r.lipColor = lipColor
    r.Set("lip_color", lipColor)
    return nil
}

func (r AlibabaAlihealthUicSwipefaceSyncdataRequest) GetLipColor() string {
    return r.lipColor
}

