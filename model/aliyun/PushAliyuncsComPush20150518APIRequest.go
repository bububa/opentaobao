package aliyun

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// PushAliyuncsComPush20150518APIRequest 云推送指令API API请求
// push.aliyuncs.com.push.20150518
//
// 阿里云推送新增API，允许一条推送指令同时发布到多个终端上。
type PushAliyuncsComPush20150518APIRequest struct {
	model.Params
	// 用户账号列表,以换行区分,仅sendType为3时有效
	_account string
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
	// 推送接收设备，多个以逗号分隔
	_deviceId string
	// iOS应用图标右上角角标
	_iOSBadge string
	// 自定义的kv结构,开发者扩展用 针对iOS设备
	_iOSExtParameters string
	// iOS通知声音
	_iOSMusic string
	// 推送时间,若空表示立即推送,推送时间不能早于当前时间
	_pushTime string
	// 通知的摘要
	_summery string
	// 推送的标题内容.
	_title string
	// 防打扰时长,取值范围为1~23
	_antiHarassDuration int64
	// 防打扰开始时间点,取值范围为0~23
	_antiHarassStartTime int64
	// 应用标识
	_appId int64
	// 设备类型,取值范围为:0~3云推送支持多种设备,各 种设备类型编号如下: 0:IOS设备; 1:Andriod设备 3:全部. 默认为3.
	_deviceType int64
	// 推送类型,取值范围:1~4; 1:所有人,无需指定tag、 deviceType等2:一群人,必须指定tag3:指定用户,根据 用户账号列表文件发送消息4:指定设备,根据设备编码列 表文件发送消息默认值为1
	_sendType int64
	// 离线消息保存时长,取值范围为1~72,若不填,则表示不保存离线消息
	_timeout int64
	// 当APP不在线时候，是否通过通知提醒
	_remind bool
	// 离线消息是否保存,若保存, 在推送时候，用户即使不在线，下一次上线则会收到
	_storeOffline bool
}

