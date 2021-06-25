package alicom

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
短信发送 APIRequest
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

func NewAlibabaAliqinTaSmsNumSendRequest() *AlibabaAliqinTaSmsNumSendRequest{
    return &AlibabaAliqinTaSmsNumSendRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAliqinTaSmsNumSendRequest) GetApiMethodName() string {
    return "alibaba.aliqin.ta.sms.num.send"
}

func (r AlibabaAliqinTaSmsNumSendRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAliqinTaSmsNumSendRequest) SetExtend(extend string) error {
    r.extend = extend
    r.Set("extend", extend)
    return nil
}

func (r AlibabaAliqinTaSmsNumSendRequest) GetExtend() string {
    return r.extend
}

func (r *AlibabaAliqinTaSmsNumSendRequest) SetSmsType(smsType string) error {
    r.smsType = smsType
    r.Set("sms_type", smsType)
    return nil
}

func (r AlibabaAliqinTaSmsNumSendRequest) GetSmsType() string {
    return r.smsType
}

func (r *AlibabaAliqinTaSmsNumSendRequest) SetSmsFreeSignName(smsFreeSignName string) error {
    r.smsFreeSignName = smsFreeSignName
    r.Set("sms_free_sign_name", smsFreeSignName)
    return nil
}

func (r AlibabaAliqinTaSmsNumSendRequest) GetSmsFreeSignName() string {
    return r.smsFreeSignName
}

func (r *AlibabaAliqinTaSmsNumSendRequest) SetSmsParam(smsParam string) error {
    r.smsParam = smsParam
    r.Set("sms_param", smsParam)
    return nil
}

func (r AlibabaAliqinTaSmsNumSendRequest) GetSmsParam() string {
    return r.smsParam
}

func (r *AlibabaAliqinTaSmsNumSendRequest) SetRecNum(recNum string) error {
    r.recNum = recNum
    r.Set("rec_num", recNum)
    return nil
}

func (r AlibabaAliqinTaSmsNumSendRequest) GetRecNum() string {
    return r.recNum
}

func (r *AlibabaAliqinTaSmsNumSendRequest) SetSmsTemplateCode(smsTemplateCode string) error {
    r.smsTemplateCode = smsTemplateCode
    r.Set("sms_template_code", smsTemplateCode)
    return nil
}

func (r AlibabaAliqinTaSmsNumSendRequest) GetSmsTemplateCode() string {
    return r.smsTemplateCode
}

func (r *AlibabaAliqinTaSmsNumSendRequest) SetExtendCode(extendCode string) error {
    r.extendCode = extendCode
    r.Set("extend_code", extendCode)
    return nil
}

func (r AlibabaAliqinTaSmsNumSendRequest) GetExtendCode() string {
    return r.extendCode
}

func (r *AlibabaAliqinTaSmsNumSendRequest) SetExtendName(extendName string) error {
    r.extendName = extendName
    r.Set("extend_name", extendName)
    return nil
}

func (r AlibabaAliqinTaSmsNumSendRequest) GetExtendName() string {
    return r.extendName
}

