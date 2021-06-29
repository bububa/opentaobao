package tmallgeniescp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
同步销售数据按照渠道类型汇总 APIRequest
alibaba.tmallgenie.scp.plan.summary.sale.qty.get

同步销售数据按照渠道类型汇总
*/
type AlibabaTmallgenieScpPlanSummarySaleQtyGetRequest struct {
    model.Params

    // 扩展参数
    requestExtendJson   string 

}

func NewAlibabaTmallgenieScpPlanSummarySaleQtyGetRequest() *AlibabaTmallgenieScpPlanSummarySaleQtyGetRequest{
    return &AlibabaTmallgenieScpPlanSummarySaleQtyGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaTmallgenieScpPlanSummarySaleQtyGetRequest) GetApiMethodName() string {
    return "alibaba.tmallgenie.scp.plan.summary.sale.qty.get"
}

func (r AlibabaTmallgenieScpPlanSummarySaleQtyGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaTmallgenieScpPlanSummarySaleQtyGetRequest) SetRequestExtendJson(requestExtendJson string) error {
    r.requestExtendJson = requestExtendJson
    r.Set("request_extend_json", requestExtendJson)
    return nil
}

func (r AlibabaTmallgenieScpPlanSummarySaleQtyGetRequest) GetRequestExtendJson() string {
    return r.requestExtendJson
}

