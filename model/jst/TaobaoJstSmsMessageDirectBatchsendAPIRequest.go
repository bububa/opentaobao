package jst

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaojstsmsmessagedirectbatchsendAPIRequest OAID批量发送，支持明文手机号发送 API请求
// taobao.jst.sms.message.direct.batchsend
//
// OAID批量发送，支持明文手机号发送
type TaobaojstsmsmessagedirectbatchsendAPIRequest struct {
	model.Params
	// 短信签名（如果是通过OAID发送短信则签名需要是数组格式，数组长度需要和OAID数量保持一致。）
	_signName string
	// 废弃字段
	_url string
	// 短信模版Code（明文发送短信和OAID发送均不要传数组格式）
	_smsTemplateCode string
	// 短信接收号码，json格式，最多200个号码
	_recNum string
	// 模板参数替换方式："[{\"msg\":\"hello1\",\"date\":\"2021-12-03\"},{\"msg\":\"hello2\",\"date\":\"2021-12-04\"},{\"msg\":\"hello3\",\"date\":\"2021-12-05\"}]"
	_smsContent string
	// 短信扩展码（JSON字符串数组格式），拓展码个数需要和电话号码个数一致。
	_extendNum string
	// 废弃字段
	_taskCode string
	// 对于taskSign相同的请求平台认为是商家的同一次短信发送，可用于对OAID的明文号码去重。
	_taskSign string
	// OAID批量发短信的入参，传该参数的时候rec_num不需要传，最大50个。
	_oaids string
	// OAID批量发短信时必传，为OAID一一对应的订单ID。
	_orderIds string
	// 如果传，必须是一个JSON对象格式的字符串。
	_extraData string
}

