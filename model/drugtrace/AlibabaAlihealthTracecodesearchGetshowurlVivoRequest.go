package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取药品扫码落地页vivo APIRequest
alibaba.alihealth.tracecodesearch.getshowurl.vivo

获取药品扫码落地页vivo
*/
type AlibabaAlihealthTracecodesearchGetshowurlVivoRequest struct {
    model.Params

    // 追溯码
    code   string 

    // 来源
    channel   string 

}

func NewAlibabaAlihealthTracecodesearchGetshowurlVivoRequest() *AlibabaAlihealthTracecodesearchGetshowurlVivoRequest{
    return &AlibabaAlihealthTracecodesearchGetshowurlVivoRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihealthTracecodesearchGetshowurlVivoRequest) GetApiMethodName() string {
    return "alibaba.alihealth.tracecodesearch.getshowurl.vivo"
}

func (r AlibabaAlihealthTracecodesearchGetshowurlVivoRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihealthTracecodesearchGetshowurlVivoRequest) SetCode(code string) error {
    r.code = code
    r.Set("code", code)
    return nil
}

func (r AlibabaAlihealthTracecodesearchGetshowurlVivoRequest) GetCode() string {
    return r.code
}

func (r *AlibabaAlihealthTracecodesearchGetshowurlVivoRequest) SetChannel(channel string) error {
    r.channel = channel
    r.Set("channel", channel)
    return nil
}

func (r AlibabaAlihealthTracecodesearchGetshowurlVivoRequest) GetChannel() string {
    return r.channel
}

