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
type AlibabaTmallgenieScpPlanInventorQtyGetAPIRequest struct {
    model.Params
    // 扩展参数
    _requestExtendJson   string
}

// 初始化AlibabaTmallgenieScpPlanInventorQtyGetAPIRequest对象
func NewAlibabaTmallgenieScpPlanInventorQtyGetRequest() *AlibabaTmallgenieScpPlanInventorQtyGetAPIRequest{
    return &AlibabaTmallgenieScpPlanInventorQtyGetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaTmallgenieScpPlanInventorQtyGetAPIRequest) GetApiMethodName() string {
    return "alibaba.tmallgenie.scp.plan.inventor.qty.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaTmallgenieScpPlanInventorQtyGetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// RequestExtendJson Setter
// 扩展参数
func (r *AlibabaTmallgenieScpPlanInventorQtyGetAPIRequest) SetRequestExtendJson(_requestExtendJson string) error {
    r._requestExtendJson = _requestExtendJson
    r.Set("request_extend_json", _requestExtendJson)
    return nil
}

// RequestExtendJson Getter
func (r AlibabaTmallgenieScpPlanInventorQtyGetAPIRequest) GetRequestExtendJson() string {
    return r._requestExtendJson
}
