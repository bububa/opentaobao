package aliyun

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// PushaliyuncscompushMsg20150318APIRequest 消息推送 API请求
// push.aliyuncs.com.pushMsg.2015-03-18
//
// 消息推送  ,支持指定用户/账号/广播等模式
type PushaliyuncscompushMsg20150318APIRequest struct {
	model.Params
	// 用户账号列表,以换行区分,仅sendType为3时有效
	_account string
	// 批次编号,用于统计活动推送效果
	_batchNumber string
	// 消息体,UTF-8编码
	_body string
	// 设备编号列表,以换行区分,仅sendType为4时有效
	_deviceId string
	// 推送时间,若空表示立即推送,推送时间不能早于当前 时间
	_pushTime string
	// 标签名称,仅支持1个标签,仅sendType为2时有效
	_tag string
	// 标题
	_title string
	// 防打扰时长,取值范围为1~23
	_antiHarassDuration int64
	// 防打扰开始时间点,取值范围为0~23
	_antiHarassStartTime int64
	// 应用标识
	_appId int64
	// 设备类型,取值范围为:0~3云推送支持多种设备, 各种设备类型编号如下:IOS设备:deviceType&amp;1=1; Andriod设备:deviceType&amp;2=2;如果存在此字段,则 向指定的设备类型推送消息。默认为全部(3);
	_deviceType int64
	// 推送类型,取值范围:1~4; 1:所有人,无需指定tag、 deviceType等2:一群人,必须指定tag3:指定用户,根 据用户账号列表文件发送消息4:指定设备,根据设备编 码列表文件发送消息默认值为1
	_sendType int64
	// 离线消息保存时长,取值范围为1~72,若不填,则表 示不保存离线消息
	_timeout int64
}

