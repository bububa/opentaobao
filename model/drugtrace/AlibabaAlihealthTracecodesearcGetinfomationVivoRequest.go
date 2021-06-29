package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取vivo banner APIRequest
alibaba.alihealth.tracecodesearc.getinfomation.vivo

获取vivo banner  url
*/
type AlibabaAlihealthTracecodesearcGetinfomationVivoRequest struct {
    model.Params

    // 渠道
    channel   string 

}

func NewAlibabaAlihealthTracecodesearcGetinfomationVivoRequest() *AlibabaAlihealthTracecodesearcGetinfomationVivoRequest{
    return &AlibabaAlihealthTracecodesearcGetinfomationVivoRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihealthTracecodesearcGetinfomationVivoRequest) GetApiMethodName() string {
    return "alibaba.alihealth.tracecodesearc.getinfomation.vivo"
}

func (r AlibabaAlihealthTracecodesearcGetinfomationVivoRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihealthTracecodesearcGetinfomationVivoRequest) SetChannel(channel string) error {
    r.channel = channel
    r.Set("channel", channel)
    return nil
}

func (r AlibabaAlihealthTracecodesearcGetinfomationVivoRequest) GetChannel() string {
    return r.channel
}

