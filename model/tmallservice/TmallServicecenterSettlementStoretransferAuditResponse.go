package tmallservice

import (
    "github.com/bububa/opentaobao/model"
)

/* 
新康众审批门店分账 APIResponse
tmall.servicecenter.settlement.storetransfer.audit

新康众审批门店分账
*/
type TmallServicecenterSettlementStoretransferAuditAPIResponse struct {
    model.CommonResponse
    // Response *TmallServicecenterSettlementStoretransferAuditResponse `json:"tmall_servicecenter_settlement_storetransfer_audit_response,omitempty"` 
    TmallServicecenterSettlementStoretransferAuditResponse
}

/* model for simplify = false
type TmallServicecenterSettlementStoretransferAuditResponse struct {

    // 分账审批通知结果
    
    Result  *struct {
        FulfilplatformResult  *FulfilplatformResult `json:"fulfilplatform_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type TmallServicecenterSettlementStoretransferAuditResponse struct {

    // 分账审批通知结果
    Result   *FulfilplatformResult `json:"result,omitempty"`

}
