package tmallgeniescp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
5-IBP同步渠道接口 APIRequest
alibaba.tmallgenie.scp.plan.channel.get

IBP同步渠道接口
*/
type AlibabaTmallgenieScpPlanChannelGetRequest struct {
    model.Params

    // 扩展参数
    requestExtendJson   string 

}

func NewAlibabaTmallgenieScpPlanChannelGetRequest() *AlibabaTmallgenieScpPlanChannelGetRequest{
    return &AlibabaTmallgenieScpPlanChannelGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaTmallgenieScpPlanChannelGetRequest) GetApiMethodName() string {
    return "alibaba.tmallgenie.scp.plan.channel.get"
}

func (r AlibabaTmallgenieScpPlanChannelGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaTmallgenieScpPlanChannelGetRequest) SetRequestExtendJson(requestExtendJson string) error {
    r.requestExtendJson = requestExtendJson
    r.Set("request_extend_json", requestExtendJson)
    return nil
}

func (r AlibabaTmallgenieScpPlanChannelGetRequest) GetRequestExtendJson() string {
    return r.requestExtendJson
}

