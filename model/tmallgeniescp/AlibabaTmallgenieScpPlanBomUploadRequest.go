package tmallgeniescp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
计划BOM同步 APIRequest
alibaba.tmallgenie.scp.plan.bom.upload

计划BOM同步
*/
type AlibabaTmallgenieScpPlanBomUploadRequest struct {
    model.Params

    // 对象
    pbomRequest   *AbstractRequest 

}

func NewAlibabaTmallgenieScpPlanBomUploadRequest() *AlibabaTmallgenieScpPlanBomUploadRequest{
    return &AlibabaTmallgenieScpPlanBomUploadRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaTmallgenieScpPlanBomUploadRequest) GetApiMethodName() string {
    return "alibaba.tmallgenie.scp.plan.bom.upload"
}

func (r AlibabaTmallgenieScpPlanBomUploadRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaTmallgenieScpPlanBomUploadRequest) SetPbomRequest(pbomRequest *AbstractRequest) error {
    r.pbomRequest = pbomRequest
    r.Set("pbom_request", pbomRequest)
    return nil
}

func (r AlibabaTmallgenieScpPlanBomUploadRequest) GetPbomRequest() *AbstractRequest {
    return r.pbomRequest
}