// NewTaobaojstsmsmessagedirectbatchsendRequest 初始化TaobaojstsmsmessagedirectbatchsendAPIRequest对象
func NewTaobaojstsmsmessagedirectbatchsendRequest() *TaobaojstsmsmessagedirectbatchsendAPIRequest {
	return &TaobaojstsmsmessagedirectbatchsendAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaojstsmsmessagedirectbatchsendAPIRequest) GetApiMethodName() string {
	return "taobao.jst.sms.message.direct.batchsend"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaojstsmsmessagedirectbatchsendAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaojstsmsmessagedirectbatchsendAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSignName is SignName Setter
// 短信签名（如果是通过OAID发送短信则签名需要是数组格式，数组长度需要和OAID数量保持一致。）
func (r *TaobaojstsmsmessagedirectbatchsendAPIRequest) SetSignName(_signName string) error {
	r._signName = _signName
	r.Set("sign_name", _signName)
	return nil
}

// GetSignName SignName Getter
func (r TaobaojstsmsmessagedirectbatchsendAPIRequest) GetSignName() string {
	return r._signName
}

// SetUrl is Url Setter
// 废弃字段
func (r *TaobaojstsmsmessagedirectbatchsendAPIRequest) SetUrl(_url string) error {
	r._url = _url
	r.Set("url", _url)
	return nil
}

// GetUrl Url Getter
func (r TaobaojstsmsmessagedirectbatchsendAPIRequest) GetUrl() string {
	return r._url
}

// SetSmsTemplateCode is SmsTemplateCode Setter
// 短信模版Code（明文发送短信和OAID发送均不要传数组格式）
func (r *TaobaojstsmsmessagedirectbatchsendAPIRequest) SetSmsTemplateCode(_smsTemplateCode string) error {
	r._smsTemplateCode = _smsTemplateCode
	r.Set("sms_template_code", _smsTemplateCode)
	return nil
}

// GetSmsTemplateCode SmsTemplateCode Getter
func (r TaobaojstsmsmessagedirectbatchsendAPIRequest) GetSmsTemplateCode() string {
	return r._smsTemplateCode
}

// SetRecNum is RecNum Setter
// 短信接收号码，json格式，最多200个号码
func (r *TaobaojstsmsmessagedirectbatchsendAPIRequest) SetRecNum(_recNum string) error {
	r._recNum = _recNum
	r.Set("rec_num", _recNum)
	return nil
}

// GetRecNum RecNum Getter
func (r TaobaojstsmsmessagedirectbatchsendAPIRequest) GetRecNum() string {
	return r._recNum
}

// SetSmsContent is SmsContent Setter
// 模板参数替换方式：&#34;[{\&#34;msg\&#34;:\&#34;hello1\&#34;,\&#34;date\&#34;:\&#34;2021-12-03\&#34;},{\&#34;msg\&#34;:\&#34;hello2\&#34;,\&#34;date\&#34;:\&#34;2021-12-04\&#34;},{\&#34;msg\&#34;:\&#34;hello3\&#34;,\&#34;date\&#34;:\&#34;2021-12-05\&#34;}]&#34;
func (r *TaobaojstsmsmessagedirectbatchsendAPIRequest) SetSmsContent(_smsContent string) error {
	r._smsContent = _smsContent
	r.Set("sms_content", _smsContent)
	return nil
}

// GetSmsContent SmsContent Getter
func (r TaobaojstsmsmessagedirectbatchsendAPIRequest) GetSmsContent() string {
	return r._smsContent
}

// SetExtendNum is ExtendNum Setter
// 短信扩展码（JSON字符串数组格式），拓展码个数需要和电话号码个数一致。
func (r *TaobaojstsmsmessagedirectbatchsendAPIRequest) SetExtendNum(_extendNum string) error {
	r._extendNum = _extendNum
	r.Set("extend_num", _extendNum)
	return nil
}

// GetExtendNum ExtendNum Getter
func (r TaobaojstsmsmessagedirectbatchsendAPIRequest) GetExtendNum() string {
	return r._extendNum
}

// SetTaskCode is TaskCode Setter
// 废弃字段
func (r *TaobaojstsmsmessagedirectbatchsendAPIRequest) SetTaskCode(_taskCode string) error {
	r._taskCode = _taskCode
	r.Set("task_code", _taskCode)
	return nil
}

// GetTaskCode TaskCode Getter
func (r TaobaojstsmsmessagedirectbatchsendAPIRequest) GetTaskCode() string {
	return r._taskCode
}

// SetTaskSign is TaskSign Setter
// 对于taskSign相同的请求平台认为是商家的同一次短信发送，可用于对OAID的明文号码去重。
func (r *TaobaojstsmsmessagedirectbatchsendAPIRequest) SetTaskSign(_taskSign string) error {
	r._taskSign = _taskSign
	r.Set("task_sign", _taskSign)
	return nil
}

// GetTaskSign TaskSign Getter
func (r TaobaojstsmsmessagedirectbatchsendAPIRequest) GetTaskSign() string {
	return r._taskSign
}

// SetOaids is Oaids Setter
// OAID批量发短信的入参，传该参数的时候rec_num不需要传，最大50个。
func (r *TaobaojstsmsmessagedirectbatchsendAPIRequest) SetOaids(_oaids string) error {
	r._oaids = _oaids
	r.Set("oaids", _oaids)
	return nil
}

// GetOaids Oaids Getter
func (r TaobaojstsmsmessagedirectbatchsendAPIRequest) GetOaids() string {
	return r._oaids
}

// SetOrderIds is OrderIds Setter
// OAID批量发短信时必传，为OAID一一对应的订单ID。
func (r *TaobaojstsmsmessagedirectbatchsendAPIRequest) SetOrderIds(_orderIds string) error {
	r._orderIds = _orderIds
	r.Set("order_ids", _orderIds)
	return nil
}

// GetOrderIds OrderIds Getter
func (r TaobaojstsmsmessagedirectbatchsendAPIRequest) GetOrderIds() string {
	return r._orderIds
}

// SetExtraData is ExtraData Setter
// 如果传，必须是一个JSON对象格式的字符串。
func (r *TaobaojstsmsmessagedirectbatchsendAPIRequest) SetExtraData(_extraData string) error {
	r._extraData = _extraData
	r.Set("extra_data", _extraData)
	return nil
}

// GetExtraData ExtraData Getter
func (r TaobaojstsmsmessagedirectbatchsendAPIRequest) GetExtraData() string {
	return r._extraData
}
