package tmallgeniescp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
【已废除】11-同步历史所有的po单 API请求
alibaba.tmallgenie.scp.plan.history.po.get

同步历史po单
*/
type AlibabaTmallgenieScpPlanHistoryPoGetRequest struct {
    model.Params
    // 扩展参数
    _requestExtendJson   string
}

// 初始化AlibabaTmallgenieScpPlanHistoryPoGetRequest对象
func NewAlibabaTmallgenieScpPlanHistoryPoGetRequest() *AlibabaTmallgenieScpPlanHistoryPoGetRequest{
    return &AlibabaTmallgenieScpPlanHistoryPoGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaTmallgenieScpPlanHistoryPoGetRequest) GetApiMethodName() string {
    return "alibaba.tmallgenie.scp.plan.history.po.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaTmallgenieScpPlanHistoryPoGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// RequestExtendJson Setter
// 扩展参数
func (r *AlibabaTmallgenieScpPlanHistoryPoGetRequest) SetRequestExtendJson(_requestExtendJson string) error {
    r._requestExtendJson = _requestExtendJson
    r.Set("request_extend_json", _requestExtendJson)
    return nil
}

// RequestExtendJson Getter
func (r AlibabaTmallgenieScpPlanHistoryPoGetRequest) GetRequestExtendJson() string {
    return r._requestExtendJson
}
