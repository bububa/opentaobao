package aliqin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
短信发送 APIRequest
alibaba.aliqin.fc.sms.num.send

向指定手机号码发送模板短信，模板内可设置部分变量。使用前需要在阿里大于管理中心添加短信签名与短信模板。测试时请直接使用正式环境HTTP请求地址。
【重要】批量发送（一次传递多个号码eg:1381111111,1382222222）会产生相应的延迟，触达时间要求高的建议单条发送
*/
type AlibabaAliqinFcSmsNumSendRequest struct {
    model.Params

    // 公共回传参数，在“消息返回”中会透传回该参数；举例：用户可以传入自己下级的会员ID，在消息返回时，该会员ID会包含在内，用户可以根据该会员ID识别是哪位会员使用了你的应用
    extend   string 

    // 短信类型，传入值请填写normal
    smsType   string 

    // 短信签名，传入的短信签名必须是在阿里大于“管理中心-验证码/短信通知/推广短信-配置短信签名”中的可用签名。如“阿里大于”已在短信签名管理中通过审核，则可传入”阿里大于“（传参时去掉引号）作为短信签名。短信效果示例：【阿里大于】欢迎使用阿里大于服务。
    smsFreeSignName   string 

    // 短信模板变量，传参规则{"key":"value"}，key的名字须和申请模板中的变量名一致，多个变量之间以逗号隔开。示例：针对模板“验证码${code}，您正在进行${product}身份验证，打死不要告诉别人哦！”，传参时需传入{"code":"1234","product":"alidayu"}
    smsParam   string 

    // 短信接收号码。支持单个或多个手机号码，传入号码为11位手机号码，不能加0或+86。群发短信需传入多个号码，以英文逗号分隔，一次调用最多传入200个号码。示例：18600000000,13911111111,13322222222
    recNum   string 

    // 短信模板ID，传入的模板必须是在阿里大于“管理中心-短信模板管理”中的可用模板。示例：SMS_585014
    smsTemplateCode   string 

}

func NewAlibabaAliqinFcSmsNumSendRequest() *AlibabaAliqinFcSmsNumSendRequest{
    return &AlibabaAliqinFcSmsNumSendRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAliqinFcSmsNumSendRequest) GetApiMethodName() string {
    return "alibaba.aliqin.fc.sms.num.send"
}

func (r AlibabaAliqinFcSmsNumSendRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAliqinFcSmsNumSendRequest) SetExtend(extend string) error {
    r.extend = extend
    r.Set("extend", extend)
    return nil
}

func (r AlibabaAliqinFcSmsNumSendRequest) GetExtend() string {
    return r.extend
}

func (r *AlibabaAliqinFcSmsNumSendRequest) SetSmsType(smsType string) error {
    r.smsType = smsType
    r.Set("sms_type", smsType)
    return nil
}

func (r AlibabaAliqinFcSmsNumSendRequest) GetSmsType() string {
    return r.smsType
}

func (r *AlibabaAliqinFcSmsNumSendRequest) SetSmsFreeSignName(smsFreeSignName string) error {
    r.smsFreeSignName = smsFreeSignName
    r.Set("sms_free_sign_name", smsFreeSignName)
    return nil
}

func (r AlibabaAliqinFcSmsNumSendRequest) GetSmsFreeSignName() string {
    return r.smsFreeSignName
}

func (r *AlibabaAliqinFcSmsNumSendRequest) SetSmsParam(smsParam string) error {
    r.smsParam = smsParam
    r.Set("sms_param", smsParam)
    return nil
}

func (r AlibabaAliqinFcSmsNumSendRequest) GetSmsParam() string {
    return r.smsParam
}

func (r *AlibabaAliqinFcSmsNumSendRequest) SetRecNum(recNum string) error {
    r.recNum = recNum
    r.Set("rec_num", recNum)
    return nil
}

func (r AlibabaAliqinFcSmsNumSendRequest) GetRecNum() string {
    return r.recNum
}

func (r *AlibabaAliqinFcSmsNumSendRequest) SetSmsTemplateCode(smsTemplateCode string) error {
    r.smsTemplateCode = smsTemplateCode
    r.Set("sms_template_code", smsTemplateCode)
    return nil
}

func (r AlibabaAliqinFcSmsNumSendRequest) GetSmsTemplateCode() string {
    return r.smsTemplateCode
}

