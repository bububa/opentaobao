package aliyun

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
云推送指令API APIRequest
push.aliyuncs.com.push.20150518

阿里云推送新增API，允许一条推送指令同时发布到多个终端上。
*/
type PushAliyuncsComPush20150518Request struct {
    model.Params

    // 用户账号列表,以换行区分,仅sendType为3时有效
    account   string 

    // Android对应的activity,仅仅当androidOpenType=2有效
    androidActivity   string 

    // 自定义的kv结构,开发者扩展用 针对android
    androidExtParameters   string 

    // android通知声音
    androidMusic   string 

    // 点击通知后动作,1:打开应用 2: 打开应用Activity 3:打开 url
    androidOpenType   string 

    // Android收到推送后打开对应的url,仅仅当androidOpenType=3有效
    androidOpenUrl   string 

    // 防打扰时长,取值范围为1~23
    antiHarassDuration   int64 

    // 防打扰开始时间点,取值范围为0~23
    antiHarassStartTime   int64 

    // 应用标识
    appId   int64 

    // 批次编号,用于活动效果统计
    batchNumber   string 

    // 推送内容
    body   string 

    // 推送接收设备，多个以逗号分隔
    deviceId   string 

    // 设备类型,取值范围为:0~3云推送支持多种设备,各 种设备类型编号如下: 0:IOS设备; 1:Andriod设备 3:全部. 默认为3.
    deviceType   int64 

    // iOS应用图标右上角角标
    iOSBadge   string 

    // 自定义的kv结构,开发者扩展用 针对iOS设备
    iOSExtParameters   string 

    // iOS通知声音
    iOSMusic   string 

    // 推送时间,若空表示立即推送,推送时间不能早于当前时间
    pushTime   string 

    // 当APP不在线时候，是否通过通知提醒
    remind   bool 

    // 推送类型,取值范围:1~4; 1:所有人,无需指定tag、 deviceType等2:一群人,必须指定tag3:指定用户,根据 用户账号列表文件发送消息4:指定设备,根据设备编码列 表文件发送消息默认值为1
    sendType   int64 

    // 离线消息是否保存,若保存, 在推送时候，用户即使不在线，下一次上线则会收到
    storeOffline   bool 

    // 通知的摘要
    summery   string 

    // 离线消息保存时长,取值范围为1~72,若不填,则表示不保存离线消息
    timeout   int64 

    // 推送的标题内容.
    title   string 

}

func NewPushAliyuncsComPush20150518Request() *PushAliyuncsComPush20150518Request{
    return &PushAliyuncsComPush20150518Request{
        Params: model.NewParams(),
    }
}

func (r PushAliyuncsComPush20150518Request) GetApiMethodName() string {
    return "push.aliyuncs.com.push.20150518"
}

func (r PushAliyuncsComPush20150518Request) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *PushAliyuncsComPush20150518Request) SetAccount(account string) error {
    r.account = account
    r.Set("Account", account)
    return nil
}

func (r PushAliyuncsComPush20150518Request) GetAccount() string {
    return r.account
}

func (r *PushAliyuncsComPush20150518Request) SetAndroidActivity(androidActivity string) error {
    r.androidActivity = androidActivity
    r.Set("AndroidActivity", androidActivity)
    return nil
}

func (r PushAliyuncsComPush20150518Request) GetAndroidActivity() string {
    return r.androidActivity
}

func (r *PushAliyuncsComPush20150518Request) SetAndroidExtParameters(androidExtParameters string) error {
    r.androidExtParameters = androidExtParameters
    r.Set("AndroidExtParameters", androidExtParameters)
    return nil
}

func (r PushAliyuncsComPush20150518Request) GetAndroidExtParameters() string {
    return r.androidExtParameters
}

func (r *PushAliyuncsComPush20150518Request) SetAndroidMusic(androidMusic string) error {
    r.androidMusic = androidMusic
    r.Set("AndroidMusic", androidMusic)
    return nil
}

func (r PushAliyuncsComPush20150518Request) GetAndroidMusic() string {
    return r.androidMusic
}

func (r *PushAliyuncsComPush20150518Request) SetAndroidOpenType(androidOpenType string) error {
    r.androidOpenType = androidOpenType
    r.Set("AndroidOpenType", androidOpenType)
    return nil
}

func (r PushAliyuncsComPush20150518Request) GetAndroidOpenType() string {
    return r.androidOpenType
}

func (r *PushAliyuncsComPush20150518Request) SetAndroidOpenUrl(androidOpenUrl string) error {
    r.androidOpenUrl = androidOpenUrl
    r.Set("AndroidOpenUrl", androidOpenUrl)
    return nil
}

func (r PushAliyuncsComPush20150518Request) GetAndroidOpenUrl() string {
    return r.androidOpenUrl
}

func (r *PushAliyuncsComPush20150518Request) SetAntiHarassDuration(antiHarassDuration int64) error {
    r.antiHarassDuration = antiHarassDuration
    r.Set("AntiHarassDuration", antiHarassDuration)
    return nil
}

