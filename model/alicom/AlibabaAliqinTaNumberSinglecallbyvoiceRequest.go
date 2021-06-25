package alicom

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据号码tts单呼 APIRequest
alibaba.aliqin.ta.number.singlecallbyvoice

根据号码语音单呼
*/
type AlibabaAliqinTaNumberSinglecallbyvoiceRequest struct {
    model.Params

    // 单呼号码
    calledNum   string 

    // 显示号码
    calledShowNum   string 

    // 语音文件code
    voiceCode   string 

    // 上下文参数 示例:{"extend":"回传参数"} extend为扩展信息作为回传参数的key
    params   string 

}

func NewAlibabaAliqinTaNumberSinglecallbyvoiceRequest() *AlibabaAliqinTaNumberSinglecallbyvoiceRequest{
    return &AlibabaAliqinTaNumberSinglecallbyvoiceRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAliqinTaNumberSinglecallbyvoiceRequest) GetApiMethodName() string {
    return "alibaba.aliqin.ta.number.singlecallbyvoice"
}

func (r AlibabaAliqinTaNumberSinglecallbyvoiceRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAliqinTaNumberSinglecallbyvoiceRequest) SetCalledNum(calledNum string) error {
    r.calledNum = calledNum
    r.Set("called_num", calledNum)
    return nil
}

func (r AlibabaAliqinTaNumberSinglecallbyvoiceRequest) GetCalledNum() string {
    return r.calledNum
}

func (r *AlibabaAliqinTaNumberSinglecallbyvoiceRequest) SetCalledShowNum(calledShowNum string) error {
    r.calledShowNum = calledShowNum
    r.Set("called_show_num", calledShowNum)
    return nil
}

func (r AlibabaAliqinTaNumberSinglecallbyvoiceRequest) GetCalledShowNum() string {
    return r.calledShowNum
}

func (r *AlibabaAliqinTaNumberSinglecallbyvoiceRequest) SetVoiceCode(voiceCode string) error {
    r.voiceCode = voiceCode
    r.Set("voice_code", voiceCode)
    return nil
}

func (r AlibabaAliqinTaNumberSinglecallbyvoiceRequest) GetVoiceCode() string {
    return r.voiceCode
}

func (r *AlibabaAliqinTaNumberSinglecallbyvoiceRequest) SetParams(params string) error {
    r.params = params
    r.Set("params", params)
    return nil
}

func (r AlibabaAliqinTaNumberSinglecallbyvoiceRequest) GetParams() string {
    return r.params
}

