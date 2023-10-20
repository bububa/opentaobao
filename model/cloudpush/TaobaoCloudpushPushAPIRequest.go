package cloudpush

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaocloudpushpushAPIRequest 百川用户使用云推送高级推送接口 API请求
// taobao.cloudpush.push
//
// 百川用户使用云推送高级推送接口
type TaobaocloudpushpushAPIRequest struct {
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
	// 批次编号,用于活动效果统计
	_batchNumber string
	// 推送内容
	_body string
	// iOS应用图标右上角角标
	_iosBadge string
	// 自定义的kv结构,开发者扩展用 针对iOS设备
	_iosExtParameters string
	// iOS通知声音
	_iosMusic string
	// 通知的摘要
	_summery string
	// 推送的标题内容.
	_title string
	// 防打扰时长,取值范围为1~23
	_antiHarassDuration int64
	// 防打扰开始时间点,取值范围为0~23
	_antiHarassStartTime int64
	// 设备类型,取值范围为:0~3云推送支持多种设备,各种设备类型编号如下:    iOS设备:deviceType=0; Andriod设备:deviceType=1;如果存在此字段,则向指定的设备类型推送消息。 默认为全部(3);
	_deviceType int64
	// 离线消息保存时长,取值范围为1~72,若不填,则表示不保存离线消息
	_timeout int64
	// 0:表示消息(默认为0),1:表示通知
	_type int64
	// 当APP不在线时候，是否通过通知提醒.  针对不同设备，处理逻辑不同。 该参数只针对iOS设备生效， (remind=true  & 发送消息的话(type=0)). 当你的目标设备不在线(既长连接通道不通, 我们会将这条消息的标题，通过苹果的apns通道再送达一次。发apns是发送生产环境的apns，需要在云推送配置的app的iOS生产证书和密码需要正确，否则也发送不了。 (remind=false & 并且是发送消息的话(type=0))，那么设备不在线，则不会再走苹果apns发送了。
	_remind bool
	// 离线消息是否保存,若保存, 在推送时候，用户即使不在线，下一次上线则会收到
	_storeOffline bool
}

