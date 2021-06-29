package tmallgeniescp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
10-同步库存现有量 APIRequest
alibaba.tmallgenie.scp.plan.inventor.qty.get

同步库存现有量
*/
type AlibabaTmallgenieScpPlanInventorQtyGetRequest struct {
    model.Params

    // 扩展参数
    requestExtendJson   string 

}

func NewAlibabaTmallgenieScpPlanInventorQtyGetRequest() *AlibabaTmallgenieScpPlanInventorQtyGetRequest{
    return &AlibabaTmallgenieScpPlanInventorQtyGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaTmallgenieScpPlanInventorQtyGetRequest) GetApiMethodName() string {
    return "alibaba.tmallgenie.scp.plan.inventor.qty.get"
}

func (r AlibabaTmallgenieScpPlanInventorQtyGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaTmallgenieScpPlanInventorQtyGetRequest) SetRequestExtendJson(requestExtendJson string) error {
    r.requestExtendJson = requestExtendJson
    r.Set("request_extend_json", requestExtendJson)
    return nil
}

func (r AlibabaTmallgenieScpPlanInventorQtyGetRequest) GetRequestExtendJson() string {
    return r.requestExtendJson
}

