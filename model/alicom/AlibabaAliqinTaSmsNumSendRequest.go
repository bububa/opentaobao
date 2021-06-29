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
    extend   string
    // 类型，normal：短信
    smsType   string
    // 短信签名
    smsFreeSignName   string
    // 短信模板变量，AckNum是变量参数
    smsParam   string
    // 接收号码
    recNum   string
    // 短信模板CODE
    smsTemplateCode   string
    // 商家自定义扩展码
    extendCode   string
    // 商家自定义扩展名,例如店铺nick
    extendName   string
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
func (r *AlibabaAliqinTaSmsNumSendRequest) SetExtend(extend string) error {
    r.extend = extend
    r.Set("extend", extend)
    return nil
}

// Extend Getter
func (r AlibabaAliqinTaSmsNumSendRequest) GetExtend() string {
    return r.extend
}
// SmsType Setter
// 类型，normal：短信
func (r *AlibabaAliqinTaSmsNumSendRequest) SetSmsType(smsType string) error {
    r.smsType = smsType
    r.Set("sms_type", smsType)
    return nil
}

// SmsType Getter
func (r AlibabaAliqinTaSmsNumSendRequest) GetSmsType() string {
    return r.smsType
}
// SmsFreeSignName Setter
// 短信签名
func (r *AlibabaAliqinTaSmsNumSendRequest) SetSmsFreeSignName(smsFreeSignName string) error {
    r.smsFreeSignName = smsFreeSignName
    r.Set("sms_free_sign_name", smsFreeSignName)
    return nil
}

// SmsFreeSignName Getter
func (r AlibabaAliqinTaSmsNumSendRequest) GetSmsFreeSignName() string {
    return r.smsFreeSignName
}
// SmsParam Setter
// 短信模板变量，AckNum是变量参数
func (r *AlibabaAliqinTaSmsNumSendRequest) SetSmsParam(smsParam string) error {
    r.smsParam = smsParam
    r.Set("sms_param", smsParam)
    return nil
}

// SmsParam Getter
func (r AlibabaAliqinTaSmsNumSendRequest) GetSmsParam() string {
    return r.smsParam
}
// RecNum Setter
// 接收号码
func (r *AlibabaAliqinTaSmsNumSendRequest) SetRecNum(recNum string) error {
    r.recNum = recNum
    r.Set("rec_num", recNum)
    return nil
}

// RecNum Getter
func (r AlibabaAliqinTaSmsNumSendRequest) GetRecNum() string {
    return r.recNum
}
// SmsTemplateCode Setter
// 短信模板CODE
func (r *AlibabaAliqinTaSmsNumSendRequest) SetSmsTemplateCode(smsTemplateCode string) error {
    r.smsTemplateCode = smsTemplateCode
    r.Set("sms_template_code", smsTemplateCode)
    return nil
}

// SmsTemplateCode Getter
func (r AlibabaAliqinTaSmsNumSendRequest) GetSmsTemplateCode() string {
    return r.smsTemplateCode
}
// ExtendCode Setter
// 商家自定义扩展码
func (r *AlibabaAliqinTaSmsNumSendRequest) SetExtendCode(extendCode string) error {
    r.extendCode = extendCode
    r.Set("extend_code", extendCode)
    return nil
}

// ExtendCode Getter
func (r AlibabaAliqinTaSmsNumSendRequest) GetExtendCode() string {
    return r.extendCode
}
// ExtendName Setter
// 商家自定义扩展名,例如店铺nick
func (r *AlibabaAliqinTaSmsNumSendRequest) SetExtendName(extendName string) error {
    r.extendName = extendName
    r.Set("extend_name", extendName)
    return nil
}

// ExtendName Getter
func (r AlibabaAliqinTaSmsNumSendRequest) GetExtendName() string {
    return r.extendName
}
