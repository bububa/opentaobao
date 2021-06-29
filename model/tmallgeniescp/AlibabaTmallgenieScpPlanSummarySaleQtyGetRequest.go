package tmallgeniescp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
同步销售数据按照渠道类型汇总 API请求
alibaba.tmallgenie.scp.plan.summary.sale.qty.get

同步销售数据按照渠道类型汇总
*/
type AlibabaTmallgenieScpPlanSummarySaleQtyGetRequest struct {
    model.Params
    // 扩展参数
    _requestExtendJson   string
}

// 初始化AlibabaTmallgenieScpPlanSummarySaleQtyGetRequest对象
func NewAlibabaTmallgenieScpPlanSummarySaleQtyGetRequest() *AlibabaTmallgenieScpPlanSummarySaleQtyGetRequest{
    return &AlibabaTmallgenieScpPlanSummarySaleQtyGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaTmallgenieScpPlanSummarySaleQtyGetRequest) GetApiMethodName() string {
    return "alibaba.tmallgenie.scp.plan.summary.sale.qty.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaTmallgenieScpPlanSummarySaleQtyGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// RequestExtendJson Setter
// 扩展参数
func (r *AlibabaTmallgenieScpPlanSummarySaleQtyGetRequest) SetRequestExtendJson(_requestExtendJson string) error {
    r._requestExtendJson = _requestExtendJson
    r.Set("request_extend_json", _requestExtendJson)
    return nil
}

// RequestExtendJson Getter
func (r AlibabaTmallgenieScpPlanSummarySaleQtyGetRequest) GetRequestExtendJson() string {
    return r._requestExtendJson
}
