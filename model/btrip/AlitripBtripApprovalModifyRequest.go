package btrip

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
修改审批单 APIRequest
alitrip.btrip.approval.modify

修改审批单
*/
type AlitripBtripApprovalModifyRequest struct {
    model.Params

    // 申请单
    addApplyRequest   *OpenApiNewApplyRq 

}

func NewAlitripBtripApprovalModifyRequest() *AlitripBtripApprovalModifyRequest{
    return &AlitripBtripApprovalModifyRequest{
        Params: model.NewParams(),
    }
}

func (r AlitripBtripApprovalModifyRequest) GetApiMethodName() string {
    return "alitrip.btrip.approval.modify"
}

func (r AlitripBtripApprovalModifyRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlitripBtripApprovalModifyRequest) SetAddApplyRequest(addApplyRequest *OpenApiNewApplyRq) error {
    r.addApplyRequest = addApplyRequest
    r.Set("add_apply_request", addApplyRequest)
    return nil
}

func (r AlitripBtripApprovalModifyRequest) GetAddApplyRequest() *OpenApiNewApplyRq {
    return r.addApplyRequest
}