// NewPushAliyuncsComPush20150518Request 初始化PushAliyuncsComPush20150518APIRequest对象
func NewPushAliyuncsComPush20150518Request() *PushAliyuncsComPush20150518APIRequest {
	return &PushAliyuncsComPush20150518APIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r PushAliyuncsComPush20150518APIRequest) GetApiMethodName() string {
	return "push.aliyuncs.com.push.20150518"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r PushAliyuncsComPush20150518APIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetAccount is Account Setter
// 用户账号列表,以换行区分,仅sendType为3时有效
func (r *PushAliyuncsComPush20150518APIRequest) SetAccount(_account string) error {
	r._account = _account
	r.Set("Account", _account)
	return nil
}

// GetAccount Account Getter
func (r PushAliyuncsComPush20150518APIRequest) GetAccount() string {
	return r._account
}

// SetAndroidActivity is AndroidActivity Setter
// Android对应的activity,仅仅当androidOpenType=2有效
func (r *PushAliyuncsComPush20150518APIRequest) SetAndroidActivity(_androidActivity string) error {
	r._androidActivity = _androidActivity
	r.Set("AndroidActivity", _androidActivity)
	return nil
}

// GetAndroidActivity AndroidActivity Getter
func (r PushAliyuncsComPush20150518APIRequest) GetAndroidActivity() string {
	return r._androidActivity
}

// SetAndroidExtParameters is AndroidExtParameters Setter
// 自定义的kv结构,开发者扩展用 针对android
func (r *PushAliyuncsComPush20150518APIRequest) SetAndroidExtParameters(_androidExtParameters string) error {
	r._androidExtParameters = _androidExtParameters
	r.Set("AndroidExtParameters", _androidExtParameters)
	return nil
}

// GetAndroidExtParameters AndroidExtParameters Getter
func (r PushAliyuncsComPush20150518APIRequest) GetAndroidExtParameters() string {
	return r._androidExtParameters
}

// SetAndroidMusic is AndroidMusic Setter
// android通知声音
func (r *PushAliyuncsComPush20150518APIRequest) SetAndroidMusic(_androidMusic string) error {
	r._androidMusic = _androidMusic
	r.Set("AndroidMusic", _androidMusic)
	return nil
}

// GetAndroidMusic AndroidMusic Getter
func (r PushAliyuncsComPush20150518APIRequest) GetAndroidMusic() string {
	return r._androidMusic
}

// SetAndroidOpenType is AndroidOpenType Setter
// 点击通知后动作,1:打开应用 2: 打开应用Activity 3:打开 url
func (r *PushAliyuncsComPush20150518APIRequest) SetAndroidOpenType(_androidOpenType string) error {
	r._androidOpenType = _androidOpenType
	r.Set("AndroidOpenType", _androidOpenType)
	return nil
}

// GetAndroidOpenType AndroidOpenType Getter
func (r PushAliyuncsComPush20150518APIRequest) GetAndroidOpenType() string {
	return r._androidOpenType
}

// SetAndroidOpenUrl is AndroidOpenUrl Setter
// Android收到推送后打开对应的url,仅仅当androidOpenType=3有效
func (r *PushAliyuncsComPush20150518APIRequest) SetAndroidOpenUrl(_androidOpenUrl string) error {
	r._androidOpenUrl = _androidOpenUrl
	r.Set("AndroidOpenUrl", _androidOpenUrl)
	return nil
}

// GetAndroidOpenUrl AndroidOpenUrl Getter
func (r PushAliyuncsComPush20150518APIRequest) GetAndroidOpenUrl() string {
	return r._androidOpenUrl
}

// SetBatchNumber is BatchNumber Setter
// 批次编号,用于活动效果统计
func (r *PushAliyuncsComPush20150518APIRequest) SetBatchNumber(_batchNumber string) error {
	r._batchNumber = _batchNumber
	r.Set("BatchNumber", _batchNumber)
	return nil
}

// GetBatchNumber BatchNumber Getter
func (r PushAliyuncsComPush20150518APIRequest) GetBatchNumber() string {
	return r._batchNumber
}

// SetBody is Body Setter
// 推送内容
func (r *PushAliyuncsComPush20150518APIRequest) SetBody(_body string) error {
	r._body = _body
	r.Set("Body", _body)
	return nil
}

// GetBody Body Getter
func (r PushAliyuncsComPush20150518APIRequest) GetBody() string {
	return r._body
}

// SetDeviceId is DeviceId Setter
// 推送接收设备，多个以逗号分隔
func (r *PushAliyuncsComPush20150518APIRequest) SetDeviceId(_deviceId string) error {
	r._deviceId = _deviceId
	r.Set("DeviceId", _deviceId)
	return nil
}

// GetDeviceId DeviceId Getter
func (r PushAliyuncsComPush20150518APIRequest) GetDeviceId() string {
	return r._deviceId
}

// SetIOSBadge is IOSBadge Setter
// iOS应用图标右上角角标
func (r *PushAliyuncsComPush20150518APIRequest) SetIOSBadge(_iOSBadge string) error {
	r._iOSBadge = _iOSBadge
	r.Set("IOSBadge", _iOSBadge)
	return nil
}

// GetIOSBadge IOSBadge Getter
func (r PushAliyuncsComPush20150518APIRequest) GetIOSBadge() string {
	return r._iOSBadge
}

// SetIOSExtParameters is IOSExtParameters Setter
// 自定义的kv结构,开发者扩展用 针对iOS设备
func (r *PushAliyuncsComPush20150518APIRequest) SetIOSExtParameters(_iOSExtParameters string) error {
	r._iOSExtParameters = _iOSExtParameters
	r.Set("IOSExtParameters", _iOSExtParameters)
	return nil
}

// GetIOSExtParameters IOSExtParameters Getter
func (r PushAliyuncsComPush20150518APIRequest) GetIOSExtParameters() string {
	return r._iOSExtParameters
}

// SetIOSMusic is IOSMusic Setter
// iOS通知声音
func (r *PushAliyuncsComPush20150518APIRequest) SetIOSMusic(_iOSMusic string) error {
	r._iOSMusic = _iOSMusic
	r.Set("IOSMusic", _iOSMusic)
	return nil
}

// GetIOSMusic IOSMusic Getter
func (r PushAliyuncsComPush20150518APIRequest) GetIOSMusic() string {
	return r._iOSMusic
}

// SetPushTime is PushTime Setter
// 推送时间,若空表示立即推送,推送时间不能早于当前时间
func (r *PushAliyuncsComPush20150518APIRequest) SetPushTime(_pushTime string) error {
	r._pushTime = _pushTime
	r.Set("PushTime", _pushTime)
	return nil
}

// GetPushTime PushTime Getter
func (r PushAliyuncsComPush20150518APIRequest) GetPushTime() string {
	return r._pushTime
}

// SetSummery is Summery Setter
// 通知的摘要
func (r *PushAliyuncsComPush20150518APIRequest) SetSummery(_summery string) error {
	r._summery = _summery
	r.Set("Summery", _summery)
	return nil
}

// GetSummery Summery Getter
func (r PushAliyuncsComPush20150518APIRequest) GetSummery() string {
	return r._summery
}

// SetTitle is Title Setter
// 推送的标题内容.
func (r *PushAliyuncsComPush20150518APIRequest) SetTitle(_title string) error {
	r._title = _title
	r.Set("Title", _title)
	return nil
}

// GetTitle Title Getter
func (r PushAliyuncsComPush20150518APIRequest) GetTitle() string {
	return r._title
}

// SetAntiHarassDuration is AntiHarassDuration Setter
// 防打扰时长,取值范围为1~23
func (r *PushAliyuncsComPush20150518APIRequest) SetAntiHarassDuration(_antiHarassDuration int64) error {
	r._antiHarassDuration = _antiHarassDuration
	r.Set("AntiHarassDuration", _antiHarassDuration)
	return nil
}

// GetAntiHarassDuration AntiHarassDuration Getter
func (r PushAliyuncsComPush20150518APIRequest) GetAntiHarassDuration() int64 {
	return r._antiHarassDuration
}

// SetAntiHarassStartTime is AntiHarassStartTime Setter
// 防打扰开始时间点,取值范围为0~23
func (r *PushAliyuncsComPush20150518APIRequest) SetAntiHarassStartTime(_antiHarassStartTime int64) error {
	r._antiHarassStartTime = _antiHarassStartTime
	r.Set("AntiHarassStartTime", _antiHarassStartTime)
	return nil
}

// GetAntiHarassStartTime AntiHarassStartTime Getter
func (r PushAliyuncsComPush20150518APIRequest) GetAntiHarassStartTime() int64 {
	return r._antiHarassStartTime
}

// SetAppId is AppId Setter
// 应用标识
func (r *PushAliyuncsComPush20150518APIRequest) SetAppId(_appId int64) error {
	r._appId = _appId
	r.Set("AppId", _appId)
	return nil
}

// GetAppId AppId Getter
func (r PushAliyuncsComPush20150518APIRequest) GetAppId() int64 {
	return r._appId
}

// SetDeviceType is DeviceType Setter
// 设备类型,取值范围为:0~3云推送支持多种设备,各 种设备类型编号如下: 0:IOS设备; 1:Andriod设备 3:全部. 默认为3.
func (r *PushAliyuncsComPush20150518APIRequest) SetDeviceType(_deviceType int64) error {
	r._deviceType = _deviceType
	r.Set("DeviceType", _deviceType)
	return nil
}

// GetDeviceType DeviceType Getter
func (r PushAliyuncsComPush20150518APIRequest) GetDeviceType() int64 {
	return r._deviceType
}

// SetSendType is SendType Setter
// 推送类型,取值范围:1~4; 1:所有人,无需指定tag、 deviceType等2:一群人,必须指定tag3:指定用户,根据 用户账号列表文件发送消息4:指定设备,根据设备编码列 表文件发送消息默认值为1
func (r *PushAliyuncsComPush20150518APIRequest) SetSendType(_sendType int64) error {
	r._sendType = _sendType
	r.Set("SendType", _sendType)
	return nil
}

// GetSendType SendType Getter
func (r PushAliyuncsComPush20150518APIRequest) GetSendType() int64 {
	return r._sendType
}

// SetTimeout is Timeout Setter
// 离线消息保存时长,取值范围为1~72,若不填,则表示不保存离线消息
func (r *PushAliyuncsComPush20150518APIRequest) SetTimeout(_timeout int64) error {
	r._timeout = _timeout
	r.Set("Timeout", _timeout)
	return nil
}

// GetTimeout Timeout Getter
func (r PushAliyuncsComPush20150518APIRequest) GetTimeout() int64 {
	return r._timeout
}

// SetRemind is Remind Setter
// 当APP不在线时候，是否通过通知提醒
func (r *PushAliyuncsComPush20150518APIRequest) SetRemind(_remind bool) error {
	r._remind = _remind
	r.Set("Remind", _remind)
	return nil
}

// GetRemind Remind Getter
func (r PushAliyuncsComPush20150518APIRequest) GetRemind() bool {
	return r._remind
}

// SetStoreOffline is StoreOffline Setter
// 离线消息是否保存,若保存, 在推送时候，用户即使不在线，下一次上线则会收到
func (r *PushAliyuncsComPush20150518APIRequest) SetStoreOffline(_storeOffline bool) error {
	r._storeOffline = _storeOffline
	r.Set("StoreOffline", _storeOffline)
	return nil
}

// GetStoreOffline StoreOffline Getter
func (r PushAliyuncsComPush20150518APIRequest) GetStoreOffline() bool {
	return r._storeOffline
}
