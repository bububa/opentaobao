package tmallgeniescp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
12-同步销售数据 APIRequest
alibaba.tmallgenie.scp.plan.sale.qty.get

同步销售数据
*/
type AlibabaTmallgenieScpPlanSaleQtyGetRequest struct {
    model.Params

    // 扩展参数
    requestExtendJson   string 

}

func NewAlibabaTmallgenieScpPlanSaleQtyGetRequest() *AlibabaTmallgenieScpPlanSaleQtyGetRequest{
    return &AlibabaTmallgenieScpPlanSaleQtyGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaTmallgenieScpPlanSaleQtyGetRequest) GetApiMethodName() string {
    return "alibaba.tmallgenie.scp.plan.sale.qty.get"
}

func (r AlibabaTmallgenieScpPlanSaleQtyGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaTmallgenieScpPlanSaleQtyGetRequest) SetRequestExtendJson(requestExtendJson string) error {
    r.requestExtendJson = requestExtendJson
    r.Set("request_extend_json", requestExtendJson)
    return nil
}

func (r AlibabaTmallgenieScpPlanSaleQtyGetRequest) GetRequestExtendJson() string {
    return r.requestExtendJson
}

