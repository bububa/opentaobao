package tmallgeniescp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
【已废除】同步历史的销售数据 APIRequest
alibaba.tmallgenie.scp.plan.history.sale.qty.get

同步历史的销售数据
*/
type AlibabaTmallgenieScpPlanHistorySaleQtyGetRequest struct {
    model.Params

    // 扩展参数
    requestExtendJson   string 

}

func NewAlibabaTmallgenieScpPlanHistorySaleQtyGetRequest() *AlibabaTmallgenieScpPlanHistorySaleQtyGetRequest{
    return &AlibabaTmallgenieScpPlanHistorySaleQtyGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaTmallgenieScpPlanHistorySaleQtyGetRequest) GetApiMethodName() string {
    return "alibaba.tmallgenie.scp.plan.history.sale.qty.get"
}

func (r AlibabaTmallgenieScpPlanHistorySaleQtyGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaTmallgenieScpPlanHistorySaleQtyGetRequest) SetRequestExtendJson(requestExtendJson string) error {
    r.requestExtendJson = requestExtendJson
    r.Set("request_extend_json", requestExtendJson)
    return nil
}

func (r AlibabaTmallgenieScpPlanHistorySaleQtyGetRequest) GetRequestExtendJson() string {
    return r.requestExtendJson
}

