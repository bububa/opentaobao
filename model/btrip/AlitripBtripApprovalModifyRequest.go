package btrip

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
修改审批单 API请求
alitrip.btrip.approval.modify

修改审批单
*/
type AlitripBtripApprovalModifyRequest struct {
    model.Params
    // 申请单
    addApplyRequest   *OpenApiNewApplyRq
}

// 初始化AlitripBtripApprovalModifyRequest对象
func NewAlitripBtripApprovalModifyRequest() *AlitripBtripApprovalModifyRequest{
    return &AlitripBtripApprovalModifyRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripBtripApprovalModifyRequest) GetApiMethodName() string {
    return "alitrip.btrip.approval.modify"
}

// IRequest interface 方法, 获取API参数
func (r AlitripBtripApprovalModifyRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// AddApplyRequest Setter
// 申请单
func (r *AlitripBtripApprovalModifyRequest) SetAddApplyRequest(addApplyRequest *OpenApiNewApplyRq) error {
    r.addApplyRequest = addApplyRequest
    r.Set("add_apply_request", addApplyRequest)
    return nil
}

// AddApplyRequest Getter
func (r AlitripBtripApprovalModifyRequest) GetAddApplyRequest() *OpenApiNewApplyRq {
    return r.addApplyRequest
}
