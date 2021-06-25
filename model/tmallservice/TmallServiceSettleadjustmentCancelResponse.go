package tmallservice

import (
    "github.com/bububa/opentaobao/model"
)

/* 
取消结算调整单 APIResponse
tmall.service.settleadjustment.cancel

提供给服务商在对取消已经发起的结算调整单。
通过说明调整单ID进行结算调整单取消。
*/
type TmallServiceSettleadjustmentCancelAPIResponse struct {
    model.CommonResponse
    Response *TmallServiceSettleadjustmentCancelResponse `json:"tmall_service_settleadjustment_cancel_response,omitempty"`
}

type TmallServiceSettleadjustmentCancelResponse struct {

    // result
    Result   *TmallServiceSettleadjustmentCancelResult `json:"result,omitempty"`

}
