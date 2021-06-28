package servicecenter

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
尾款处置方案回传 APIResponse
tmall.car.lease.tailpaymentback

尾款处置方案回传
*/
type TmallCarLeaseTailpaymentbackAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"tmall_car_lease_tailpaymentback_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // result
    
    Result   *TmallCarLeaseTailpaymentbackResult `json:"result,omitempty" xml:"