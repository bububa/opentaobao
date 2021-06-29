package btrip

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
更新审批单 APIRequest
alitrip.btrip.approval.update

更新审批单
*/
type AlitripBtripApprovalUpdateRequest struct {
    model.Params

    // 审批请求对象
    approveApplyRequest   *OpenApproveApplyRq 

}

func NewAlitripBtripApprovalUpdateRequest() *AlitripBtripApprovalUpdateRequest{
    return &AlitripBtripApprovalUpdateRequest{
        Params: model.NewParams(),
    }
}

func (r AlitripBtripApprovalUpdateRequest) GetApiMethodName() string {
    return "alitrip.btrip.approval.update"
}

func (r AlitripBtripApprovalUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlitripBtripApprovalUpdateRequest) SetApproveApplyRequest(approveApplyRequest *OpenApproveApplyRq) error {
    r.approveApplyRequest = approveApplyRequest
    r.Set("approve_apply_request", approveApplyRequest)
    return nil
}

func (r AlitripBtripApprovalUpdateRequest) GetApproveApplyRequest() *OpenApproveApplyRq {
    return r.approveApplyRequest
}