// NewTaobaocloudpushpushRequest 初始化TaobaocloudpushpushAPIRequest对象
func NewTaobaocloudpushpushRequest() *TaobaocloudpushpushAPIRequest {
	return &TaobaocloudpushpushAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaocloudpushpushAPIRequest) GetApiMethodName() string {
	return "taobao.cloudpush.push"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaocloudpushpushAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaocloudpushpushAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTarget is Target Setter
// 推送目标: device:推送给设备; account:推送给指定帐号,tag:推送给自定义标签; all: 推送给全部
func (r *TaobaocloudpushpushAPIRequest) SetTarget(_target string) error {
	r._target = _target
	r.Set("target", _target)
	return nil
}

// GetTarget Target Getter
func (r TaobaocloudpushpushAPIRequest) GetTarget() string {
	return r._target
}

// SetTargetValue is TargetValue Setter
// 根据Target来设定，如Target=device, 则对应的值为 设备id1,设备id2. 多个值使用逗号分隔.(帐号与设备有一次最多100个的限制)
func (r *TaobaocloudpushpushAPIRequest) SetTargetValue(_targetValue string) error {
	r._targetValue = _targetValue
	r.Set("target_value", _targetValue)
	return nil
}

// GetTargetValue TargetValue Getter
func (r TaobaocloudpushpushAPIRequest) GetTargetValue() string {
	return r._targetValue
}

// SetAndroidActivity is AndroidActivity Setter
// Android对应的activity,仅仅当androidOpenType=2有效
func (r *TaobaocloudpushpushAPIRequest) SetAndroidActivity(_androidActivity string) error {
	r._androidActivity = _androidActivity
	r.Set("android_activity", _androidActivity)
	return nil
}

// GetAndroidActivity AndroidActivity Getter
func (r TaobaocloudpushpushAPIRequest) GetAndroidActivity() string {
	return r._androidActivity
}

// SetAndroidExtParameters is AndroidExtParameters Setter
// 自定义的kv结构,开发者扩展用 针对android
func (r *TaobaocloudpushpushAPIRequest) SetAndroidExtParameters(_androidExtParameters string) error {
	r._androidExtParameters = _androidExtParameters
	r.Set("android_ext_parameters", _androidExtParameters)
	return nil
}

// GetAndroidExtParameters AndroidExtParameters Getter
func (r TaobaocloudpushpushAPIRequest) GetAndroidExtParameters() string {
	return r._androidExtParameters
}

// SetAndroidMusic is AndroidMusic Setter
// android通知声音
func (r *TaobaocloudpushpushAPIRequest) SetAndroidMusic(_androidMusic string) error {
	r._androidMusic = _androidMusic
	r.Set("android_music", _androidMusic)
	return nil
}

// GetAndroidMusic AndroidMusic Getter
func (r TaobaocloudpushpushAPIRequest) GetAndroidMusic() string {
	return r._androidMusic
}

// SetAndroidOpenType is AndroidOpenType Setter
// 点击通知后动作,1:打开应用 2: 打开应用Activity 3:打开 url
func (r *TaobaocloudpushpushAPIRequest) SetAndroidOpenType(_androidOpenType string) error {
	r._androidOpenType = _androidOpenType
	r.Set("android_open_type", _androidOpenType)
	return nil
}

// GetAndroidOpenType AndroidOpenType Getter
func (r TaobaocloudpushpushAPIRequest) GetAndroidOpenType() string {
	return r._androidOpenType
}

// SetAndroidOpenUrl is AndroidOpenUrl Setter
// Android收到推送后打开对应的url,仅仅当androidOpenType=3有效
func (r *TaobaocloudpushpushAPIRequest) SetAndroidOpenUrl(_androidOpenUrl string) error {
	r._androidOpenUrl = _androidOpenUrl
	r.Set("android_open_url", _androidOpenUrl)
	return nil
}

// GetAndroidOpenUrl AndroidOpenUrl Getter
func (r TaobaocloudpushpushAPIRequest) GetAndroidOpenUrl() string {
	return r._androidOpenUrl
}

// SetBatchNumber is BatchNumber Setter
// 批次编号,用于活动效果统计
func (r *TaobaocloudpushpushAPIRequest) SetBatchNumber(_batchNumber string) error {
	r._batchNumber = _batchNumber
	r.Set("batch_number", _batchNumber)
	return nil
}

// GetBatchNumber BatchNumber Getter
func (r TaobaocloudpushpushAPIRequest) GetBatchNumber() string {
	return r._batchNumber
}

// SetBody is Body Setter
// 推送内容
func (r *TaobaocloudpushpushAPIRequest) SetBody(_body string) error {
	r._body = _body
	r.Set("body", _body)
	return nil
}

// GetBody Body Getter
func (r TaobaocloudpushpushAPIRequest) GetBody() string {
	return r._body
}

// SetIosBadge is IosBadge Setter
// iOS应用图标右上角角标
func (r *TaobaocloudpushpushAPIRequest) SetIosBadge(_iosBadge string) error {
	r._iosBadge = _iosBadge
	r.Set("ios_badge", _iosBadge)
	return nil
}

// GetIosBadge IosBadge Getter
func (r TaobaocloudpushpushAPIRequest) GetIosBadge() string {
	return r._iosBadge
}

// SetIosExtParameters is IosExtParameters Setter
// 自定义的kv结构,开发者扩展用 针对iOS设备
func (r *TaobaocloudpushpushAPIRequest) SetIosExtParameters(_iosExtParameters string) error {
	r._iosExtParameters = _iosExtParameters
	r.Set("ios_ext_parameters", _iosExtParameters)
	return nil
}

// GetIosExtParameters IosExtParameters Getter
func (r TaobaocloudpushpushAPIRequest) GetIosExtParameters() string {
	return r._iosExtParameters
}

// SetIosMusic is IosMusic Setter
// iOS通知声音
func (r *TaobaocloudpushpushAPIRequest) SetIosMusic(_iosMusic string) error {
	r._iosMusic = _iosMusic
	r.Set("ios_music", _iosMusic)
	return nil
}

// GetIosMusic IosMusic Getter
func (r TaobaocloudpushpushAPIRequest) GetIosMusic() string {
	return r._iosMusic
}

// SetSummery is Summery Setter
// 通知的摘要
func (r *TaobaocloudpushpushAPIRequest) SetSummery(_summery string) error {
	r._summery = _summery
	r.Set("summery", _summery)
	return nil
}

// GetSummery Summery Getter
func (r TaobaocloudpushpushAPIRequest) GetSummery() string {
	return r._summery
}

// SetTitle is Title Setter
// 推送的标题内容.
func (r *TaobaocloudpushpushAPIRequest) SetTitle(_title string) error {
	r._title = _title
	r.Set("title", _title)
	return nil
}

// GetTitle Title Getter
func (r TaobaocloudpushpushAPIRequest) GetTitle() string {
	return r._title
}

// SetAntiHarassDuration is AntiHarassDuration Setter
// 防打扰时长,取值范围为1~23
func (r *TaobaocloudpushpushAPIRequest) SetAntiHarassDuration(_antiHarassDuration int64) error {
	r._antiHarassDuration = _antiHarassDuration
	r.Set("anti_harass_duration", _antiHarassDuration)
	return nil
}

// GetAntiHarassDuration AntiHarassDuration Getter
func (r TaobaocloudpushpushAPIRequest) GetAntiHarassDuration() int64 {
	return r._antiHarassDuration
}

// SetAntiHarassStartTime is AntiHarassStartTime Setter
// 防打扰开始时间点,取值范围为0~23
func (r *TaobaocloudpushpushAPIRequest) SetAntiHarassStartTime(_antiHarassStartTime int64) error {
	r._antiHarassStartTime = _antiHarassStartTime
	r.Set("anti_harass_start_time", _antiHarassStartTime)
	return nil
}

// GetAntiHarassStartTime AntiHarassStartTime Getter
func (r TaobaocloudpushpushAPIRequest) GetAntiHarassStartTime() int64 {
	return r._antiHarassStartTime
}

// SetDeviceType is DeviceType Setter
// 设备类型,取值范围为:0~3云推送支持多种设备,各种设备类型编号如下:    iOS设备:deviceType=0; Andriod设备:deviceType=1;如果存在此字段,则向指定的设备类型推送消息。 默认为全部(3);
func (r *TaobaocloudpushpushAPIRequest) SetDeviceType(_deviceType int64) error {
	r._deviceType = _deviceType
	r.Set("device_type", _deviceType)
	return nil
}

// GetDeviceType DeviceType Getter
func (r TaobaocloudpushpushAPIRequest) GetDeviceType() int64 {
	return r._deviceType
}

// SetTimeout is Timeout Setter
// 离线消息保存时长,取值范围为1~72,若不填,则表示不保存离线消息
func (r *TaobaocloudpushpushAPIRequest) SetTimeout(_timeout int64) error {
	r._timeout = _timeout
	r.Set("timeout", _timeout)
	return nil
}

// GetTimeout Timeout Getter
func (r TaobaocloudpushpushAPIRequest) GetTimeout() int64 {
	return r._timeout
}

// SetType is Type Setter
// 0:表示消息(默认为0),1:表示通知
func (r *TaobaocloudpushpushAPIRequest) SetType(_type int64) error {
	r._type = _type
	r.Set("type", _type)
	return nil
}

// GetType Type Getter
func (r TaobaocloudpushpushAPIRequest) GetType() int64 {
	return r._type
}

// SetRemind is Remind Setter
// 当APP不在线时候，是否通过通知提醒.  针对不同设备，处理逻辑不同。 该参数只针对iOS设备生效， (remind=true  &amp; 发送消息的话(type=0)). 当你的目标设备不在线(既长连接通道不通, 我们会将这条消息的标题，通过苹果的apns通道再送达一次。发apns是发送生产环境的apns，需要在云推送配置的app的iOS生产证书和密码需要正确，否则也发送不了。 (remind=false &amp; 并且是发送消息的话(type=0))，那么设备不在线，则不会再走苹果apns发送了。
func (r *TaobaocloudpushpushAPIRequest) SetRemind(_remind bool) error {
	r._remind = _remind
	r.Set("remind", _remind)
	return nil
}

// GetRemind Remind Getter
func (r TaobaocloudpushpushAPIRequest) GetRemind() bool {
	return r._remind
}

// SetStoreOffline is StoreOffline Setter
// 离线消息是否保存,若保存, 在推送时候，用户即使不在线，下一次上线则会收到
func (r *TaobaocloudpushpushAPIRequest) SetStoreOffline(_storeOffline bool) error {
	r._storeOffline = _storeOffline
	r.Set("store_offline", _storeOffline)
	return nil
}

// GetStoreOffline StoreOffline Getter
func (r TaobaocloudpushpushAPIRequest) GetStoreOffline() bool {
	return r._storeOffline
}