// NewPushaliyuncscompushMsg20150318Request 初始化PushaliyuncscompushMsg20150318APIRequest对象
func NewPushaliyuncscompushMsg20150318Request() *PushaliyuncscompushMsg20150318APIRequest {
	return &PushaliyuncscompushMsg20150318APIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r PushaliyuncscompushMsg20150318APIRequest) GetApiMethodName() string {
	return "push.aliyuncs.com.pushMsg.2015-03-18"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r PushaliyuncscompushMsg20150318APIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r PushaliyuncscompushMsg20150318APIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAccount is Account Setter
// 用户账号列表,以换行区分,仅sendType为3时有效
func (r *PushaliyuncscompushMsg20150318APIRequest) SetAccount(_account string) error {
	r._account = _account
	r.Set("Account", _account)
	return nil
}

// GetAccount Account Getter
func (r PushaliyuncscompushMsg20150318APIRequest) GetAccount() string {
	return r._account
}

// SetBatchNumber is BatchNumber Setter
// 批次编号,用于统计活动推送效果
func (r *PushaliyuncscompushMsg20150318APIRequest) SetBatchNumber(_batchNumber string) error {
	r._batchNumber = _batchNumber
	r.Set("BatchNumber", _batchNumber)
	return nil
}

// GetBatchNumber BatchNumber Getter
func (r PushaliyuncscompushMsg20150318APIRequest) GetBatchNumber() string {
	return r._batchNumber
}

// SetBody is Body Setter
// 消息体,UTF-8编码
func (r *PushaliyuncscompushMsg20150318APIRequest) SetBody(_body string) error {
	r._body = _body
	r.Set("Body", _body)
	return nil
}

// GetBody Body Getter
func (r PushaliyuncscompushMsg20150318APIRequest) GetBody() string {
	return r._body
}

// SetDeviceId is DeviceId Setter
// 设备编号列表,以换行区分,仅sendType为4时有效
func (r *PushaliyuncscompushMsg20150318APIRequest) SetDeviceId(_deviceId string) error {
	r._deviceId = _deviceId
	r.Set("DeviceId", _deviceId)
	return nil
}

// GetDeviceId DeviceId Getter
func (r PushaliyuncscompushMsg20150318APIRequest) GetDeviceId() string {
	return r._deviceId
}

// SetPushTime is PushTime Setter
// 推送时间,若空表示立即推送,推送时间不能早于当前 时间
func (r *PushaliyuncscompushMsg20150318APIRequest) SetPushTime(_pushTime string) error {
	r._pushTime = _pushTime
	r.Set("PushTime", _pushTime)
	return nil
}

// GetPushTime PushTime Getter
func (r PushaliyuncscompushMsg20150318APIRequest) GetPushTime() string {
	return r._pushTime
}

// SetTag is Tag Setter
// 标签名称,仅支持1个标签,仅sendType为2时有效
func (r *PushaliyuncscompushMsg20150318APIRequest) SetTag(_tag string) error {
	r._tag = _tag
	r.Set("Tag", _tag)
	return nil
}

// GetTag Tag Getter
func (r PushaliyuncscompushMsg20150318APIRequest) GetTag() string {
	return r._tag
}

// SetTitle is Title Setter
// 标题
func (r *PushaliyuncscompushMsg20150318APIRequest) SetTitle(_title string) error {
	r._title = _title
	r.Set("Title", _title)
	return nil
}

// GetTitle Title Getter
func (r PushaliyuncscompushMsg20150318APIRequest) GetTitle() string {
	return r._title
}

// SetAntiHarassDuration is AntiHarassDuration Setter
// 防打扰时长,取值范围为1~23
func (r *PushaliyuncscompushMsg20150318APIRequest) SetAntiHarassDuration(_antiHarassDuration int64) error {
	r._antiHarassDuration = _antiHarassDuration
	r.Set("AntiHarassDuration", _antiHarassDuration)
	return nil
}

// GetAntiHarassDuration AntiHarassDuration Getter
func (r PushaliyuncscompushMsg20150318APIRequest) GetAntiHarassDuration() int64 {
	return r._antiHarassDuration
}

// SetAntiHarassStartTime is AntiHarassStartTime Setter
// 防打扰开始时间点,取值范围为0~23
func (r *PushaliyuncscompushMsg20150318APIRequest) SetAntiHarassStartTime(_antiHarassStartTime int64) error {
	r._antiHarassStartTime = _antiHarassStartTime
	r.Set("AntiHarassStartTime", _antiHarassStartTime)
	return nil
}

// GetAntiHarassStartTime AntiHarassStartTime Getter
func (r PushaliyuncscompushMsg20150318APIRequest) GetAntiHarassStartTime() int64 {
	return r._antiHarassStartTime
}

// SetAppId is AppId Setter
// 应用标识
func (r *PushaliyuncscompushMsg20150318APIRequest) SetAppId(_appId int64) error {
	r._appId = _appId
	r.Set("AppId", _appId)
	return nil
}

// GetAppId AppId Getter
func (r PushaliyuncscompushMsg20150318APIRequest) GetAppId() int64 {
	return r._appId
}

// SetDeviceType is DeviceType Setter
// 设备类型,取值范围为:0~3云推送支持多种设备, 各种设备类型编号如下:IOS设备:deviceType&amp;amp;1=1; Andriod设备:deviceType&amp;amp;2=2;如果存在此字段,则 向指定的设备类型推送消息。默认为全部(3);
func (r *PushaliyuncscompushMsg20150318APIRequest) SetDeviceType(_deviceType int64) error {
	r._deviceType = _deviceType
	r.Set("DeviceType", _deviceType)
	return nil
}

// GetDeviceType DeviceType Getter
func (r PushaliyuncscompushMsg20150318APIRequest) GetDeviceType() int64 {
	return r._deviceType
}

// SetSendType is SendType Setter
// 推送类型,取值范围:1~4; 1:所有人,无需指定tag、 deviceType等2:一群人,必须指定tag3:指定用户,根 据用户账号列表文件发送消息4:指定设备,根据设备编 码列表文件发送消息默认值为1
func (r *PushaliyuncscompushMsg20150318APIRequest) SetSendType(_sendType int64) error {
	r._sendType = _sendType
	r.Set("SendType", _sendType)
	return nil
}

// GetSendType SendType Getter
func (r PushaliyuncscompushMsg20150318APIRequest) GetSendType() int64 {
	return r._sendType
}

// SetTimeout is Timeout Setter
// 离线消息保存时长,取值范围为1~72,若不填,则表 示不保存离线消息
func (r *PushaliyuncscompushMsg20150318APIRequest) SetTimeout(_timeout int64) error {
	r._timeout = _timeout
	r.Set("Timeout", _timeout)
	return nil
}

// GetTimeout Timeout Getter
func (r PushaliyuncscompushMsg20150318APIRequest) GetTimeout() int64 {
	return r._timeout
}
