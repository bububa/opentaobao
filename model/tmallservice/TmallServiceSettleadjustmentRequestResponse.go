package tmallservice

import (
    "github.com/bububa/opentaobao/model"
)

/* 
创建结算调整单 APIResponse
tmall.service.settleadjustment.request

提供给服务商在对结算有异议时，发起结算调整单。
通过说明工单ID，调整费用值，调整原因进行新建结算调整单。
*/
type TmallServiceSettleadjustmentRequestAPIResponse struct {
    model.CommonResponse
    // Response *TmallServiceSettleadjustmentRequestResponse `json:"tmall_service_settleadjustment_request_response,omitempty"` 
    TmallServiceSettleadjustmentRequestResponse
}

/* model for simplify = false
type TmallServiceSettleadjustmentRequestResponse struct {

    // result
    
    Result  *struct {
        TmallServiceSettleadjustmentRequestResult  *TmallServiceSettleadjustmentRequestResult `json:"tmall_service_settleadjustment_request_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type TmallServiceSettleadjustmentRequestResponse struct {

    // result
    Result   *TmallServiceSettleadjustmentRequestResult `json:"result,omitempty"`

}
