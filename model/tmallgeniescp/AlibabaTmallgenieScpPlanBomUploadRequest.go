package tmallgeniescp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
计划BOM同步 API请求
alibaba.tmallgenie.scp.plan.bom.upload

计划BOM同步
*/
type AlibabaTmallgenieScpPlanBomUploadRequest struct {
    model.Params
    // 对象
    pbomRequest   *AbstractRequest
}

// 初始化AlibabaTmallgenieScpPlanBomUploadRequest对象
func NewAlibabaTmallgenieScpPlanBomUploadRequest() *AlibabaTmallgenieScpPlanBomUploadRequest{
    return &AlibabaTmallgenieScpPlanBomUploadRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaTmallgenieScpPlanBomUploadRequest) GetApiMethodName() string {
    return "alibaba.tmallgenie.scp.plan.bom.upload"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaTmallgenieScpPlanBomUploadRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// PbomRequest Setter
// 对象
func (r *AlibabaTmallgenieScpPlanBomUploadRequest) SetPbomRequest(pbomRequest *AbstractRequest) error {
    r.pbomRequest = pbomRequest
    r.Set("pbom_request", pbomRequest)
    return nil
}

// PbomRequest Getter
func (r AlibabaTmallgenieScpPlanBomUploadRequest) GetPbomRequest() *AbstractRequest {
    return r.pbomRequest
}
