package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取vivo banner API请求
alibaba.alihealth.tracecodesearc.getinfomation.vivo

获取vivo banner  url
*/
type AlibabaAlihealthTracecodesearcGetinfomationVivoRequest struct {
    model.Params
    // 渠道
    channel   string
}

// 初始化AlibabaAlihealthTracecodesearcGetinfomationVivoRequest对象
func NewAlibabaAlihealthTracecodesearcGetinfomationVivoRequest() *AlibabaAlihealthTracecodesearcGetinfomationVivoRequest{
    return &AlibabaAlihealthTracecodesearcGetinfomationVivoRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthTracecodesearcGetinfomationVivoRequest) GetApiMethodName() string {
    return "alibaba.alihealth.tracecodesearc.getinfomation.vivo"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthTracecodesearcGetinfomationVivoRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Channel Setter
// 渠道
func (r *AlibabaAlihealthTracecodesearcGetinfomationVivoRequest) SetChannel(channel string) error {
    r.channel = channel
    r.Set("channel", channel)
    return nil
}

// Channel Getter
func (r AlibabaAlihealthTracecodesearcGetinfomationVivoRequest) GetChannel() string {
    return r.channel
}
