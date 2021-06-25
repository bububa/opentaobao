package servicecenter

import (
    "github.com/bububa/opentaobao/model"
)

/* 
尾款处置方案回传 APIResponse
tmall.car.lease.tailpaymentback

尾款处置方案回传
*/
type TmallCarLeaseTailpaymentbackAPIResponse struct {
    model.CommonResponse
    Response *TmallCarLeaseTailpaymentbackResponse `json:"tmall_car_lease_tailpaymentback_response,omitempty"`
}

type TmallCarLeaseTailpaymentbackResponse struct {

    // result
    Result   *TmallCarLeaseTailpaymentbackResult `json:"result,omitempty"`

}
