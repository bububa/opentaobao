package alicom

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据号码tts单呼 API请求
alibaba.aliqin.ta.number.singlecallbytts

将语音验证码和语音通知发布至聚石塔渠道
*/
type AlibabaAliqinTaNumberSinglecallbyttsRequest struct {
    model.Params
    // 被叫号码
    calledNum   string
    // 显示号码
    calledShowNum   string
    // tts文本模板code
    ttsCode   string
    // 上下文参数,tts模板含有变量的, 此参数需填写。示例:{"date":"2015年 " ,"name":"测试","extend":"回传参数"} date、name 为模板里的变量名作为key,extend为扩展信息作为回传参数的key
    params   string
}

// 初始化AlibabaAliqinTaNumberSinglecallbyttsRequest对象
func NewAlibabaAliqinTaNumberSinglecallbyttsRequest() *AlibabaAliqinTaNumberSinglecallbyttsRequest{
    return &AlibabaAliqinTaNumberSinglecallbyttsRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAliqinTaNumberSinglecallbyttsRequest) GetApiMethodName() string {
    return "alibaba.aliqin.ta.number.singlecallbytts"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAliqinTaNumberSinglecallbyttsRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CalledNum Setter
// 被叫号码
func (r *AlibabaAliqinTaNumberSinglecallbyttsRequest) SetCalledNum(calledNum string) error {
    r.calledNum = calledNum
    r.Set("called_num", calledNum)
    return nil
}

// CalledNum Getter
func (r AlibabaAliqinTaNumberSinglecallbyttsRequest) GetCalledNum() string {
    return r.calledNum
}
// CalledShowNum Setter
// 显示号码
func (r *AlibabaAliqinTaNumberSinglecallbyttsRequest) SetCalledShowNum(calledShowNum string) error {
    r.calledShowNum = calledShowNum
    r.Set("called_show_num", calledShowNum)
    return nil
}

// CalledShowNum Getter
func (r AlibabaAliqinTaNumberSinglecallbyttsRequest) GetCalledShowNum() string {
    return r.calledShowNum
}
// TtsCode Setter
// tts文本模板code
func (r *AlibabaAliqinTaNumberSinglecallbyttsRequest) SetTtsCode(ttsCode string) error {
    r.ttsCode = ttsCode
    r.Set("tts_code", ttsCode)
    return nil
}

// TtsCode Getter
func (r AlibabaAliqinTaNumberSinglecallbyttsRequest) GetTtsCode() string {
    return r.ttsCode
}
// Params Setter
// 上下文参数,tts模板含有变量的, 此参数需填写。示例:{"date":"2015年 " ,"name":"测试","extend":"回传参数"} date、name 为模板里的变量名作为key,extend为扩展信息作为回传参数的key
func (r *AlibabaAliqinTaNumberSinglecallbyttsRequest) SetParams(params string) error {
    r.params = params
    r.Set("params", params)
    return nil
}

// Params Getter
func (r AlibabaAliqinTaNumberSinglecallbyttsRequest) GetParams() string {
    return r.params
}