func (r PushAliyuncsComPush20150518Request) GetAntiHarassDuration() int64 {
    return r.antiHarassDuration
}

func (r *PushAliyuncsComPush20150518Request) SetAntiHarassStartTime(antiHarassStartTime int64) error {
    r.antiHarassStartTime = antiHarassStartTime
    r.Set("AntiHarassStartTime", antiHarassStartTime)
    return nil
}

func (r PushAliyuncsComPush20150518Request) GetAntiHarassStartTime() int64 {
    return r.antiHarassStartTime
}

func (r *PushAliyuncsComPush20150518Request) SetAppId(appId int64) error {
    r.appId = appId
    r.Set("AppId", appId)
    return nil
}

func (r PushAliyuncsComPush20150518Request) GetAppId() int64 {
    return r.appId
}

func (r *PushAliyuncsComPush20150518Request) SetBatchNumber(batchNumber string) error {
    r.batchNumber = batchNumber
    r.Set("BatchNumber", batchNumber)
    return nil
}

func (r PushAliyuncsComPush20150518Request) GetBatchNumber() string {
    return r.batchNumber
}

func (r *PushAliyuncsComPush20150518Request) SetBody(body string) error {
    r.body = body
    r.Set("Body", body)
    return nil
}

func (r PushAliyuncsComPush20150518Request) GetBody() string {
    return r.body
}

func (r *PushAliyuncsComPush20150518Request) SetDeviceId(deviceId string) error {
    r.deviceId = deviceId
    r.Set("DeviceId", deviceId)
    return nil
}

func (r PushAliyuncsComPush20150518Request) GetDeviceId() string {
    return r.deviceId
}

func (r *PushAliyuncsComPush20150518Request) SetDeviceType(deviceType int64) error {
    r.deviceType = deviceType
    r.Set("DeviceType", deviceType)
    return nil
}

func (r PushAliyuncsComPush20150518Request) GetDeviceType() int64 {
    return r.deviceType
}

func (r *PushAliyuncsComPush20150518Request) SetIOSBadge(iOSBadge string) error {
    r.iOSBadge = iOSBadge
    r.Set("IOSBadge", iOSBadge)
    return nil
}

func (r PushAliyuncsComPush20150518Request) GetIOSBadge() string {
    return r.iOSBadge
}

func (r *PushAliyuncsComPush20150518Request) SetIOSExtParameters(iOSExtParameters string) error {
    r.iOSExtParameters = iOSExtParameters
    r.Set("IOSExtParameters", iOSExtParameters)
    return nil
}

func (r PushAliyuncsComPush20150518Request) GetIOSExtParameters() string {
    return r.iOSExtParameters
}

func (r *PushAliyuncsComPush20150518Request) SetIOSMusic(iOSMusic string) error {
    r.iOSMusic = iOSMusic
    r.Set("IOSMusic", iOSMusic)
    return nil
}

func (r PushAliyuncsComPush20150518Request) GetIOSMusic() string {
    return r.iOSMusic
}

func (r *PushAliyuncsComPush20150518Request) SetPushTime(pushTime string) error {
    r.pushTime = pushTime
    r.Set("PushTime", pushTime)
    return nil
}

func (r PushAliyuncsComPush20150518Request) GetPushTime() string {
    return r.pushTime
}

func (r *PushAliyuncsComPush20150518Request) SetRemind(remind bool) error {
    r.remind = remind
    r.Set("Remind", remind)
    return nil
}

func (r PushAliyuncsComPush20150518Request) GetRemind() bool {
    return r.remind
}

func (r *PushAliyuncsComPush20150518Request) SetSendType(sendType int64) error {
    r.sendType = sendType
    r.Set("SendType", sendType)
    return nil
}

func (r PushAliyuncsComPush20150518Request) GetSendType() int64 {
    return r.sendType
}

func (r *PushAliyuncsComPush20150518Request) SetStoreOffline(storeOffline bool) error {
    r.storeOffline = storeOffline
    r.Set("StoreOffline", storeOffline)
    return nil
}

func (r PushAliyuncsComPush20150518Request) GetStoreOffline() bool {
    return r.storeOffline
}

func (r *PushAliyuncsComPush20150518Request) SetSummery(summery string) error {
    r.summery = summery
    r.Set("Summery", summery)
    return nil
}

func (r PushAliyuncsComPush20150518Request) GetSummery() string {
    return r.summery
}

func (r *PushAliyuncsComPush20150518Request) SetTimeout(timeout int64) error {
    r.timeout = timeout
    r.Set("Timeout", timeout)
    return nil
}

func (r PushAliyuncsComPush20150518Request) GetTimeout() int64 {
    return r.timeout
}

func (r *PushAliyuncsComPush20150518Request) SetTitle(title string) error {
    r.title = title
    r.Set("Title", title)
    return nil
}

func (r PushAliyuncsComPush20150518Request) GetTitle() string {
    return r.title
}

