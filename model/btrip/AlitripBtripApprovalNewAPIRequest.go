package btrip

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
新建审批单 API请求
alitrip.btrip.approval.new

用户新建审批单
*/
type AlitripBtripApprovalNewAPIRequest struct {
    model.Params
    // 申请单
    _addApplyRequest   *OpenAddApplyRq
}

// 初始化AlitripBtripApprovalNewAPIRequest对象
func NewAlitripBtripApprovalNewRequest() *AlitripBtripApprovalNewAPIRequest{
    return &AlitripBtripApprovalNewAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripBtripApprovalNewAPIRequest) GetApiMethodName() string {
    return "alitrip.btrip.approval.new"
}

// IRequest interface 方法, 获取API参数
func (r AlitripBtripApprovalNewAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// AddApplyRequest Setter
// 申请单
func (r *AlitripBtripApprovalNewAPIRequest) SetAddApplyRequest(_addApplyRequest *OpenAddApplyRq) error {
    r._addApplyRequest = _addApplyRequest
    r.Set("add_apply_request", _addApplyRequest)
    return nil
}

// AddApplyRequest Getter
func (r AlitripBtripApprovalNewAPIRequest) GetAddApplyRequest() *OpenAddApplyRq {
    return r._addApplyRequest
}
