package jst

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
聚石塔新短信发送接口 APIRequest
taobao.jst.sms.message.direct.batchsend

聚石塔所见即所得的短信发送接口
*/
type TaobaoJstSmsMessageDirectBatchsendRequest struct {
    model.Params

    // 短信签名
    signName   string 

    // 短信中带的短链，如果不传则没有短信效果数据
    url   string 

    // 短信模版CODE，必须为全变量模板
    smsTemplateCode   string 

    // 短信接收号码，json格式，最多200个号码
    recNum   string 

    // 短信内容，如果带${url}则会被入参url替换
    smsContent   string 

    // 短信扩展码（JSON字符串数组格式），拓展码个数需要和电话号码个数一致。
    extendNum   string 

    // 短信任务code，没有请先创建。
    taskCode   string 

    // 一个在taskcode下唯一的随机字符串，对于taskSign相同的请求平台认为是商家的同一次短信发送。
    taskSign   string 

}

func NewTaobaoJstSmsMessageDirectBatchsendRequest() *TaobaoJstSmsMessageDirectBatchsendRequest{
    return &TaobaoJstSmsMessageDirectBatchsendRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoJstSmsMessageDirectBatchsendRequest) GetApiMethodName() string {
    return "taobao.jst.sms.message.direct.batchsend"
}

func (r TaobaoJstSmsMessageDirectBatchsendRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoJstSmsMessageDirectBatchsendRequest) SetSignName(signName string) error {
    r.signName = signName
    r.Set("sign_name", signName)
    return nil
}

func (r TaobaoJstSmsMessageDirectBatchsendRequest) GetSignName() string {
    return r.signName
}

func (r *TaobaoJstSmsMessageDirectBatchsendRequest) SetUrl(url string) error {
    r.url = url
    r.Set("url", url)
    return nil
}

func (r TaobaoJstSmsMessageDirectBatchsendRequest) GetUrl() string {
    return r.url
}

func (r *TaobaoJstSmsMessageDirectBatchsendRequest) SetSmsTemplateCode(smsTemplateCode string) error {
    r.smsTemplateCode = smsTemplateCode
    r.Set("sms_template_code", smsTemplateCode)
    return nil
}

func (r TaobaoJstSmsMessageDirectBatchsendRequest) GetSmsTemplateCode() string {
    return r.smsTemplateCode
}

func (r *TaobaoJstSmsMessageDirectBatchsendRequest) SetRecNum(recNum string) error {
    r.recNum = recNum
    r.Set("rec_num", recNum)
    return nil
}

func (r TaobaoJstSmsMessageDirectBatchsendRequest) GetRecNum() string {
    return r.recNum
}

func (r *TaobaoJstSmsMessageDirectBatchsendRequest) SetSmsContent(smsContent string) error {
    r.smsContent = smsContent
    r.Set("sms_content", smsContent)
    return nil
}

func (r TaobaoJstSmsMessageDirectBatchsendRequest) GetSmsContent() string {
    return r.smsContent
}

func (r *TaobaoJstSmsMessageDirectBatchsendRequest) SetExtendNum(extendNum string) error {
    r.extendNum = extendNum
    r.Set("extend_num", extendNum)
    return nil
}

func (r TaobaoJstSmsMessageDirectBatchsendRequest) GetExtendNum() string {
    return r.extendNum
}

func (r *TaobaoJstSmsMessageDirectBatchsendRequest) SetTaskCode(taskCode string) error {
    r.taskCode = taskCode
    r.Set("task_code", taskCode)
    return nil
}

func (r TaobaoJstSmsMessageDirectBatchsendRequest) GetTaskCode() string {
    return r.taskCode
}

func (r *TaobaoJstSmsMessageDirectBatchsendRequest) SetTaskSign(taskSign string) error {
    r.taskSign = taskSign
    r.Set("task_sign", taskSign)
    return nil
}

func (r TaobaoJstSmsMessageDirectBatchsendRequest) GetTaskSign() string {
    return r.taskSign
}

