package aliyun

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// PushAliyuncsComPushNotification2015_03_18APIRequest 推送通知 API请求
// push.aliyuncs.com.pushNotification.2015-03-18
//
// pushNotification
type PushAliyuncsComPushNotification2015_03_18APIRequest struct {
	model.Params
	// 用户账号列表,以换行区分,仅sendType为3时有效
	_account string
	// 自定义的kv结构,开发者扩展用
	_androidExtraMap string
	// 通知声音
	_androidMusic string
	// 通知类型 1:震动 2:响铃
	_androidNotifyType int64
	// 打开app指定位置
	_androidOpenActivity string
	// 点击通知后动作
	_androidOpenType int64
	// 打开应用,网页
	_androidOpenUrl string
	// 防打扰时长,取值范围为1~23
	_antiHarassDuration int64
	// 防打扰开始时间点,取值范围为0~23
	_antiHarassStartTime int64
	// 应用标识
	_appId int64
	// 批次编号,用于活动效果统计
	_batchNumber string
	// 设备编号列表,以换行区分,仅sendType为4时有效
	_deviceId string
	// 设备类型,取值范围为:0~3云推送支持多种设备,各 种设备类型编号如下:IOS设备:deviceType&amp;1=1; Andriod设备:deviceType&amp;2=2;如果存在此字段,则向 指定的设备类型推送消息。默认为全部(3);
	_deviceType int64
	// 自定义的kv结构,开发者扩展用
	_iosExtraMap string
	// 角标
	_iosFooter int64
	// 默认音乐
	_iosMusic string
	// 推送时间,若空表示立即推送,推送时间不能早于当前时间
	_pushTime string
	// 推送类型,取值范围:1~4; 1:所有人,无需指定tag、 deviceType等2:一群人,必须指定tag3:指定用户,根据 用户账号列表文件发送消息4:指定设备,根据设备编码列 表文件发送消息默认值为1
	_sendType int64
	// 摘要
	_summary string
	// 标签名称,仅支持1个标签,仅sendType为2时有效
	_tag string
	// 离线消息保存时长,取值范围为1~72,若不填,则表示不保存离线消息
	_timeout int64
	// 标题
	_title string
}

