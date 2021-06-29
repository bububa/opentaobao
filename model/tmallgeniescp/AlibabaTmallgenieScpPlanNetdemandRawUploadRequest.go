package tmallgeniescp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
二级物料净需求回传(TL+1) API请求
alibaba.tmallgenie.scp.plan.netdemand.raw.upload

二级物料净需求回传(TL+1)
*/
type AlibabaTmallgenieScpPlanNetdemandRawUploadRequest struct {
    model.Params
    // 对象
    _netDemandRawRequest   *NetDemandRawRequest
}

// 初始化AlibabaTmallgenieScpPlanNetdemandRawUploadRequest对象
func NewAlibabaTmallgenieScpPlanNetdemandRawUploadRequest() *AlibabaTmallgenieScpPlanNetdemandRawUploadRequest{
    return &AlibabaTmallgenieScpPlanNetdemandRawUploadRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaTmallgenieScpPlanNetdemandRawUploadRequest) GetApiMethodName() string {
    return "alibaba.tmallgenie.scp.plan.netdemand.raw.upload"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaTmallgenieScpPlanNetdemandRawUploadRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// NetDemandRawRequest Setter
// 对象
func (r *AlibabaTmallgenieScpPlanNetdemandRawUploadRequest) SetNetDemandRawRequest(_netDemandRawRequest *NetDemandRawRequest) error {
    r._netDemandRawRequest = _netDemandRawRequest
    r.Set("net_demand_raw_request", _netDemandRawRequest)
    return nil
}

// NetDemandRawRequest Getter
func (r AlibabaTmallgenieScpPlanNetdemandRawUploadRequest) GetNetDemandRawRequest() *NetDemandRawRequest {
    return r._netDemandRawRequest
}