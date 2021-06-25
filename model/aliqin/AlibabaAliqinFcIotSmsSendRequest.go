package aliqin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
物联网短信发送 APIRequest
alibaba.aliqin.fc.iot.sms.send

发送物联网短信，只允许使用物联网短信模板
*/
type AlibabaAliqinFcIotSmsSendRequest struct {
    model.Params

    // 公共回传参数，在“消息返回”中会透传回该参数；举例：用户可以传入自己下级的会员ID，在消息返回时，该会员ID会包含在内，用户可以根据该会员ID识别是哪位会员使用了你的应用
    extend   string 

    // 短信类型，传入值请填写normal
    smsType   string 

    // 短信模板变量，传参规则{"key":"value"}，key的名字须和申请模板中的变量名一致，多个变量之间以逗号隔开。示例：针对模板“验证码${code}，您正在进行${product}身份验证，打死不要告诉别人哦！”，传参时需传入{"code":"1234","product":"alidayu"}
    smsParam   string 

    // 短信接收号码。
    recNum   string 

    // 短信模板ID，传入的模板必须是在阿里大于“管理中心-短信模板管理”中的可用模板。示例：SMS_585014
    smsTemplateCode   string 

}

func NewAlibabaAliqinFcIotSmsSendRequest() *AlibabaAliqinFcIotSmsSendRequest{
    return &AlibabaAliqinFcIotSmsSendRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAliqinFcIotSmsSendRequest) GetApiMethodName() string {
    return "alibaba.aliqin.fc.iot.sms.send"
}

func (r AlibabaAliqinFcIotSmsSendRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAliqinFcIotSmsSendRequest) SetExtend(extend string) error {
    r.extend = extend
    r.Set("extend", extend)
    return nil
}

func (r AlibabaAliqinFcIotSmsSendRequest) GetExtend() string {
    return r.extend
}

func (r *AlibabaAliqinFcIotSmsSendRequest) SetSmsType(smsType string) error {
    r.smsType = smsType
    r.Set("sms_type", smsType)
    return nil
}

func (r AlibabaAliqinFcIotSmsSendRequest) GetSmsType() string {
    return r.smsType
}

func (r *AlibabaAliqinFcIotSmsSendRequest) SetSmsParam(smsParam string) error {
    r.smsParam = smsParam
    r.Set("sms_param", smsParam)
    return nil
}

func (r AlibabaAliqinFcIotSmsSendRequest) GetSmsParam() string {
    return r.smsParam
}

func (r *AlibabaAliqinFcIotSmsSendRequest) SetRecNum(recNum string) error {
    r.recNum = recNum
    r.Set("rec_num", recNum)
    return nil
}

func (r AlibabaAliqinFcIotSmsSendRequest) GetRecNum() string {
    return r.recNum
}

func (r *AlibabaAliqinFcIotSmsSendRequest) SetSmsTemplateCode(smsTemplateCode string) error {
    r.smsTemplateCode = smsTemplateCode
    r.Set("sms_template_code", smsTemplateCode)
    return nil
}

func (r AlibabaAliqinFcIotSmsSendRequest) GetSmsTemplateCode() string {
    return r.smsTemplateCode
}

