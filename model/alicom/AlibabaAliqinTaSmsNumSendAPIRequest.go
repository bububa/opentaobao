package alicom

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAliqinTaSmsNumSendAPIRequest 短信发送 API请求
// alibaba.aliqin.ta.sms.num.send
//
// 短信发送
type AlibabaAliqinTaSmsNumSendAPIRequest struct {
	model.Params
	// 接收号码
	_recNum string
	// 短信模板CODE
	_smsTemplateCode string
	// 公共回传参数
	_extend string
	// 短信签名
	_smsFreeSignName string
	// 短信模板变量，AckNum是变量参数
	_smsParam string
	// 类型，normal：短信
	_smsType string
	// 商家自定义扩展码
	_extendCode string
	// 商家自定义扩展名,例如店铺nick
	_extendName string
}

// NewAlibabaAliqinTaSmsNumSendRequest 初始化AlibabaAliqinTaSmsNumSendAPIRequest对象
func NewAlibabaAliqinTaSmsNumSendRequest() *AlibabaAliqinTaSmsNumSendAPIRequest {
	return &AlibabaAliqinTaSmsNumSendAPIRequest{
		Params: model.NewParams(8),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAliqinTaSmsNumSendAPIRequest) Reset() {
	r._recNum = ""
	r._smsTemplateCode = ""
	r._extend = ""
	r._smsFreeSignName = ""
	r._smsParam = ""
	r._smsType = ""
	r._extendCode = ""
	r._extendName = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAliqinTaSmsNumSendAPIRequest) GetApiMethodName() string {
	return "alibaba.aliqin.ta.sms.num.send"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAliqinTaSmsNumSendAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAliqinTaSmsNumSendAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRecNum is RecNum Setter
// 接收号码
func (r *AlibabaAliqinTaSmsNumSendAPIRequest) SetRecNum(_recNum string) error {
	r._recNum = _recNum
	r.Set("rec_num", _recNum)
	return nil
}

// GetRecNum RecNum Getter
func (r AlibabaAliqinTaSmsNumSendAPIRequest) GetRecNum() string {
	return r._recNum
}

// SetSmsTemplateCode is SmsTemplateCode Setter
// 短信模板CODE
func (r *AlibabaAliqinTaSmsNumSendAPIRequest) SetSmsTemplateCode(_smsTemplateCode string) error {
	r._smsTemplateCode = _smsTemplateCode
	r.Set("sms_template_code", _smsTemplateCode)
	return nil
}

// GetSmsTemplateCode SmsTemplateCode Getter
func (r AlibabaAliqinTaSmsNumSendAPIRequest) GetSmsTemplateCode() string {
	return r._smsTemplateCode
}

// SetExtend is Extend Setter
// 公共回传参数
func (r *AlibabaAliqinTaSmsNumSendAPIRequest) SetExtend(_extend string) error {
	r._extend = _extend
	r.Set("extend", _extend)
	return nil
}

// GetExtend Extend Getter
func (r AlibabaAliqinTaSmsNumSendAPIRequest) GetExtend() string {
	return r._extend
}

// SetSmsFreeSignName is SmsFreeSignName Setter
// 短信签名
func (r *AlibabaAliqinTaSmsNumSendAPIRequest) SetSmsFreeSignName(_smsFreeSignName string) error {
	r._smsFreeSignName = _smsFreeSignName
	r.Set("sms_free_sign_name", _smsFreeSignName)
	return nil
}

// GetSmsFreeSignName SmsFreeSignName Getter
func (r AlibabaAliqinTaSmsNumSendAPIRequest) GetSmsFreeSignName() string {
	return r._smsFreeSignName
}

// SetSmsParam is SmsParam Setter
// 短信模板变量，AckNum是变量参数
func (r *AlibabaAliqinTaSmsNumSendAPIRequest) SetSmsParam(_smsParam string) error {
	r._smsParam = _smsParam
	r.Set("sms_param", _smsParam)
	return nil
}

// GetSmsParam SmsParam Getter
func (r AlibabaAliqinTaSmsNumSendAPIRequest) GetSmsParam() string {
	return r._smsParam
}

// SetSmsType is SmsType Setter
// 类型，normal：短信
func (r *AlibabaAliqinTaSmsNumSendAPIRequest) SetSmsType(_smsType string) error {
	r._smsType = _smsType
	r.Set("sms_type", _smsType)
	return nil
}

// GetSmsType SmsType Getter
func (r AlibabaAliqinTaSmsNumSendAPIRequest) GetSmsType() string {
	return r._smsType
}

// SetExtendCode is ExtendCode Setter
// 商家自定义扩展码
func (r *AlibabaAliqinTaSmsNumSendAPIRequest) SetExtendCode(_extendCode string) error {
	r._extendCode = _extendCode
	r.Set("extend_code", _extendCode)
	return nil
}

// GetExtendCode ExtendCode Getter
func (r AlibabaAliqinTaSmsNumSendAPIRequest) GetExtendCode() string {
	return r._extendCode
}

// SetExtendName is ExtendName Setter
// 商家自定义扩展名,例如店铺nick
func (r *AlibabaAliqinTaSmsNumSendAPIRequest) SetExtendName(_extendName string) error {
	r._extendName = _extendName
	r.Set("extend_name", _extendName)
	return nil
}

// GetExtendName ExtendName Getter
func (r AlibabaAliqinTaSmsNumSendAPIRequest) GetExtendName() string {
	return r._extendName
}

var poolAlibabaAliqinTaSmsNumSendAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAliqinTaSmsNumSendRequest()
	},
}

// GetAlibabaAliqinTaSmsNumSendRequest 从 sync.Pool 获取 AlibabaAliqinTaSmsNumSendAPIRequest
func GetAlibabaAliqinTaSmsNumSendAPIRequest() *AlibabaAliqinTaSmsNumSendAPIRequest {
	return poolAlibabaAliqinTaSmsNumSendAPIRequest.Get().(*AlibabaAliqinTaSmsNumSendAPIRequest)
}

// ReleaseAlibabaAliqinTaSmsNumSendAPIRequest 将 AlibabaAliqinTaSmsNumSendAPIRequest 放入 sync.Pool
func ReleaseAlibabaAliqinTaSmsNumSendAPIRequest(v *AlibabaAliqinTaSmsNumSendAPIRequest) {
	v.Reset()
	poolAlibabaAliqinTaSmsNumSendAPIRequest.Put(v)
}
