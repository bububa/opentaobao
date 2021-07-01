package jst

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoJstSmsMessageDirectBatchsendAPIRequest
聚石塔新短信发送接口 API请求
taobao.jst.sms.message.direct.batchsend

聚石塔所见即所得的短信发送接口 */
type TaobaoJstSmsMessageDirectBatchsendAPIRequest struct {
	model.Params
	// 短信签名
	_signName string
	// 短信中带的短链，如果不传则没有短信效果数据
	_url string
	// 短信模版CODE，必须为全变量模板
	_smsTemplateCode string
	// 短信接收号码，json格式，最多200个号码
	_recNum string
	// 短信内容，如果带${url}则会被入参url替换
	_smsContent string
	// 短信扩展码（JSON字符串数组格式），拓展码个数需要和电话号码个数一致。
	_extendNum string
	// 短信任务code，没有请先创建。
	_taskCode string
	// 一个在taskcode下唯一的随机字符串，对于taskSign相同的请求平台认为是商家的同一次短信发送。
	_taskSign string
}

// NewTaobaoJstSmsMessageDirectBatchsendRequest 初始化TaobaoJstSmsMessageDirectBatchsendAPIRequest对象
func NewTaobaoJstSmsMessageDirectBatchsendRequest() *TaobaoJstSmsMessageDirectBatchsendAPIRequest {
	return &TaobaoJstSmsMessageDirectBatchsendAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoJstSmsMessageDirectBatchsendAPIRequest) GetApiMethodName() string {
	return "taobao.jst.sms.message.direct.batchsend"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoJstSmsMessageDirectBatchsendAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is SignName Setter
// 短信签名
func (r *TaobaoJstSmsMessageDirectBatchsendAPIRequest) SetSignName(_signName string) error {
	r._signName = _signName
	r.Set("sign_name", _signName)
	return nil
}

// Get SignName Getter
func (r TaobaoJstSmsMessageDirectBatchsendAPIRequest) GetSignName() string {
	return r._signName
}

// Set is Url Setter
// 短信中带的短链，如果不传则没有短信效果数据
func (r *TaobaoJstSmsMessageDirectBatchsendAPIRequest) SetUrl(_url string) error {
	r._url = _url
	r.Set("url", _url)
	return nil
}

// Get Url Getter
func (r TaobaoJstSmsMessageDirectBatchsendAPIRequest) GetUrl() string {
	return r._url
}

// Set is SmsTemplateCode Setter
// 短信模版CODE，必须为全变量模板
func (r *TaobaoJstSmsMessageDirectBatchsendAPIRequest) SetSmsTemplateCode(_smsTemplateCode string) error {
	r._smsTemplateCode = _smsTemplateCode
	r.Set("sms_template_code", _smsTemplateCode)
	return nil
}

// Get SmsTemplateCode Getter
func (r TaobaoJstSmsMessageDirectBatchsendAPIRequest) GetSmsTemplateCode() string {
	return r._smsTemplateCode
}

// Set is RecNum Setter
// 短信接收号码，json格式，最多200个号码
func (r *TaobaoJstSmsMessageDirectBatchsendAPIRequest) SetRecNum(_recNum string) error {
	r._recNum = _recNum
	r.Set("rec_num", _recNum)
	return nil
}

// Get RecNum Getter
func (r TaobaoJstSmsMessageDirectBatchsendAPIRequest) GetRecNum() string {
	return r._recNum
}

// Set is SmsContent Setter
// 短信内容，如果带${url}则会被入参url替换
func (r *TaobaoJstSmsMessageDirectBatchsendAPIRequest) SetSmsContent(_smsContent string) error {
	r._smsContent = _smsContent
	r.Set("sms_content", _smsContent)
	return nil
}

// Get SmsContent Getter
func (r TaobaoJstSmsMessageDirectBatchsendAPIRequest) GetSmsContent() string {
	return r._smsContent
}

// Set is ExtendNum Setter
// 短信扩展码（JSON字符串数组格式），拓展码个数需要和电话号码个数一致。
func (r *TaobaoJstSmsMessageDirectBatchsendAPIRequest) SetExtendNum(_extendNum string) error {
	r._extendNum = _extendNum
	r.Set("extend_num", _extendNum)
	return nil
}

// Get ExtendNum Getter
func (r TaobaoJstSmsMessageDirectBatchsendAPIRequest) GetExtendNum() string {
	return r._extendNum
}

// Set is TaskCode Setter
// 短信任务code，没有请先创建。
func (r *TaobaoJstSmsMessageDirectBatchsendAPIRequest) SetTaskCode(_taskCode string) error {
	r._taskCode = _taskCode
	r.Set("task_code", _taskCode)
	return nil
}

// Get TaskCode Getter
func (r TaobaoJstSmsMessageDirectBatchsendAPIRequest) GetTaskCode() string {
	return r._taskCode
}

// Set is TaskSign Setter
// 一个在taskcode下唯一的随机字符串，对于taskSign相同的请求平台认为是商家的同一次短信发送。
func (r *TaobaoJstSmsMessageDirectBatchsendAPIRequest) SetTaskSign(_taskSign string) error {
	r._taskSign = _taskSign
	r.Set("task_sign", _taskSign)
	return nil
}

// Get TaskSign Getter
func (r TaobaoJstSmsMessageDirectBatchsendAPIRequest) GetTaskSign() string {
	return r._taskSign
}
