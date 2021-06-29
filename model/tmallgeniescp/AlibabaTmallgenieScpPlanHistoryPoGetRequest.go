package tmallgeniescp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
【已废除】11-同步历史所有的po单 APIRequest
alibaba.tmallgenie.scp.plan.history.po.get

同步历史po单
*/
type AlibabaTmallgenieScpPlanHistoryPoGetRequest struct {
    model.Params

    // 扩展参数
    requestExtendJson   string 

}

func NewAlibabaTmallgenieScpPlanHistoryPoGetRequest() *AlibabaTmallgenieScpPlanHistoryPoGetRequest{
    return &AlibabaTmallgenieScpPlanHistoryPoGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaTmallgenieScpPlanHistoryPoGetRequest) GetApiMethodName() string {
    return "alibaba.tmallgenie.scp.plan.history.po.get"
}

func (r AlibabaTmallgenieScpPlanHistoryPoGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaTmallgenieScpPlanHistoryPoGetRequest) SetRequestExtendJson(requestExtendJson string) error {
    r.requestExtendJson = requestExtendJson
    r.Set("request_extend_json", requestExtendJson)
    return nil
}

func (r AlibabaTmallgenieScpPlanHistoryPoGetRequest) GetRequestExtendJson() string {
    return r.requestExtendJson
}