// NewPushAliyuncsComPushNotification2015_03_18Request 初始化PushAliyuncsComPushNotification2015_03_18APIRequest对象
func NewPushAliyuncsComPushNotification2015_03_18Request() *PushAliyuncsComPushNotification2015_03_18APIRequest {
	return &PushAliyuncsComPushNotification2015_03_18APIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r PushAliyuncsComPushNotification2015_03_18APIRequest) GetApiMethodName() string {
	return "push.aliyuncs.com.pushNotification.2015-03-18"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r PushAliyuncsComPushNotification2015_03_18APIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetAccount is Account Setter
// 用户账号列表,以换行区分,仅sendType为3时有效
func (r *PushAliyuncsComPushNotification2015_03_18APIRequest) SetAccount(_account string) error {
	r._account = _account
	r.Set("Account", _account)
	return nil
}

// GetAccount Account Getter
func (r PushAliyuncsComPushNotification2015_03_18APIRequest) GetAccount() string {
	return r._account
}

// SetAndroidExtraMap is AndroidExtraMap Setter
// 自定义的kv结构,开发者扩展用
func (r *PushAliyuncsComPushNotification2015_03_18APIRequest) SetAndroidExtraMap(_androidExtraMap string) error {
	r._androidExtraMap = _androidExtraMap
	r.Set("AndroidExtraMap", _androidExtraMap)
	return nil
}

// GetAndroidExtraMap AndroidExtraMap Getter
func (r PushAliyuncsComPushNotification2015_03_18APIRequest) GetAndroidExtraMap() string {
	return r._androidExtraMap
}

// SetAndroidMusic is AndroidMusic Setter
// 通知声音
func (r *PushAliyuncsComPushNotification2015_03_18APIRequest) SetAndroidMusic(_androidMusic string) error {
	r._androidMusic = _androidMusic
	r.Set("AndroidMusic", _androidMusic)
	return nil
}

// GetAndroidMusic AndroidMusic Getter
func (r PushAliyuncsComPushNotification2015_03_18APIRequest) GetAndroidMusic() string {
	return r._androidMusic
}

// SetAndroidNotifyType is AndroidNotifyType Setter
// 通知类型 1:震动 2:响铃
func (r *PushAliyuncsComPushNotification2015_03_18APIRequest) SetAndroidNotifyType(_androidNotifyType int64) error {
	r._androidNotifyType = _androidNotifyType
	r.Set("AndroidNotifyType", _androidNotifyType)
	return nil
}

// GetAndroidNotifyType AndroidNotifyType Getter
func (r PushAliyuncsComPushNotification2015_03_18APIRequest) GetAndroidNotifyType() int64 {
	return r._androidNotifyType
}

// SetAndroidOpenActivity is AndroidOpenActivity Setter
// 打开app指定位置
func (r *PushAliyuncsComPushNotification2015_03_18APIRequest) SetAndroidOpenActivity(_androidOpenActivity string) error {
	r._androidOpenActivity = _androidOpenActivity
	r.Set("AndroidOpenActivity", _androidOpenActivity)
	return nil
}

// GetAndroidOpenActivity AndroidOpenActivity Getter
func (r PushAliyuncsComPushNotification2015_03_18APIRequest) GetAndroidOpenActivity() string {
	return r._androidOpenActivity
}

// SetAndroidOpenType is AndroidOpenType Setter
// 点击通知后动作
func (r *PushAliyuncsComPushNotification2015_03_18APIRequest) SetAndroidOpenType(_androidOpenType int64) error {
	r._androidOpenType = _androidOpenType
	r.Set("AndroidOpenType", _androidOpenType)
	return nil
}

// GetAndroidOpenType AndroidOpenType Getter
func (r PushAliyuncsComPushNotification2015_03_18APIRequest) GetAndroidOpenType() int64 {
	return r._androidOpenType
}

// SetAndroidOpenUrl is AndroidOpenUrl Setter
// 打开应用,网页
func (r *PushAliyuncsComPushNotification2015_03_18APIRequest) SetAndroidOpenUrl(_androidOpenUrl string) error {
	r._androidOpenUrl = _androidOpenUrl
	r.Set("AndroidOpenUrl", _androidOpenUrl)
	return nil
}

// GetAndroidOpenUrl AndroidOpenUrl Getter
func (r PushAliyuncsComPushNotification2015_03_18APIRequest) GetAndroidOpenUrl() string {
	return r._androidOpenUrl
}

// SetAntiHarassDuration is AntiHarassDuration Setter
// 防打扰时长,取值范围为1~23
func (r *PushAliyuncsComPushNotification2015_03_18APIRequest) SetAntiHarassDuration(_antiHarassDuration int64) error {
	r._antiHarassDuration = _antiHarassDuration
	r.Set("AntiHarassDuration", _antiHarassDuration)
	return nil
}

// GetAntiHarassDuration AntiHarassDuration Getter
func (r PushAliyuncsComPushNotification2015_03_18APIRequest) GetAntiHarassDuration() int64 {
	return r._antiHarassDuration
}

// SetAntiHarassStartTime is AntiHarassStartTime Setter
// 防打扰开始时间点,取值范围为0~23
func (r *PushAliyuncsComPushNotification2015_03_18APIRequest) SetAntiHarassStartTime(_antiHarassStartTime int64) error {
	r._antiHarassStartTime = _antiHarassStartTime
	r.Set("AntiHarassStartTime", _antiHarassStartTime)
	return nil
}

// GetAntiHarassStartTime AntiHarassStartTime Getter
func (r PushAliyuncsComPushNotification2015_03_18APIRequest) GetAntiHarassStartTime() int64 {
	return r._antiHarassStartTime
}

// SetAppId is AppId Setter
// 应用标识
func (r *PushAliyuncsComPushNotification2015_03_18APIRequest) SetAppId(_appId int64) error {
	r._appId = _appId
	r.Set("AppId", _appId)
	return nil
}

// GetAppId AppId Getter
func (r PushAliyuncsComPushNotification2015_03_18APIRequest) GetAppId() int64 {
	return r._appId
}

// SetBatchNumber is BatchNumber Setter
// 批次编号,用于活动效果统计
func (r *PushAliyuncsComPushNotification2015_03_18APIRequest) SetBatchNumber(_batchNumber string) error {
	r._batchNumber = _batchNumber
	r.Set("BatchNumber", _batchNumber)
	return nil
}

// GetBatchNumber BatchNumber Getter
func (r PushAliyuncsComPushNotification2015_03_18APIRequest) GetBatchNumber() string {
	return r._batchNumber
}

// SetDeviceId is DeviceId Setter
// 设备编号列表,以换行区分,仅sendType为4时有效
func (r *PushAliyuncsComPushNotification2015_03_18APIRequest) SetDeviceId(_deviceId string) error {
	r._deviceId = _deviceId
	r.Set("DeviceId", _deviceId)
	return nil
}

// GetDeviceId DeviceId Getter
func (r PushAliyuncsComPushNotification2015_03_18APIRequest) GetDeviceId() string {
	return r._deviceId
}

// SetDeviceType is DeviceType Setter
// 设备类型,取值范围为:0~3云推送支持多种设备,各 种设备类型编号如下:IOS设备:deviceType&amp;1=1; Andriod设备:deviceType&amp;2=2;如果存在此字段,则向 指定的设备类型推送消息。默认为全部(3);
func (r *PushAliyuncsComPushNotification2015_03_18APIRequest) SetDeviceType(_deviceType int64) error {
	r._deviceType = _deviceType
	r.Set("DeviceType", _deviceType)
	return nil
}

// GetDeviceType DeviceType Getter
func (r PushAliyuncsComPushNotification2015_03_18APIRequest) GetDeviceType() int64 {
	return r._deviceType
}

// SetIosExtraMap is IosExtraMap Setter
// 自定义的kv结构,开发者扩展用
func (r *PushAliyuncsComPushNotification2015_03_18APIRequest) SetIosExtraMap(_iosExtraMap string) error {
	r._iosExtraMap = _iosExtraMap
	r.Set("IosExtraMap", _iosExtraMap)
	return nil
}

// GetIosExtraMap IosExtraMap Getter
func (r PushAliyuncsComPushNotification2015_03_18APIRequest) GetIosExtraMap() string {
	return r._iosExtraMap
}

// SetIosFooter is IosFooter Setter
// 角标
func (r *PushAliyuncsComPushNotification2015_03_18APIRequest) SetIosFooter(_iosFooter int64) error {
	r._iosFooter = _iosFooter
	r.Set("IosFooter", _iosFooter)
	return nil
}

// GetIosFooter IosFooter Getter
func (r PushAliyuncsComPushNotification2015_03_18APIRequest) GetIosFooter() int64 {
	return r._iosFooter
}

// SetIosMusic is IosMusic Setter
// 默认音乐
func (r *PushAliyuncsComPushNotification2015_03_18APIRequest) SetIosMusic(_iosMusic string) error {
	r._iosMusic = _iosMusic
	r.Set("IosMusic", _iosMusic)
	return nil
}

// GetIosMusic IosMusic Getter
func (r PushAliyuncsComPushNotification2015_03_18APIRequest) GetIosMusic() string {
	return r._iosMusic
}

// SetPushTime is PushTime Setter
// 推送时间,若空表示立即推送,推送时间不能早于当前时间
func (r *PushAliyuncsComPushNotification2015_03_18APIRequest) SetPushTime(_pushTime string) error {
	r._pushTime = _pushTime
	r.Set("PushTime", _pushTime)
	return nil
}

// GetPushTime PushTime Getter
func (r PushAliyuncsComPushNotification2015_03_18APIRequest) GetPushTime() string {
	return r._pushTime
}

// SetSendType is SendType Setter
// 推送类型,取值范围:1~4; 1:所有人,无需指定tag、 deviceType等2:一群人,必须指定tag3:指定用户,根据 用户账号列表文件发送消息4:指定设备,根据设备编码列 表文件发送消息默认值为1
func (r *PushAliyuncsComPushNotification2015_03_18APIRequest) SetSendType(_sendType int64) error {
	r._sendType = _sendType
	r.Set("SendType", _sendType)
	return nil
}

// GetSendType SendType Getter
func (r PushAliyuncsComPushNotification2015_03_18APIRequest) GetSendType() int64 {
	return r._sendType
}

// SetSummary is Summary Setter
// 摘要
func (r *PushAliyuncsComPushNotification2015_03_18APIRequest) SetSummary(_summary string) error {
	r._summary = _summary
	r.Set("Summary", _summary)
	return nil
}

// GetSummary Summary Getter
func (r PushAliyuncsComPushNotification2015_03_18APIRequest) GetSummary() string {
	return r._summary
}

// SetTag is Tag Setter
// 标签名称,仅支持1个标签,仅sendType为2时有效
func (r *PushAliyuncsComPushNotification2015_03_18APIRequest) SetTag(_tag string) error {
	r._tag = _tag
	r.Set("Tag", _tag)
	return nil
}

// GetTag Tag Getter
func (r PushAliyuncsComPushNotification2015_03_18APIRequest) GetTag() string {
	return r._tag
}

// SetTimeout is Timeout Setter
// 离线消息保存时长,取值范围为1~72,若不填,则表示不保存离线消息
func (r *PushAliyuncsComPushNotification2015_03_18APIRequest) SetTimeout(_timeout int64) error {
	r._timeout = _timeout
	r.Set("Timeout", _timeout)
	return nil
}

// GetTimeout Timeout Getter
func (r PushAliyuncsComPushNotification2015_03_18APIRequest) GetTimeout() int64 {
	return r._timeout
}

// SetTitle is Title Setter
// 标题
func (r *PushAliyuncsComPushNotification2015_03_18APIRequest) SetTitle(_title string) error {
	r._title = _title
	r.Set("Title", _title)
	return nil
}

// GetTitle Title Getter
func (r PushAliyuncsComPushNotification2015_03_18APIRequest) GetTitle() string {
	return r._title
}
