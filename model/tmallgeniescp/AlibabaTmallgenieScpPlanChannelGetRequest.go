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
type AlibabaTmallgenieScpPlanChannelGetAPIRequest struct {
    model.Params
    // 扩展参数
    _requestExtendJson   string
}

// 初始化AlibabaTmallgenieScpPlanChannelGetAPIRequest对象
func NewAlibabaTmallgenieScpPlanChannelGetRequest() *AlibabaTmallgenieScpPlanChannelGetAPIRequest{
    return &AlibabaTmallgenieScpPlanChannelGetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaTmallgenieScpPlanChannelGetAPIRequest) GetApiMethodName() string {
    return "alibaba.tmallgenie.scp.plan.channel.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaTmallgenieScpPlanChannelGetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// RequestExtendJson Setter
// 扩展参数
func (r *AlibabaTmallgenieScpPlanChannelGetAPIRequest) SetRequestExtendJson(_requestExtendJson string) error {
    r._requestExtendJson = _requestExtendJson
    r.Set("request_extend_json", _requestExtendJson)
    return nil
}

// RequestExtendJson Getter
func (r AlibabaTmallgenieScpPlanChannelGetAPIRequest) GetRequestExtendJson() string {
    return r._requestExtendJson
}
