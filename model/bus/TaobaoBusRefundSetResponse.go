package bus

import (
    "github.com/bububa/opentaobao/model"
)

/* 
B2B退票申请接口 APIResponse
taobao.bus.refund.set

B2B业务支持退票
*/
type TaobaoBusRefundSetAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoBusRefundSetResponse `json:"bus_refund_set_response,omitempty"` 
    TaobaoBusRefundSetResponse
}

/* model for simplify = false
type TaobaoBusRefundSetResponse struct {

    // result
    
    Result  *struct {
        B2BRefundOrderRp  *B2BRefundOrderRp `json:"b2b_refund_order_rp,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type TaobaoBusRefundSetResponse struct {

    // result
    Result   *B2BRefundOrderRp `json:"result,omitempty"`

}
