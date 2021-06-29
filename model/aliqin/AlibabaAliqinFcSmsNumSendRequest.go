package aliqin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
短信发送 API请求
alibaba.aliqin.fc.sms.num.send

向指定手机号码发送模板短信，模板内可设置部分变量。使用前需要在阿里大于管理中心添加短信签名与短信模板。测试时请直接使用正式环境HTTP请求地址。
【重要】批量发送（一次传递多个号码eg:1381111111,1382222222）会产生相应的延迟，触达时间要求高的建议单条发送
*/
type AlibabaAliqinFcSmsNumSendRequest struct {
    model.Params
    // 公共回传参数，在“消息返回”中会透传回该参数；举例：用户可以传入自己下级的会员ID，在消息返回时，该会员ID会包含在内，用户可以根据该会员ID识别是哪位会员使用了你的应用
    _extend   string
    // 短信类型，传入值请填写normal
    _smsType   string
    // 短信签名，传入的短信签名必须是在阿里大于“管理中心-验证码/短信通知/推广短信-配置短信签名”中的可用签名。如“阿里大于”已在短信签名管理中通过审核，则可传入”阿里大于“（传参时去掉引号）作为短信签名。短信效果示例：【阿里大于】欢迎使用阿里大于服务。
    _smsFreeSignName   string
    // 短信模板变量，传参规则{"key":"value"}，key的名字须和申请模板中的变量名一致，多个变量之间以逗号隔开。示例：针对模板“验证码${code}，您正在进行${product}身份验证，打死不要告诉别人哦！”，传参时需传入{"code":"1234","product":"alidayu"}
    _smsParam   string
    // 短信接收号码。支持单个或多个手机号码，传入号码为11位手机号码，不能加0或+86。群发短信需传入多个号码，以英文逗号分隔，一次调用最多传入200个号码。示例：18600000000,13911111111,13322222222
    _recNum   string
    // 短信模板ID，传入的模板必须是在阿里大于“管理中心-短信模板管理”中的可用模板。示例：SMS_585014
    _smsTemplateCode   string
}

// 初始化AlibabaAliqinFcSmsNumSendRequest对象
func NewAlibabaAliqinFcSmsNumSendRequest() *AlibabaAliqinFcSmsNumSendRequest{
    return &AlibabaAliqinFcSmsNumSendRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAliqinFcSmsNumSendRequest) GetApiMethodName() string {
    return "alibaba.aliqin.fc.sms.num.send"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAliqinFcSmsNumSendRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Extend Setter
// 公共回传参数，在“消息返回”中会透传回该参数；举例：用户可以传入自己下级的会员ID，在消息返回时，该会员ID会包含在内，用户可以根据该会员ID识别是哪位会员使用了你的应用
func (r *AlibabaAliqinFcSmsNumSendRequest) SetExtend(_extend string) error {
    r._extend = _extend
    r.Set("extend", _extend)
    return nil
}

// Extend Getter
func (r AlibabaAliqinFcSmsNumSendRequest) GetExtend() string {
    return r._extend
}
// SmsType Setter
// 短信类型，传入值请填写normal
func (r *AlibabaAliqinFcSmsNumSendRequest) SetSmsType(_smsType string) error {
    r._smsType = _smsType
    r.Set("sms_type", _smsType)
    return nil
}

// SmsType Getter
func (r AlibabaAliqinFcSmsNumSendRequest) GetSmsType() string {
    return r._smsType
}
// SmsFreeSignName Setter
// 短信签名，传入的短信签名必须是在阿里大于“管理中心-验证码/短信通知/推广短信-配置短信签名”中的可用签名。如“阿里大于”已在短信签名管理中通过审核，则可传入”阿里大于“（传参时去掉引号）作为短信签名。短信效果示例：【阿里大于】欢迎使用阿里大于服务。
func (r *AlibabaAliqinFcSmsNumSendRequest) SetSmsFreeSignName(_smsFreeSignName string) error {
    r._smsFreeSignName = _smsFreeSignName
    r.Set("sms_free_sign_name", _smsFreeSignName)
    return nil
}

// SmsFreeSignName Getter
func (r AlibabaAliqinFcSmsNumSendRequest) GetSmsFreeSignName() string {
    return r._smsFreeSignName
}
// SmsParam Setter
// 短信模板变量，传参规则{"key":"value"}，key的名字须和申请模板中的变量名一致，多个变量之间以逗号隔开。示例：针对模板“验证码${code}，您正在进行${product}身份验证，打死不要告诉别人哦！”，传参时需传入{"code":"1234","product":"alidayu"}
func (r *AlibabaAliqinFcSmsNumSendRequest) SetSmsParam(_smsParam string) error {
    r._smsParam = _smsParam
    r.Set("sms_param", _smsParam)
    return nil
}

// SmsParam Getter
func (r AlibabaAliqinFcSmsNumSendRequest) GetSmsParam() string {
    return r._smsParam
}
// RecNum Setter
// 短信接收号码。支持单个或多个手机号码，传入号码为11位手机号码，不能加0或+86。群发短信需传入多个号码，以英文逗号分隔，一次调用最多传入200个号码。示例：18600000000,13911111111,13322222222
func (r *AlibabaAliqinFcSmsNumSendRequest) SetRecNum(_recNum string) error {
    r._recNum = _recNum
    r.Set("rec_num", _recNum)
    return nil
}

// RecNum Getter
func (r AlibabaAliqinFcSmsNumSendRequest) GetRecNum() string {
    return r._recNum
}
// SmsTemplateCode Setter
// 短信模板ID，传入的模板必须是在阿里大于“管理中心-短信模板管理”中的可用模板。示例：SMS_585014
func (r *AlibabaAliqinFcSmsNumSendRequest) SetSmsTemplateCode(_smsTemplateCode string) error {
    r._smsTemplateCode = _smsTemplateCode
    r.Set("sms_template_code", _smsTemplateCode)
    return nil
}

// SmsTemplateCode Getter
func (r AlibabaAliqinFcSmsNumSendRequest) GetSmsTemplateCode() string {
    return r._smsTemplateCode
}
