package tmallgeniescp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
二级物料-PO数据同步 API请求
alibaba.tmallgenie.scp.plan.current.rawpo.get

二级物料-PO数据同步（WO-W[TL])
*/
type AlibabaTmallgenieScpPlanCurrentRawpoGetRequest struct {
    model.Params
    // 系统自动生成
    paramRequest   *AbstractRequest
}

// 初始化AlibabaTmallgenieScpPlanCurrentRawpoGetRequest对象
func NewAlibabaTmallgenieScpPlanCurrentRawpoGetRequest() *AlibabaTmallgenieScpPlanCurrentRawpoGetRequest{
    return &AlibabaTmallgenieScpPlanCurrentRawpoGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaTmallgenieScpPlanCurrentRawpoGetRequest) GetApiMethodName() string {
    return "alibaba.tmallgenie.scp.plan.current.rawpo.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaTmallgenieScpPlanCurrentRawpoGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamRequest Setter
// 系统自动生成
func (r *AlibabaTmallgenieScpPlanCurrentRawpoGetRequest) SetParamRequest(paramRequest *AbstractRequest) error {
    r.paramRequest = paramRequest
    r.Set("param_request", paramRequest)
    return nil
}

// ParamRequest Getter
func (r AlibabaTmallgenieScpPlanCurrentRawpoGetRequest) GetParamRequest() *AbstractRequest {
    return r.paramRequest
}
