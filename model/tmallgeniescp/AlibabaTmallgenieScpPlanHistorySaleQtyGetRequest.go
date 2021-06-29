package tmallgeniescp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
【已废除】同步历史的销售数据 API请求
alibaba.tmallgenie.scp.plan.history.sale.qty.get

同步历史的销售数据
*/
type AlibabaTmallgenieScpPlanHistorySaleQtyGetRequest struct {
    model.Params
    // 扩展参数
    _requestExtendJson   string
}

// 初始化AlibabaTmallgenieScpPlanHistorySaleQtyGetRequest对象
func NewAlibabaTmallgenieScpPlanHistorySaleQtyGetRequest() *AlibabaTmallgenieScpPlanHistorySaleQtyGetRequest{
    return &AlibabaTmallgenieScpPlanHistorySaleQtyGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaTmallgenieScpPlanHistorySaleQtyGetRequest) GetApiMethodName() string {
    return "alibaba.tmallgenie.scp.plan.history.sale.qty.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaTmallgenieScpPlanHistorySaleQtyGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// RequestExtendJson Setter
// 扩展参数
func (r *AlibabaTmallgenieScpPlanHistorySaleQtyGetRequest) SetRequestExtendJson(_requestExtendJson string) error {
    r._requestExtendJson = _requestExtendJson
    r.Set("request_extend_json", _requestExtendJson)
    return nil
}

// RequestExtendJson Getter
func (r AlibabaTmallgenieScpPlanHistorySaleQtyGetRequest) GetRequestExtendJson() string {
    return r._requestExtendJson
}
