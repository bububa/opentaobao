package tmallgeniescp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
23-Net Demand净需求回传 API请求
alibaba.tmallgenie.scp.plan.netdemand.upload

Net Demand净需求回传
*/
type AlibabaTmallgenieScpPlanNetdemandUploadRequest struct {
    model.Params
    // 对象
    netDemandRequest   *NetDemandRequest
}

// 初始化AlibabaTmallgenieScpPlanNetdemandUploadRequest对象
func NewAlibabaTmallgenieScpPlanNetdemandUploadRequest() *AlibabaTmallgenieScpPlanNetdemandUploadRequest{
    return &AlibabaTmallgenieScpPlanNetdemandUploadRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaTmallgenieScpPlanNetdemandUploadRequest) GetApiMethodName() string {
    return "alibaba.tmallgenie.scp.plan.netdemand.upload"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaTmallgenieScpPlanNetdemandUploadRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// NetDemandRequest Setter
// 对象
func (r *AlibabaTmallgenieScpPlanNetdemandUploadRequest) SetNetDemandRequest(netDemandRequest *NetDemandRequest) error {
    r.netDemandRequest = netDemandRequest
    r.Set("net_demand_request", netDemandRequest)
    return nil
}

// NetDemandRequest Getter
func (r AlibabaTmallgenieScpPlanNetdemandUploadRequest) GetNetDemandRequest() *NetDemandRequest {
    return r.netDemandRequest
}
