package tmallgeniescp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
1-IBP同步物料接口 APIRequest
alibaba.tmallgenie.scp.plan.materiel.get

IBP同步物料接口
*/
type AlibabaTmallgenieScpPlanMaterielGetRequest struct {
    model.Params

    // 扩展字段
    requestExtendJson   string 

}

func NewAlibabaTmallgenieScpPlanMaterielGetRequest() *AlibabaTmallgenieScpPlanMaterielGetRequest{
    return &AlibabaTmallgenieScpPlanMaterielGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaTmallgenieScpPlanMaterielGetRequest) GetApiMethodName() string {
    return "alibaba.tmallgenie.scp.plan.materiel.get"
}

func (r AlibabaTmallgenieScpPlanMaterielGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaTmallgenieScpPlanMaterielGetRequest) SetRequestExtendJson(requestExtendJson string) error {
    r.requestExtendJson = requestExtendJson
    r.Set("request_extend_json", requestExtendJson)
    return nil
}

func (r AlibabaTmallgenieScpPlanMaterielGetRequest) GetRequestExtendJson() string {
    return r.requestExtendJson
}

