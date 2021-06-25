package cloudpush

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
百川用户使用云推送高级推送接口 APIRequest
taobao.cloudpush.push

百川用户使用云推送高级推送接口
*/
type TaobaoCloudpushPushRequest struct {
    model.Params

    // 推送目标: device:推送给设备; account:推送给指定帐号,tag:推送给自定义标签; all: 推送给全部
    target   string 

    // 根据Target来设定，如Target=device, 则对应的值为 设备id1,设备id2. 多个值使用逗号分隔.(帐号与设备有一次最多100个的限制)
    targetValue   string 

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

    // 批次编号,用于活动效果统计
    batchNumber   string 

    // 推送内容
    body   string 

    // 设备类型,取值范围为:0~3云推送支持多种设备,各种设备类型编号如下:    iOS设备:deviceType=0; Andriod设备:deviceType=1;如果存在此字段,则向指定的设备类型推送消息。 默认为全部(3);
    deviceType   int64 

    // iOS应用图标右上角角标
    iosBadge   string 

    // 自定义的kv结构,开发者扩展用 针对iOS设备
    iosExtParameters   string 

    // iOS通知声音
    iosMusic   string 

    // 当APP不在线时候，是否通过通知提醒.  针对不同设备，处理逻辑不同。 该参数只针对iOS设备生效， (remind=true  & 发送消息的话(type=0)). 当你的目标设备不在线(既长连接通道不通, 我们会将这条消息的标题，通过苹果的apns通道再送达一次。发apns是发送生产环境的apns，需要在云推送配置的app的iOS生产证书和密码需要正确，否则也发送不了。 (remind=false & 并且是发送消息的话(type=0))，那么设备不在线，则不会再走苹果apns发送了。
    remind   bool 

    // 离线消息是否保存,若保存, 在推送时候，用户即使不在线，下一次上线则会收到
    storeOffline   bool 

    // 通知的摘要
    summery   string 

    // 离线消息保存时长,取值范围为1~72,若不填,则表示不保存离线消息
    timeout   int64 

    // 推送的标题内容.
    title   string 

    // 0:表示消息(默认为0),1:表示通知
    type   int64 

}

func NewTaobaoCloudpushPushRequest() *TaobaoCloudpushPushRequest{
    return &TaobaoCloudpushPushRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoCloudpushPushRequest) GetApiMethodName() string {
    return "taobao.cloudpush.push"
}

func (r TaobaoCloudpushPushRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoCloudpushPushRequest) SetTarget(target string) error {
    r.target = target
    r.Set("target", target)
    return nil
}

func (r TaobaoCloudpushPushRequest) GetTarget() string {
    return r.target
}

func (r *TaobaoCloudpushPushRequest) SetTargetValue(targetValue string) error {
    r.targetValue = targetValue
    r.Set("target_value", targetValue)
    return nil
}

func (r TaobaoCloudpushPushRequest) GetTargetValue() string {
    return r.targetValue
}

func (r *TaobaoCloudpushPushRequest) SetAndroidActivity(androidActivity string) error {
    r.androidActivity = androidActivity
    r.Set("android_activity", androidActivity)
    return nil
}

func (r TaobaoCloudpushPushRequest) GetAndroidActivity() string {
    return r.androidActivity
}

func (r *TaobaoCloudpushPushRequest) SetAndroidExtParameters(androidExtParameters string) error {
    r.androidExtParameters = androidExtParameters
    r.Set("android_ext_parameters", androidExtParameters)
    return nil
}

func (r TaobaoCloudpushPushRequest) GetAndroidExtParameters() string {
    return r.androidExtParameters
}

func (r *TaobaoCloudpushPushRequest) SetAndroidMusic(androidMusic string) error {
    r.androidMusic = androidMusic
    r.Set("android_music", androidMusic)
    return nil
}

func (r TaobaoCloudpushPushRequest) GetAndroidMusic() string {
    return r.androidMusic
}

func (r *TaobaoCloudpushPushRequest) SetAndroidOpenType(androidOpenType string) error {
    r.androidOpenType = androidOpenType
    r.Set("android_open_type", androidOpenType)
    return nil
}

func (r TaobaoCloudpushPushRequest) GetAndroidOpenType() string {
    return r.androidOpenType
}

func (r *TaobaoCloudpushPushRequest) SetAndroidOpenUrl(androidOpenUrl string) error {
    r.androidOpenUrl = androidOpenUrl
    r.Set("android_open_url", androidOpenUrl)
    return nil
}

