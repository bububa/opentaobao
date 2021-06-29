package alicom

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
短信发送 API请求
alibaba.aliqin.ta.sms.num.send

短信发送
*/
type AlibabaAliqinTaSmsNumSendRequest struct {
    model.Params
    // 公共回传参数
    _extend   string
    // 类型，normal：短信
    _smsType   string
    // 短信签名
    _smsFreeSignName   string
    // 短信模板变量，AckNum是变量参数
    _smsParam   string
    // 接收号码
    _recNum   string
    // 短信模板CODE
    _smsTemplateCode   string
    // 商家自定义扩展码
    _extendCode   string
    // 商家自定义扩展名,例如店铺nick
    _extendName   string
}

// 初始化AlibabaAliqinTaSmsNumSendRequest对象
func NewAlibabaAliqinTaSmsNumSendRequest() *AlibabaAliqinTaSmsNumSendRequest{
    return &AlibabaAliqinTaSmsNumSendRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAliqinTaSmsNumSendRequest) GetApiMethodName() string {
    return "alibaba.aliqin.ta.sms.num.send"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAliqinTaSmsNumSendRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Extend Setter
// 公共回传参数
func (r *AlibabaAliqinTaSmsNumSendRequest) SetExtend(_extend string) error {
    r._extend = _extend
    r.Set("extend", _extend)
    return nil
}

// Extend Getter
func (r AlibabaAliqinTaSmsNumSendRequest) GetExtend() string {
    return r._extend
}
// SmsType Setter
// 类型，normal：短信
func (r *AlibabaAliqinTaSmsNumSendRequest) SetSmsType(_smsType string) error {
    r._smsType = _smsType
    r.Set("sms_type", _smsType)
    return nil
}

// SmsType Getter
func (r AlibabaAliqinTaSmsNumSendRequest) GetSmsType() string {
    return r._smsType
}
// SmsFreeSignName Setter
// 短信签名
func (r *AlibabaAliqinTaSmsNumSendRequest) SetSmsFreeSignName(_smsFreeSignName string) error {
    r._smsFreeSignName = _smsFreeSignName
    r.Set("sms_free_sign_name", _smsFreeSignName)
    return nil
}

// SmsFreeSignName Getter
func (r AlibabaAliqinTaSmsNumSendRequest) GetSmsFreeSignName() string {
    return r._smsFreeSignName
}
// SmsParam Setter
// 短信模板变量，AckNum是变量参数
func (r *AlibabaAliqinTaSmsNumSendRequest) SetSmsParam(_smsParam string) error {
    r._smsParam = _smsParam
    r.Set("sms_param", _smsParam)
    return nil
}

// SmsParam Getter
func (r AlibabaAliqinTaSmsNumSendRequest) GetSmsParam() string {
    return r._smsParam
}
// RecNum Setter
// 接收号码
func (r *AlibabaAliqinTaSmsNumSendRequest) SetRecNum(_recNum string) error {
    r._recNum = _recNum
    r.Set("rec_num", _recNum)
    return nil
}

// RecNum Getter
func (r AlibabaAliqinTaSmsNumSendRequest) GetRecNum() string {
    return r._recNum
}
// SmsTemplateCode Setter
// 短信模板CODE
func (r *AlibabaAliqinTaSmsNumSendRequest) SetSmsTemplateCode(_smsTemplateCode string) error {
    r._smsTemplateCode = _smsTemplateCode
    r.Set("sms_template_code", _smsTemplateCode)
    return nil
}

// SmsTemplateCode Getter
func (r AlibabaAliqinTaSmsNumSendRequest) GetSmsTemplateCode() string {
    return r._smsTemplateCode
}
// ExtendCode Setter
// 商家自定义扩展码
func (r *AlibabaAliqinTaSmsNumSendRequest) SetExtendCode(_extendCode string) error {
    r._extendCode = _extendCode
    r.Set("extend_code", _extendCode)
    return nil
}

// ExtendCode Getter
func (r AlibabaAliqinTaSmsNumSendRequest) GetExtendCode() string {
    return r._extendCode
}
// ExtendName Setter
// 商家自定义扩展名,例如店铺nick
func (r *AlibabaAliqinTaSmsNumSendRequest) SetExtendName(_extendName string) error {
    r._extendName = _extendName
    r.Set("extend_name", _extendName)
    return nil
}

// ExtendName Getter
func (r AlibabaAliqinTaSmsNumSendRequest) GetExtendName() string {
    return r._extendName
}
