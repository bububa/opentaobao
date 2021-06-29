package tmallgeniescp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
11-同步本周的po单（从W-1周到W+4周） APIRequest
alibaba.tmallgenie.scp.plan.current.po.get

11-同步本周的po单（从W-1周到W+4周）
*/
type AlibabaTmallgenieScpPlanCurrentPoGetRequest struct {
    model.Params

    // 扩展参数
    requestExtendJson   string 

}

func NewAlibabaTmallgenieScpPlanCurrentPoGetRequest() *AlibabaTmallgenieScpPlanCurrentPoGetRequest{
    return &AlibabaTmallgenieScpPlanCurrentPoGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaTmallgenieScpPlanCurrentPoGetRequest) GetApiMethodName() string {
    return "alibaba.tmallgenie.scp.plan.current.po.get"
}

func (r AlibabaTmallgenieScpPlanCurrentPoGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaTmallgenieScpPlanCurrentPoGetRequest) SetRequestExtendJson(requestExtendJson string) error {
    r.requestExtendJson = requestExtendJson
    r.Set("request_extend_json", requestExtendJson)
    return nil
}

func (r AlibabaTmallgenieScpPlanCurrentPoGetRequest) GetRequestExtendJson() string {
    return r.requestExtendJson
}

