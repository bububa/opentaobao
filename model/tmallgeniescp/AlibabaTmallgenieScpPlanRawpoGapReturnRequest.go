package tmallgeniescp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
二级物料-LT内的POGAP数据回传 API请求
alibaba.tmallgenie.scp.plan.rawpo.gap.return

二级物料-LT内的POGAP数据回传
*/
type AlibabaTmallgenieScpPlanRawpoGapReturnRequest struct {
    model.Params
    // 请求对象
    _rawPogapRequest   *RawPurchaseOrderGapRequest
}

// 初始化AlibabaTmallgenieScpPlanRawpoGapReturnRequest对象
func NewAlibabaTmallgenieScpPlanRawpoGapReturnRequest() *AlibabaTmallgenieScpPlanRawpoGapReturnRequest{
    return &AlibabaTmallgenieScpPlanRawpoGapReturnRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaTmallgenieScpPlanRawpoGapReturnRequest) GetApiMethodName() string {
    return "alibaba.tmallgenie.scp.plan.rawpo.gap.return"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaTmallgenieScpPlanRawpoGapReturnRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// RawPogapRequest Setter
// 请求对象
func (r *AlibabaTmallgenieScpPlanRawpoGapReturnRequest) SetRawPogapRequest(_rawPogapRequest *RawPurchaseOrderGapRequest) error {
    r._rawPogapRequest = _rawPogapRequest
    r.Set("raw_pogap_request", _rawPogapRequest)
    return nil
}

// RawPogapRequest Getter
func (r AlibabaTmallgenieScpPlanRawpoGapReturnRequest) GetRawPogapRequest() *RawPurchaseOrderGapRequest {
    return r._rawPogapRequest
}
