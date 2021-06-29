package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取药品扫码落地页vivo API请求
alibaba.alihealth.tracecodesearch.getshowurl.vivo

获取药品扫码落地页vivo
*/
type AlibabaAlihealthTracecodesearchGetshowurlVivoRequest struct {
    model.Params
    // 追溯码
    _code   string
    // 来源
    _channel   string
}

// 初始化AlibabaAlihealthTracecodesearchGetshowurlVivoRequest对象
func NewAlibabaAlihealthTracecodesearchGetshowurlVivoRequest() *AlibabaAlihealthTracecodesearchGetshowurlVivoRequest{
    return &AlibabaAlihealthTracecodesearchGetshowurlVivoRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthTracecodesearchGetshowurlVivoRequest) GetApiMethodName() string {
    return "alibaba.alihealth.tracecodesearch.getshowurl.vivo"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthTracecodesearchGetshowurlVivoRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Code Setter
// 追溯码
func (r *AlibabaAlihealthTracecodesearchGetshowurlVivoRequest) SetCode(_code string) error {
    r._code = _code
    r.Set("code", _code)
    return nil
}

// Code Getter
func (r AlibabaAlihealthTracecodesearchGetshowurlVivoRequest) GetCode() string {
    return r._code
}
// Channel Setter
// 来源
func (r *AlibabaAlihealthTracecodesearchGetshowurlVivoRequest) SetChannel(_channel string) error {
    r._channel = _channel
    r.Set("channel", _channel)
    return nil
}

// Channel Getter
func (r AlibabaAlihealthTracecodesearchGetshowurlVivoRequest) GetChannel() string {
    return r._channel
}
