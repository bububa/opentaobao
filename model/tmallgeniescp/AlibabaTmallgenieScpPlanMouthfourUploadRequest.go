package tmallgeniescp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
21-M+4PR 回传接口接口 API请求
alibaba.tmallgenie.scp.plan.mouthfour.upload

M+4 PR 回传接口
*/
type AlibabaTmallgenieScpPlanMouthfourUploadAPIRequest struct {
    model.Params
    // 请求参数
    _monthFourPrRequest   *MonthFourPrRequest
}

// 初始化AlibabaTmallgenieScpPlanMouthfourUploadAPIRequest对象
func NewAlibabaTmallgenieScpPlanMouthfourUploadRequest() *AlibabaTmallgenieScpPlanMouthfourUploadAPIRequest{
    return &AlibabaTmallgenieScpPlanMouthfourUploadAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaTmallgenieScpPlanMouthfourUploadAPIRequest) GetApiMethodName() string {
    return "alibaba.tmallgenie.scp.plan.mouthfour.upload"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaTmallgenieScpPlanMouthfourUploadAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// MonthFourPrRequest Setter
// 请求参数
func (r *AlibabaTmallgenieScpPlanMouthfourUploadAPIRequest) SetMonthFourPrRequest(_monthFourPrRequest *MonthFourPrRequest) error {
    r._monthFourPrRequest = _monthFourPrRequest
    r.Set("month_four_pr_request", _monthFourPrRequest)
    return nil
}

// MonthFourPrRequest Getter
func (r AlibabaTmallgenieScpPlanMouthfourUploadAPIRequest) GetMonthFourPrRequest() *MonthFourPrRequest {
    return r._monthFourPrRequest
}
