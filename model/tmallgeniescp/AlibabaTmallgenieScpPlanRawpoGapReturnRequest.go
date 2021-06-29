package tmallgeniescp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
二级物料-LT内的POGAP数据回传 APIRequest
alibaba.tmallgenie.scp.plan.rawpo.gap.return

二级物料-LT内的POGAP数据回传
*/
type AlibabaTmallgenieScpPlanRawpoGapReturnRequest struct {
    model.Params

    // 请求对象
    rawPogapRequest   *RawPurchaseOrderGapRequest 

}

func NewAlibabaTmallgenieScpPlanRawpoGapReturnRequest() *AlibabaTmallgenieScpPlanRawpoGapReturnRequest{
    return &AlibabaTmallgenieScpPlanRawpoGapReturnRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaTmallgenieScpPlanRawpoGapReturnRequest) GetApiMethodName() string {
    return "alibaba.tmallgenie.scp.plan.rawpo.gap.return"
}

func (r AlibabaTmallgenieScpPlanRawpoGapReturnRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaTmallgenieScpPlanRawpoGapReturnRequest) SetRawPogapRequest(rawPogapRequest *RawPurchaseOrderGapRequest) error {
    r.rawPogapRequest = rawPogapRequest
    r.Set("raw_pogap_request", rawPogapRequest)
    return nil
}

func (r AlibabaTmallgenieScpPlanRawpoGapReturnRequest) GetRawPogapRequest() *RawPurchaseOrderGapRequest {
    return r.rawPogapRequest
}

