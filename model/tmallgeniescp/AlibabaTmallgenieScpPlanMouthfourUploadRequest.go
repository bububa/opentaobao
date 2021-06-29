package tmallgeniescp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
21-M+4PR 回传接口接口 APIRequest
alibaba.tmallgenie.scp.plan.mouthfour.upload

M+4 PR 回传接口
*/
type AlibabaTmallgenieScpPlanMouthfourUploadRequest struct {
    model.Params

    // 请求参数
    monthFourPrRequest   *MonthFourPrRequest 

}

func NewAlibabaTmallgenieScpPlanMouthfourUploadRequest() *AlibabaTmallgenieScpPlanMouthfourUploadRequest{
    return &AlibabaTmallgenieScpPlanMouthfourUploadRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaTmallgenieScpPlanMouthfourUploadRequest) GetApiMethodName() string {
    return "alibaba.tmallgenie.scp.plan.mouthfour.upload"
}

func (r AlibabaTmallgenieScpPlanMouthfourUploadRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaTmallgenieScpPlanMouthfourUploadRequest) SetMonthFourPrRequest(monthFourPrRequest *MonthFourPrRequest) error {
    r.monthFourPrRequest = monthFourPrRequest
    r.Set("month_four_pr_request", monthFourPrRequest)
    return nil
}

func (r AlibabaTmallgenieScpPlanMouthfourUploadRequest) GetMonthFourPrRequest() *MonthFourPrRequest {
    return r.monthFourPrRequest
}

