package aliyun

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
推送通知 APIRequest
push.aliyuncs.com.pushNotification.2015-03-18

pushNotification
*/
type PushAliyuncsComPushNotification2015-03-18Request struct {
    model.Params

    // 用户账号列表,以换行区分,仅sendType为3时有效
    account   string 

    // 自定义的kv结构,开发者扩展用
    androidExtraMap   string 

    // 通知声音
    androidMusic   string 

    // 通知类型 1:震动 2:响铃
    androidNotifyType   int64 

    // 打开app指定位置
    androidOpenActivity   string 

    // 点击通知后动作
    androidOpenType   int64 

    // 打开应用,网页
    androidOpenUrl   string 

    // 防打扰时长,取值范围为1~23
    antiHarassDuration   int64 

    // 防打扰开始时间点,取值范围为0~23
    antiHarassStartTime   int64 

    // 应用标识
    appId   int64 

    // 批次编号,用于活动效果统计
    batchNumber   string 

    // 设备编号列表,以换行区分,仅sendType为4时有效
    deviceId   string 

    // 设备类型,取值范围为:0~3云推送支持多种设备,各 种设备类型编号如下:IOS设备:deviceType&amp;1=1; Andriod设备:deviceType&amp;2=2;如果存在此字段,则向 指定的设备类型推送消息。默认为全部(3);
    deviceType   int64 

    // 自定义的kv结构,开发者扩展用
    iosExtraMap   string 

    // 角标
    iosFooter   int64 

    // 默认音乐
    iosMusic   string 

    // 推送时间,若空表示立即推送,推送时间不能早于当前时间
    pushTime   string 

    // 推送类型,取值范围:1~4; 1:所有人,无需指定tag、 deviceType等2:一群人,必须指定tag3:指定用户,根据 用户账号列表文件发送消息4:指定设备,根据设备编码列 表文件发送消息默认值为1
    sendType   int64 

    // 摘要
    summary   string 

    // 标签名称,仅支持1个标签,仅sendType为2时有效
    tag   string 

    // 离线消息保存时长,取值范围为1~72,若不填,则表示不保存离线消息
    timeout   int64 

    // 标题
    title   string 

}

func NewPushAliyuncsComPushNotification2015-03-18Request() *PushAliyuncsComPushNotification2015-03-18Request{
    return &PushAliyuncsComPushNotification2015-03-18Request{
        Params: model.NewParams(),
    }
}

func (r PushAliyuncsComPushNotification2015-03-18Request) GetApiMethodName() string {
    return "push.aliyuncs.com.pushNotification.2015-03-18"
}

func (r PushAliyuncsComPushNotification2015-03-18Request) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *PushAliyuncsComPushNotification2015-03-18Request) SetAccount(account string) error {
    r.account = account
    r.Set("Account", account)
    return nil
}

func (r PushAliyuncsComPushNotification2015-03-18Request) GetAccount() string {
    return r.account
}

func (r *PushAliyuncsComPushNotification2015-03-18Request) SetAndroidExtraMap(androidExtraMap string) error {
    r.androidExtraMap = androidExtraMap
    r.Set("AndroidExtraMap", androidExtraMap)
    return nil
}

func (r PushAliyuncsComPushNotification2015-03-18Request) GetAndroidExtraMap() string {
    return r.androidExtraMap
}

func (r *PushAliyuncsComPushNotification2015-03-18Request) SetAndroidMusic(androidMusic string) error {
    r.androidMusic = androidMusic
    r.Set("AndroidMusic", androidMusic)
    return nil
}

func (r PushAliyuncsComPushNotification2015-03-18Request) GetAndroidMusic() string {
    return r.androidMusic
}

func (r *PushAliyuncsComPushNotification2015-03-18Request) SetAndroidNotifyType(androidNotifyType int64) error {
    r.androidNotifyType = androidNotifyType
    r.Set("AndroidNotifyType", androidNotifyType)
    return nil
}

func (r PushAliyuncsComPushNotification2015-03-18Request) GetAndroidNotifyType() int64 {
    return r.androidNotifyType
}

func (r *PushAliyuncsComPushNotification2015-03-18Request) SetAndroidOpenActivity(androidOpenActivity string) error {
    r.androidOpenActivity = androidOpenActivity
    r.Set("AndroidOpenActivity", androidOpenActivity)
    return nil
}

func (r PushAliyuncsComPushNotification2015-03-18Request) GetAndroidOpenActivity() string {
    return r.androidOpenActivity
}

func (r *PushAliyuncsComPushNotification2015-03-18Request) SetAndroidOpenType(androidOpenType int64) error {
    r.androidOpenType = androidOpenType
    r.Set("AndroidOpenType", androidOpenType)
    return nil
}

func (r PushAliyuncsComPushNotification2015-03-18Request) GetAndroidOpenType() int64 {
    return r.androidOpenType
}

func (r *PushAliyuncsComPushNotification2015-03-18Request) SetAndroidOpenUrl(androidOpenUrl string) error {
    r.androidOpenUrl = androidOpenUrl
    r.Set("AndroidOpenUrl", androidOpenUrl)
    return nil
}

