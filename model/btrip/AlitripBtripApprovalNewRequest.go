package btrip

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
新建审批单 APIRequest
alitrip.btrip.approval.new

用户新建审批单
*/
type AlitripBtripApprovalNewRequest struct {
    model.Params

    // 申请单
    addApplyRequest   *OpenAddApplyRq 

}

func NewAlitripBtripApprovalNewRequest() *AlitripBtripApprovalNewRequest{
    return &AlitripBtripApprovalNewRequest{
        Params: model.NewParams(),
    }
}

func (r AlitripBtripApprovalNewRequest) GetApiMethodName() string {
    return "alitrip.btrip.approval.new"
}

func (r AlitripBtripApprovalNewRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlitripBtripApprovalNewRequest) SetAddApplyRequest(addApplyRequest *OpenAddApplyRq) error {
    r.addApplyRequest = addApplyRequest
    r.Set("add_apply_request", addApplyRequest)
    return nil
}

func (r AlitripBtripApprovalNewRequest) GetAddApplyRequest() *OpenAddApplyRq {
    return r.addApplyRequest
}

