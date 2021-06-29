package tmallgeniescp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
12-同步销售数据 API请求
alibaba.tmallgenie.scp.plan.sale.qty.get

同步销售数据
*/
type AlibabaTmallgenieScpPlanSaleQtyGetRequest struct {
    model.Params
    // 扩展参数
    _requestExtendJson   string
}

// 初始化AlibabaTmallgenieScpPlanSaleQtyGetRequest对象
func NewAlibabaTmallgenieScpPlanSaleQtyGetRequest() *AlibabaTmallgenieScpPlanSaleQtyGetRequest{
    return &AlibabaTmallgenieScpPlanSaleQtyGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaTmallgenieScpPlanSaleQtyGetRequest) GetApiMethodName() string {
    return "alibaba.tmallgenie.scp.plan.sale.qty.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaTmallgenieScpPlanSaleQtyGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// RequestExtendJson Setter
// 扩展参数
func (r *AlibabaTmallgenieScpPlanSaleQtyGetRequest) SetRequestExtendJson(_requestExtendJson string) error {
    r._requestExtendJson = _requestExtendJson
    r.Set("request_extend_json", _requestExtendJson)
    return nil
}

// RequestExtendJson Getter
func (r AlibabaTmallgenieScpPlanSaleQtyGetRequest) GetRequestExtendJson() string {
    return r._requestExtendJson
}