func (r TaobaoCloudpushPushRequest) GetAndroidOpenUrl() string {
    return r.androidOpenUrl
}

func (r *TaobaoCloudpushPushRequest) SetAntiHarassDuration(antiHarassDuration int64) error {
    r.antiHarassDuration = antiHarassDuration
    r.Set("anti_harass_duration", antiHarassDuration)
    return nil
}

func (r TaobaoCloudpushPushRequest) GetAntiHarassDuration() int64 {
    return r.antiHarassDuration
}

func (r *TaobaoCloudpushPushRequest) SetAntiHarassStartTime(antiHarassStartTime int64) error {
    r.antiHarassStartTime = antiHarassStartTime
    r.Set("anti_harass_start_time", antiHarassStartTime)
    return nil
}

func (r TaobaoCloudpushPushRequest) GetAntiHarassStartTime() int64 {
    return r.antiHarassStartTime
}

func (r *TaobaoCloudpushPushRequest) SetBatchNumber(batchNumber string) error {
    r.batchNumber = batchNumber
    r.Set("batch_number", batchNumber)
    return nil
}

func (r TaobaoCloudpushPushRequest) GetBatchNumber() string {
    return r.batchNumber
}

func (r *TaobaoCloudpushPushRequest) SetBody(body string) error {
    r.body = body
    r.Set("body", body)
    return nil
}

func (r TaobaoCloudpushPushRequest) GetBody() string {
    return r.body
}

func (r *TaobaoCloudpushPushRequest) SetDeviceType(deviceType int64) error {
    r.deviceType = deviceType
    r.Set("device_type", deviceType)
    return nil
}

func (r TaobaoCloudpushPushRequest) GetDeviceType() int64 {
    return r.deviceType
}

func (r *TaobaoCloudpushPushRequest) SetIosBadge(iosBadge string) error {
    r.iosBadge = iosBadge
    r.Set("ios_badge", iosBadge)
    return nil
}

func (r TaobaoCloudpushPushRequest) GetIosBadge() string {
    return r.iosBadge
}

func (r *TaobaoCloudpushPushRequest) SetIosExtParameters(iosExtParameters string) error {
    r.iosExtParameters = iosExtParameters
    r.Set("ios_ext_parameters", iosExtParameters)
    return nil
}

func (r TaobaoCloudpushPushRequest) GetIosExtParameters() string {
    return r.iosExtParameters
}

func (r *TaobaoCloudpushPushRequest) SetIosMusic(iosMusic string) error {
    r.iosMusic = iosMusic
    r.Set("ios_music", iosMusic)
    return nil
}

func (r TaobaoCloudpushPushRequest) GetIosMusic() string {
    return r.iosMusic
}

func (r *TaobaoCloudpushPushRequest) SetRemind(remind bool) error {
    r.remind = remind
    r.Set("remind", remind)
    return nil
}

func (r TaobaoCloudpushPushRequest) GetRemind() bool {
    return r.remind
}

func (r *TaobaoCloudpushPushRequest) SetStoreOffline(storeOffline bool) error {
    r.storeOffline = storeOffline
    r.Set("store_offline", storeOffline)
    return nil
}

func (r TaobaoCloudpushPushRequest) GetStoreOffline() bool {
    return r.storeOffline
}

func (r *TaobaoCloudpushPushRequest) SetSummery(summery string) error {
    r.summery = summery
    r.Set("summery", summery)
    return nil
}

func (r TaobaoCloudpushPushRequest) GetSummery() string {
    return r.summery
}

func (r *TaobaoCloudpushPushRequest) SetTimeout(timeout int64) error {
    r.timeout = timeout
    r.Set("timeout", timeout)
    return nil
}

func (r TaobaoCloudpushPushRequest) GetTimeout() int64 {
    return r.timeout
}

func (r *TaobaoCloudpushPushRequest) SetTitle(title string) error {
    r.title = title
    r.Set("title", title)
    return nil
}

func (r TaobaoCloudpushPushRequest) GetTitle() string {
    return r.title
}

func (r *TaobaoCloudpushPushRequest) SetType(type int64) error {
    r.type = type
    r.Set("type", type)
    return nil
}

func (r TaobaoCloudpushPushRequest) GetType() int64 {
    return r.type
}

