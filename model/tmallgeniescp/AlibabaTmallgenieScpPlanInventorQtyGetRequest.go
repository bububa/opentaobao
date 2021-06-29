package tmallgeniescp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
10-同步库存现有量 API请求
alibaba.tmallgenie.scp.plan.inventor.qty.get

同步库存现有量
*/
type AlibabaTmallgenieScpPlanInventorQtyGetRequest struct {
    model.Params
    // 扩展参数
    requestExtendJson   string
}

// 初始化AlibabaTmallgenieScpPlanInventorQtyGetRequest对象
func NewAlibabaTmallgenieScpPlanInventorQtyGetRequest() *AlibabaTmallgenieScpPlanInventorQtyGetRequest{
    return &AlibabaTmallgenieScpPlanInventorQtyGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaTmallgenieScpPlanInventorQtyGetRequest) GetApiMethodName() string {
    return "alibaba.tmallgenie.scp.plan.inventor.qty.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaTmallgenieScpPlanInventorQtyGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// RequestExtendJson Setter
// 扩展参数
func (r *AlibabaTmallgenieScpPlanInventorQtyGetRequest) SetRequestExtendJson(requestExtendJson string) error {
    r.requestExtendJson = requestExtendJson
    r.Set("request_extend_json", requestExtendJson)
    return nil
}

// RequestExtendJson Getter
func (r AlibabaTmallgenieScpPlanInventorQtyGetRequest) GetRequestExtendJson() string {
    return r.requestExtendJson
}
