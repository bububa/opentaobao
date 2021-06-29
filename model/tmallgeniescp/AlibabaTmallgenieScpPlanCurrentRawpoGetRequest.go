package tmallgeniescp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
二级物料-PO数据同步 APIRequest
alibaba.tmallgenie.scp.plan.current.rawpo.get

二级物料-PO数据同步（WO-W[TL])
*/
type AlibabaTmallgenieScpPlanCurrentRawpoGetRequest struct {
    model.Params

    // 系统自动生成
    paramRequest   *AbstractRequest 

}

func NewAlibabaTmallgenieScpPlanCurrentRawpoGetRequest() *AlibabaTmallgenieScpPlanCurrentRawpoGetRequest{
    return &AlibabaTmallgenieScpPlanCurrentRawpoGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaTmallgenieScpPlanCurrentRawpoGetRequest) GetApiMethodName() string {
    return "alibaba.tmallgenie.scp.plan.current.rawpo.get"
}

func (r AlibabaTmallgenieScpPlanCurrentRawpoGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaTmallgenieScpPlanCurrentRawpoGetRequest) SetParamRequest(paramRequest *AbstractRequest) error {
    r.paramRequest = paramRequest
    r.Set("param_request", paramRequest)
    return nil
}

func (r AlibabaTmallgenieScpPlanCurrentRawpoGetRequest) GetParamRequest() *AbstractRequest {
    return r.paramRequest
}

