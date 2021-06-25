package user

import (
    "github.com/bububa/opentaobao/model"
)

/* 
修改结算调整单 APIResponse
tmall.service.settleadjustment.modify

提供给服务商在对结算有异议时，发起结算调整单。
通过说明调整单ID，调整费用值，调整原因进行结算调整单修改。
*/
type TmallServiceSettleadjustmentModifyAPIResponse struct {
    model.CommonResponse
    Response *TmallServiceSettleadjustmentModifyResponse `json:"tmall_service_settleadjustment_modify_response,omitempty"`
}

type TmallServiceSettleadjustmentModifyResponse struct {

    // result
    Result   *TmallServiceSettleadjustmentModifyResult `json:"result,omitempty"`

}