func (r PushAliyuncsComPushNotification2015-03-18Request) GetAndroidOpenUrl() string {
    return r.androidOpenUrl
}

func (r *PushAliyuncsComPushNotification2015-03-18Request) SetAntiHarassDuration(antiHarassDuration int64) error {
    r.antiHarassDuration = antiHarassDuration
    r.Set("AntiHarassDuration", antiHarassDuration)
    return nil
}

func (r PushAliyuncsComPushNotification2015-03-18Request) GetAntiHarassDuration() int64 {
    return r.antiHarassDuration
}

func (r *PushAliyuncsComPushNotification2015-03-18Request) SetAntiHarassStartTime(antiHarassStartTime int64) error {
    r.antiHarassStartTime = antiHarassStartTime
    r.Set("AntiHarassStartTime", antiHarassStartTime)
    return nil
}

func (r PushAliyuncsComPushNotification2015-03-18Request) GetAntiHarassStartTime() int64 {
    return r.antiHarassStartTime
}

func (r *PushAliyuncsComPushNotification2015-03-18Request) SetAppId(appId int64) error {
    r.appId = appId
    r.Set("AppId", appId)
    return nil
}

func (r PushAliyuncsComPushNotification2015-03-18Request) GetAppId() int64 {
    return r.appId
}

func (r *PushAliyuncsComPushNotification2015-03-18Request) SetBatchNumber(batchNumber string) error {
    r.batchNumber = batchNumber
    r.Set("BatchNumber", batchNumber)
    return nil
}

func (r PushAliyuncsComPushNotification2015-03-18Request) GetBatchNumber() string {
    return r.batchNumber
}

func (r *PushAliyuncsComPushNotification2015-03-18Request) SetDeviceId(deviceId string) error {
    r.deviceId = deviceId
    r.Set("DeviceId", deviceId)
    return nil
}

func (r PushAliyuncsComPushNotification2015-03-18Request) GetDeviceId() string {
    return r.deviceId
}

func (r *PushAliyuncsComPushNotification2015-03-18Request) SetDeviceType(deviceType int64) error {
    r.deviceType = deviceType
    r.Set("DeviceType", deviceType)
    return nil
}

func (r PushAliyuncsComPushNotification2015-03-18Request) GetDeviceType() int64 {
    return r.deviceType
}

func (r *PushAliyuncsComPushNotification2015-03-18Request) SetIosExtraMap(iosExtraMap string) error {
    r.iosExtraMap = iosExtraMap
    r.Set("IosExtraMap", iosExtraMap)
    return nil
}

func (r PushAliyuncsComPushNotification2015-03-18Request) GetIosExtraMap() string {
    return r.iosExtraMap
}

func (r *PushAliyuncsComPushNotification2015-03-18Request) SetIosFooter(iosFooter int64) error {
    r.iosFooter = iosFooter
    r.Set("IosFooter", iosFooter)
    return nil
}

func (r PushAliyuncsComPushNotification2015-03-18Request) GetIosFooter() int64 {
    return r.iosFooter
}

func (r *PushAliyuncsComPushNotification2015-03-18Request) SetIosMusic(iosMusic string) error {
    r.iosMusic = iosMusic
    r.Set("IosMusic", iosMusic)
    return nil
}

func (r PushAliyuncsComPushNotification2015-03-18Request) GetIosMusic() string {
    return r.iosMusic
}

func (r *PushAliyuncsComPushNotification2015-03-18Request) SetPushTime(pushTime string) error {
    r.pushTime = pushTime
    r.Set("PushTime", pushTime)
    return nil
}

func (r PushAliyuncsComPushNotification2015-03-18Request) GetPushTime() string {
    return r.pushTime
}

func (r *PushAliyuncsComPushNotification2015-03-18Request) SetSendType(sendType int64) error {
    r.sendType = sendType
    r.Set("SendType", sendType)
    return nil
}

func (r PushAliyuncsComPushNotification2015-03-18Request) GetSendType() int64 {
    return r.sendType
}

func (r *PushAliyuncsComPushNotification2015-03-18Request) SetSummary(summary string) error {
    r.summary = summary
    r.Set("Summary", summary)
    return nil
}

func (r PushAliyuncsComPushNotification2015-03-18Request) GetSummary() string {
    return r.summary
}

func (r *PushAliyuncsComPushNotification2015-03-18Request) SetTag(tag string) error {
    r.tag = tag
    r.Set("Tag", tag)
    return nil
}

func (r PushAliyuncsComPushNotification2015-03-18Request) GetTag() string {
    return r.tag
}

func (r *PushAliyuncsComPushNotification2015-03-18Request) SetTimeout(timeout int64) error {
    r.timeout = timeout
    r.Set("Timeout", timeout)
    return nil
}

func (r PushAliyuncsComPushNotification2015-03-18Request) GetTimeout() int64 {
    return r.timeout
}

func (r *PushAliyuncsComPushNotification2015-03-18Request) SetTitle(title string) error {
    r.title = title
    r.Set("Title", title)
    return nil
}

func (r PushAliyuncsComPushNotification2015-03-18Request) GetTitle() string {
    return r.title
}

