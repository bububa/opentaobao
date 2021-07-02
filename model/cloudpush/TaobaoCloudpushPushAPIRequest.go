package cloudpush

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoCloudpushPushAPIRequest 百川用户使用云推送高级推送接口 API请求
// taobao.cloudpush.push
//
// 百川用户使用云推送高级推送接口
type TaobaoCloudpushPushAPIRequest struct {
	model.Params
	// 推送目标: device:推送给设备; account:推送给指定帐号,tag:推送给自定义标签; all: 推送给全部
	_target string
	// 根据Target来设定，如Target=device, 则对应的值为 设备id1,设备id2. 多个值使用逗号分隔.(帐号与设备有一次最多100个的限制)
	_targetValue string
	// Android对应的activity,仅仅当androidOpenType=2有效
	_androidActivity string
	// 自定义的kv结构,开发者扩展用 针对android
	_androidExtParameters string
	// android通知声音
	_androidMusic string
	// 点击通知后动作,1:打开应用 2: 打开应用Activity 3:打开 url
	_androidOpenType string
	// Android收到推送后打开对应的url,仅仅当androidOpenType=3有效
	_androidOpenUrl string
	// 防打扰时长,取值范围为1~23
	_antiHarassDuration int64
	// 防打扰开始时间点,取值范围为0~23
	_antiHarassStartTime int64
	// 批次编号,用于活动效果统计
	_batchNumber string
	// 推送内容
	_body string
	// 设备类型,取值范围为:0~3云推送支持多种设备,各种设备类型编号如下:    iOS设备:deviceType=0; Andriod设备:deviceType=1;如果存在此字段,则向指定的设备类型推送消息。 默认为全部(3);
	_deviceType int64
	// iOS应用图标右上角角标
	_iosBadge string
	// 自定义的kv结构,开发者扩展用 针对iOS设备
	_iosExtParameters string
	// iOS通知声音
	_iosMusic string
	// 当APP不在线时候，是否通过通知提醒.  针对不同设备，处理逻辑不同。 该参数只针对iOS设备生效， (remind=true  & 发送消息的话(type=0)). 当你的目标设备不在线(既长连接通道不通, 我们会将这条消息的标题，通过苹果的apns通道再送达一次。发apns是发送生产环境的apns，需要在云推送配置的app的iOS生产证书和密码需要正确，否则也发送不了。 (remind=false & 并且是发送消息的话(type=0))，那么设备不在线，则不会再走苹果apns发送了。
	_remind bool
	// 离线消息是否保存,若保存, 在推送时候，用户即使不在线，下一次上线则会收到
	_storeOffline bool
	// 通知的摘要
	_summery string
	// 离线消息保存时长,取值范围为1~72,若不填,则表示不保存离线消息
	_timeout int64
	// 推送的标题内容.
	_title string
	// 0:表示消息(默认为0),1:表示通知
	_type int64
}

