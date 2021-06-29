package tmallgeniescp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
5-IBP同步渠道接口 API请求
alibaba.tmallgenie.scp.plan.channel.get

IBP同步渠道接口
*/
type AlibabaTmallgenieScpPlanChannelGetRequest struct {
    model.Params
    // 扩展参数
    requestExtendJson   string
}

// 初始化AlibabaTmallgenieScpPlanChannelGetRequest对象
func NewAlibabaTmallgenieScpPlanChannelGetRequest() *AlibabaTmallgenieScpPlanChannelGetRequest{
    return &AlibabaTmallgenieScpPlanChannelGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaTmallgenieScpPlanChannelGetRequest) GetApiMethodName() string {
    return "alibaba.tmallgenie.scp.plan.channel.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaTmallgenieScpPlanChannelGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// RequestExtendJson Setter
// 扩展参数
func (r *AlibabaTmallgenieScpPlanChannelGetRequest) SetRequestExtendJson(requestExtendJson string) error {
    r.requestExtendJson = requestExtendJson
    r.Set("request_extend_json", requestExtendJson)
    return nil
}

// RequestExtendJson Getter
func (r AlibabaTmallgenieScpPlanChannelGetRequest) GetRequestExtendJson() string {
    return r.requestExtendJson
}
