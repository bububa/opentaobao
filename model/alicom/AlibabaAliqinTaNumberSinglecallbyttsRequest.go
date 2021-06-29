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
    _calledNum   string
    // 显示号码
    _calledShowNum   string
    // tts文本模板code
    _ttsCode   string
    // 上下文参数,tts模板含有变量的, 此参数需填写。示例:{"date":"2015年 " ,"name":"测试","extend":"回传参数"} date、name 为模板里的变量名作为key,extend为扩展信息作为回传参数的key
    _params   string
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
func (r *AlibabaAliqinTaNumberSinglecallbyttsRequest) SetCalledNum(_calledNum string) error {
    r._calledNum = _calledNum
    r.Set("called_num", _calledNum)
    return nil
}

// CalledNum Getter
func (r AlibabaAliqinTaNumberSinglecallbyttsRequest) GetCalledNum() string {
    return r._calledNum
}
// CalledShowNum Setter
// 显示号码
func (r *AlibabaAliqinTaNumberSinglecallbyttsRequest) SetCalledShowNum(_calledShowNum string) error {
    r._calledShowNum = _calledShowNum
    r.Set("called_show_num", _calledShowNum)
    return nil
}

// CalledShowNum Getter
func (r AlibabaAliqinTaNumberSinglecallbyttsRequest) GetCalledShowNum() string {
    return r._calledShowNum
}
// TtsCode Setter
// tts文本模板code
func (r *AlibabaAliqinTaNumberSinglecallbyttsRequest) SetTtsCode(_ttsCode string) error {
    r._ttsCode = _ttsCode
    r.Set("tts_code", _ttsCode)
    return nil
}

// TtsCode Getter
func (r AlibabaAliqinTaNumberSinglecallbyttsRequest) GetTtsCode() string {
    return r._ttsCode
}
// Params Setter
// 上下文参数,tts模板含有变量的, 此参数需填写。示例:{"date":"2015年 " ,"name":"测试","extend":"回传参数"} date、name 为模板里的变量名作为key,extend为扩展信息作为回传参数的key
func (r *AlibabaAliqinTaNumberSinglecallbyttsRequest) SetParams(_params string) error {
    r._params = _params
    r.Set("params", _params)
    return nil
}

// Params Getter
func (r AlibabaAliqinTaNumberSinglecallbyttsRequest) GetParams() string {
    return r._params
}
