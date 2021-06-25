package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
saas 售后逆向 商户拒绝用户逆向申请 APIRequest
alibaba.tcls.aelophy.refund.disagree

saas 售后逆向 商户拒绝用户逆向申请
*/
type AlibabaTclsAelophyRefundDisagreeRequest struct {
    model.Params

    // 退款单ID
    refundId   string 

    // 拒绝原因
    rejectReason   string 

}

func NewAlibabaTclsAelophyRefundDisagreeRequest() *AlibabaTclsAelophyRefundDisagreeRequest{
    return &AlibabaTclsAelophyRefundDisagreeRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaTclsAelophyRefundDisagreeRequest) GetApiMethodName() string {
    return "alibaba.tcls.aelophy.refund.disagree"
}

func (r AlibabaTclsAelophyRefundDisagreeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaTclsAelophyRefundDisagreeRequest) SetRefundId(refundId string) error {
    r.refundId = refundId
    r.Set("refund_id", refundId)
    return nil
}

func (r AlibabaTclsAelophyRefundDisagreeRequest) GetRefundId() string {
    return r.refundId
}

func (r *AlibabaTclsAelophyRefundDisagreeRequest) SetRejectReason(rejectReason string) error {
    r.rejectReason = rejectReason
    r.Set("reject_reason", rejectReason)
    return nil
}

func (r AlibabaTclsAelophyRefundDisagreeRequest) GetRejectReason() string {
    return r.rejectReason
}

