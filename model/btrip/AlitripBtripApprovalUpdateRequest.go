package btrip

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
更新审批单 API请求
alitrip.btrip.approval.update

更新审批单
*/
type AlitripBtripApprovalUpdateRequest struct {
    model.Params
    // 审批请求对象
    _approveApplyRequest   *OpenApproveApplyRq
}

// 初始化AlitripBtripApprovalUpdateRequest对象
func NewAlitripBtripApprovalUpdateRequest() *AlitripBtripApprovalUpdateRequest{
    return &AlitripBtripApprovalUpdateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripBtripApprovalUpdateRequest) GetApiMethodName() string {
    return "alitrip.btrip.approval.update"
}

// IRequest interface 方法, 获取API参数
func (r AlitripBtripApprovalUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ApproveApplyRequest Setter
// 审批请求对象
func (r *AlitripBtripApprovalUpdateRequest) SetApproveApplyRequest(_approveApplyRequest *OpenApproveApplyRq) error {
    r._approveApplyRequest = _approveApplyRequest
    r.Set("approve_apply_request", _approveApplyRequest)
    return nil
}

// ApproveApplyRequest Getter
func (r AlitripBtripApprovalUpdateRequest) GetApproveApplyRequest() *OpenApproveApplyRq {
    return r._approveApplyRequest
}