// NewTaobaoCloudpushPushRequest 初始化TaobaoCloudpushPushAPIRequest对象
func NewTaobaoCloudpushPushRequest() *TaobaoCloudpushPushAPIRequest {
	return &TaobaoCloudpushPushAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoCloudpushPushAPIRequest) GetApiMethodName() string {
	return "taobao.cloudpush.push"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoCloudpushPushAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Target Setter
// 推送目标: device:推送给设备; account:推送给指定帐号,tag:推送给自定义标签; all: 推送给全部
func (r *TaobaoCloudpushPushAPIRequest) SetTarget(_target string) error {
	r._target = _target
	r.Set("target", _target)
	return nil
}

// Get Target Getter
func (r TaobaoCloudpushPushAPIRequest) GetTarget() string {
	return r._target
}

// Set is TargetValue Setter
// 根据Target来设定，如Target=device, 则对应的值为 设备id1,设备id2. 多个值使用逗号分隔.(帐号与设备有一次最多100个的限制)
func (r *TaobaoCloudpushPushAPIRequest) SetTargetValue(_targetValue string) error {
	r._targetValue = _targetValue
	r.Set("target_value", _targetValue)
	return nil
}

// Get TargetValue Getter
func (r TaobaoCloudpushPushAPIRequest) GetTargetValue() string {
	return r._targetValue
}

// Set is AndroidActivity Setter
// Android对应的activity,仅仅当androidOpenType=2有效
func (r *TaobaoCloudpushPushAPIRequest) SetAndroidActivity(_androidActivity string) error {
	r._androidActivity = _androidActivity
	r.Set("android_activity", _androidActivity)
	return nil
}

// Get AndroidActivity Getter
func (r TaobaoCloudpushPushAPIRequest) GetAndroidActivity() string {
	return r._androidActivity
}

// Set is AndroidExtParameters Setter
// 自定义的kv结构,开发者扩展用 针对android
func (r *TaobaoCloudpushPushAPIRequest) SetAndroidExtParameters(_androidExtParameters string) error {
	r._androidExtParameters = _androidExtParameters
	r.Set("android_ext_parameters", _androidExtParameters)
	return nil
}

// Get AndroidExtParameters Getter
func (r TaobaoCloudpushPushAPIRequest) GetAndroidExtParameters() string {
	return r._androidExtParameters
}

// Set is AndroidMusic Setter
// android通知声音
func (r *TaobaoCloudpushPushAPIRequest) SetAndroidMusic(_androidMusic string) error {
	r._androidMusic = _androidMusic
	r.Set("android_music", _androidMusic)
	return nil
}

// Get AndroidMusic Getter
func (r TaobaoCloudpushPushAPIRequest) GetAndroidMusic() string {
	return r._androidMusic
}

// Set is AndroidOpenType Setter
// 点击通知后动作,1:打开应用 2: 打开应用Activity 3:打开 url
func (r *TaobaoCloudpushPushAPIRequest) SetAndroidOpenType(_androidOpenType string) error {
	r._androidOpenType = _androidOpenType
	r.Set("android_open_type", _androidOpenType)
	return nil
}

// Get AndroidOpenType Getter
func (r TaobaoCloudpushPushAPIRequest) GetAndroidOpenType() string {
	return r._androidOpenType
}

// Set is AndroidOpenUrl Setter
// Android收到推送后打开对应的url,仅仅当androidOpenType=3有效
func (r *TaobaoCloudpushPushAPIRequest) SetAndroidOpenUrl(_androidOpenUrl string) error {
	r._androidOpenUrl = _androidOpenUrl
	r.Set("android_open_url", _androidOpenUrl)
	return nil
}

// Get AndroidOpenUrl Getter
func (r TaobaoCloudpushPushAPIRequest) GetAndroidOpenUrl() string {
	return r._androidOpenUrl
}

// Set is AntiHarassDuration Setter
// 防打扰时长,取值范围为1~23
func (r *TaobaoCloudpushPushAPIRequest) SetAntiHarassDuration(_antiHarassDuration int64) error {
	r._antiHarassDuration = _antiHarassDuration
	r.Set("anti_harass_duration", _antiHarassDuration)
	return nil
}

// Get AntiHarassDuration Getter
func (r TaobaoCloudpushPushAPIRequest) GetAntiHarassDuration() int64 {
	return r._antiHarassDuration
}

// Set is AntiHarassStartTime Setter
// 防打扰开始时间点,取值范围为0~23
func (r *TaobaoCloudpushPushAPIRequest) SetAntiHarassStartTime(_antiHarassStartTime int64) error {
	r._antiHarassStartTime = _antiHarassStartTime
	r.Set("anti_harass_start_time", _antiHarassStartTime)
	return nil
}

// Get AntiHarassStartTime Getter
func (r TaobaoCloudpushPushAPIRequest) GetAntiHarassStartTime() int64 {
	return r._antiHarassStartTime
}

// Set is BatchNumber Setter
// 批次编号,用于活动效果统计
func (r *TaobaoCloudpushPushAPIRequest) SetBatchNumber(_batchNumber string) error {
	r._batchNumber = _batchNumber
	r.Set("batch_number", _batchNumber)
	return nil
}

// Get BatchNumber Getter
func (r TaobaoCloudpushPushAPIRequest) GetBatchNumber() string {
	return r._batchNumber
}

// Set is Body Setter
// 推送内容
func (r *TaobaoCloudpushPushAPIRequest) SetBody(_body string) error {
	r._body = _body
	r.Set("body", _body)
	return nil
}

// Get Body Getter
func (r TaobaoCloudpushPushAPIRequest) GetBody() string {
	return r._body
}

// Set is DeviceType Setter
// 设备类型,取值范围为:0~3云推送支持多种设备,各种设备类型编号如下:    iOS设备:deviceType=0; Andriod设备:deviceType=1;如果存在此字段,则向指定的设备类型推送消息。 默认为全部(3);
func (r *TaobaoCloudpushPushAPIRequest) SetDeviceType(_deviceType int64) error {
	r._deviceType = _deviceType
	r.Set("device_type", _deviceType)
	return nil
}

// Get DeviceType Getter
func (r TaobaoCloudpushPushAPIRequest) GetDeviceType() int64 {
	return r._deviceType
}

// Set is IosBadge Setter
// iOS应用图标右上角角标
func (r *TaobaoCloudpushPushAPIRequest) SetIosBadge(_iosBadge string) error {
	r._iosBadge = _iosBadge
	r.Set("ios_badge", _iosBadge)
	return nil
}

// Get IosBadge Getter
func (r TaobaoCloudpushPushAPIRequest) GetIosBadge() string {
	return r._iosBadge
}

// Set is IosExtParameters Setter
// 自定义的kv结构,开发者扩展用 针对iOS设备
func (r *TaobaoCloudpushPushAPIRequest) SetIosExtParameters(_iosExtParameters string) error {
	r._iosExtParameters = _iosExtParameters
	r.Set("ios_ext_parameters", _iosExtParameters)
	return nil
}

// Get IosExtParameters Getter
func (r TaobaoCloudpushPushAPIRequest) GetIosExtParameters() string {
	return r._iosExtParameters
}

// Set is IosMusic Setter
// iOS通知声音
func (r *TaobaoCloudpushPushAPIRequest) SetIosMusic(_iosMusic string) error {
	r._iosMusic = _iosMusic
	r.Set("ios_music", _iosMusic)
	return nil
}

// Get IosMusic Getter
func (r TaobaoCloudpushPushAPIRequest) GetIosMusic() string {
	return r._iosMusic
}

// Set is Remind Setter
// 当APP不在线时候，是否通过通知提醒.  针对不同设备，处理逻辑不同。 该参数只针对iOS设备生效， (remind=true  & 发送消息的话(type=0)). 当你的目标设备不在线(既长连接通道不通, 我们会将这条消息的标题，通过苹果的apns通道再送达一次。发apns是发送生产环境的apns，需要在云推送配置的app的iOS生产证书和密码需要正确，否则也发送不了。 (remind=false & 并且是发送消息的话(type=0))，那么设备不在线，则不会再走苹果apns发送了。
func (r *TaobaoCloudpushPushAPIRequest) SetRemind(_remind bool) error {
	r._remind = _remind
	r.Set("remind", _remind)
	return nil
}

// Get Remind Getter
func (r TaobaoCloudpushPushAPIRequest) GetRemind() bool {
	return r._remind
}

// Set is StoreOffline Setter
// 离线消息是否保存,若保存, 在推送时候，用户即使不在线，下一次上线则会收到
func (r *TaobaoCloudpushPushAPIRequest) SetStoreOffline(_storeOffline bool) error {
	r._storeOffline = _storeOffline
	r.Set("store_offline", _storeOffline)
	return nil
}

// Get StoreOffline Getter
func (r TaobaoCloudpushPushAPIRequest) GetStoreOffline() bool {
	return r._storeOffline
}

// Set is Summery Setter
// 通知的摘要
func (r *TaobaoCloudpushPushAPIRequest) SetSummery(_summery string) error {
	r._summery = _summery
	r.Set("summery", _summery)
	return nil
}

// Get Summery Getter
func (r TaobaoCloudpushPushAPIRequest) GetSummery() string {
	return r._summery
}

// Set is Timeout Setter
// 离线消息保存时长,取值范围为1~72,若不填,则表示不保存离线消息
func (r *TaobaoCloudpushPushAPIRequest) SetTimeout(_timeout int64) error {
	r._timeout = _timeout
	r.Set("timeout", _timeout)
	return nil
}

// Get Timeout Getter
func (r TaobaoCloudpushPushAPIRequest) GetTimeout() int64 {
	return r._timeout
}

// Set is Title Setter
// 推送的标题内容.
func (r *TaobaoCloudpushPushAPIRequest) SetTitle(_title string) error {
	r._title = _title
	r.Set("title", _title)
	return nil
}

// Get Title Getter
func (r TaobaoCloudpushPushAPIRequest) GetTitle() string {
	return r._title
}

// Set is Type Setter
// 0:表示消息(默认为0),1:表示通知
func (r *TaobaoCloudpushPushAPIRequest) SetType(_type int64) error {
	r._type = _type
	r.Set("type", _type)
	return nil
}

// Get Type Getter
func (r TaobaoCloudpushPushAPIRequest) GetType() int64 {
	return r._type
}
