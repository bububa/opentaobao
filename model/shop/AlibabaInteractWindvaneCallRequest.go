package shop

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
Weex页面调用windvane APIRequest
alibaba.interact.windvane.call

客户端鉴权使用，实际不会发送或接收数据
*/
type AlibabaInteractWindvaneCallRequest struct {
    model.Params

    // 客户端鉴权使用，实际不会发送或接收数据
    unNamed   string 

}

func NewAlibabaInteractWindvaneCallRequest() *AlibabaInteractWindvaneCallRequest{
    return &AlibabaInteractWindvaneCallRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaInteractWindvaneCallRequest) GetApiMethodName() string {
    return "alibaba.interact.windvane.call"
}

func (r AlibabaInteractWindvaneCallRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaInteractWindvaneCallRequest) SetUnNamed(unNamed string) error {
    r.unNamed = unNamed
    r.Set("un_named", unNamed)
    return nil
}

func (r AlibabaInteractWindvaneCallRequest) GetUnNamed() string {
    return r.unNamed
}

