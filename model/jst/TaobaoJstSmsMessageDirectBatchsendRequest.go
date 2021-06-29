package jst

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
聚石塔新短信发送接口 API请求
taobao.jst.sms.message.direct.batchsend

聚石塔所见即所得的短信发送接口
*/
type TaobaoJstSmsMessageDirectBatchsendRequest struct {
    model.Params
    // 短信签名
    _signName   string
    // 短信中带的短链，如果不传则没有短信效果数据
    _url   string
    // 短信模版CODE，必须为全变量模板
    _smsTemplateCode   string
    // 短信接收号码，json格式，最多200个号码
    _recNum   string
    // 短信内容，如果带${url}则会被入参url替换
    _smsContent   string
    // 短信扩展码（JSON字符串数组格式），拓展码个数需要和电话号码个数一致。
    _extendNum   string
    // 短信任务code，没有请先创建。
    _taskCode   string
    // 一个在taskcode下唯一的随机字符串，对于taskSign相同的请求平台认为是商家的同一次短信发送。
    _taskSign   string
}

// 初始化TaobaoJstSmsMessageDirectBatchsendRequest对象
func NewTaobaoJstSmsMessageDirectBatchsendRequest() *TaobaoJstSmsMessageDirectBatchsendRequest{
    return &TaobaoJstSmsMessageDirectBatchsendRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoJstSmsMessageDirectBatchsendRequest) GetApiMethodName() string {
    return "taobao.jst.sms.message.direct.batchsend"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoJstSmsMessageDirectBatchsendRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// SignName Setter
// 短信签名
func (r *TaobaoJstSmsMessageDirectBatchsendRequest) SetSignName(_signName string) error {
    r._signName = _signName
    r.Set("sign_name", _signName)
    return nil
}

// SignName Getter
func (r TaobaoJstSmsMessageDirectBatchsendRequest) GetSignName() string {
    return r._signName
}
// Url Setter
// 短信中带的短链，如果不传则没有短信效果数据
func (r *TaobaoJstSmsMessageDirectBatchsendRequest) SetUrl(_url string) error {
    r._url = _url
    r.Set("url", _url)
    return nil
}

// Url Getter
func (r TaobaoJstSmsMessageDirectBatchsendRequest) GetUrl() string {
    return r._url
}
// SmsTemplateCode Setter
// 短信模版CODE，必须为全变量模板
func (r *TaobaoJstSmsMessageDirectBatchsendRequest) SetSmsTemplateCode(_smsTemplateCode string) error {
    r._smsTemplateCode = _smsTemplateCode
    r.Set("sms_template_code", _smsTemplateCode)
    return nil
}

// SmsTemplateCode Getter
func (r TaobaoJstSmsMessageDirectBatchsendRequest) GetSmsTemplateCode() string {
    return r._smsTemplateCode
}
// RecNum Setter
// 短信接收号码，json格式，最多200个号码
func (r *TaobaoJstSmsMessageDirectBatchsendRequest) SetRecNum(_recNum string) error {
    r._recNum = _recNum
    r.Set("rec_num", _recNum)
    return nil
}

// RecNum Getter
func (r TaobaoJstSmsMessageDirectBatchsendRequest) GetRecNum() string {
    return r._recNum
}
// SmsContent Setter
// 短信内容，如果带${url}则会被入参url替换
func (r *TaobaoJstSmsMessageDirectBatchsendRequest) SetSmsContent(_smsContent string) error {
    r._smsContent = _smsContent
    r.Set("sms_content", _smsContent)
    return nil
}

// SmsContent Getter
func (r TaobaoJstSmsMessageDirectBatchsendRequest) GetSmsContent() string {
    return r._smsContent
}
// ExtendNum Setter
// 短信扩展码（JSON字符串数组格式），拓展码个数需要和电话号码个数一致。
func (r *TaobaoJstSmsMessageDirectBatchsendRequest) SetExtendNum(_extendNum string) error {
    r._extendNum = _extendNum
    r.Set("extend_num", _extendNum)
    return nil
}

// ExtendNum Getter
func (r TaobaoJstSmsMessageDirectBatchsendRequest) GetExtendNum() string {
    return r._extendNum
}
// TaskCode Setter
// 短信任务code，没有请先创建。
func (r *TaobaoJstSmsMessageDirectBatchsendRequest) SetTaskCode(_taskCode string) error {
    r._taskCode = _taskCode
    r.Set("task_code", _taskCode)
    return nil
}

// TaskCode Getter
func (r TaobaoJstSmsMessageDirectBatchsendRequest) GetTaskCode() string {
    return r._taskCode
}
// TaskSign Setter
// 一个在taskcode下唯一的随机字符串，对于taskSign相同的请求平台认为是商家的同一次短信发送。
func (r *TaobaoJstSmsMessageDirectBatchsendRequest) SetTaskSign(_taskSign string) error {
    r._taskSign = _taskSign
    r.Set("task_sign", _taskSign)
    return nil
}

// TaskSign Getter
func (r TaobaoJstSmsMessageDirectBatchsendRequest) GetTaskSign() string {
    return r._taskSign
}
