package bus

import (
    "github.com/bububa/opentaobao/model"
)

/* 
查询退票费用明细 APIResponse
taobao.bus.refundfee.get

查询退票的费用信息
*/
type TaobaoBusRefundfeeGetAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoBusRefundfeeGetResponse `json:"bus_refundfee_get_response,omitempty"` 
    TaobaoBusRefundfeeGetResponse
}

/* model for simplify = false
type TaobaoBusRefundfeeGetResponse struct {

    // result
    
    Result  *struct {
        B2BQueryRefundFeeRp  *B2BQueryRefundFeeRp `json:"b2b_query_refund_fee_rp,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type TaobaoBusRefundfeeGetResponse struct {

    // result
    Result   *B2BQueryRefundFeeRp `json:"result,omitempty"`

}
