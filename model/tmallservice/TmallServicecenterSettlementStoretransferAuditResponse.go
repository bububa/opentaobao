package tmallservice

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
新康众审批门店分账 APIResponse
tmall.servicecenter.settlement.storetransfer.audit

新康众审批门店分账
*/
type TmallServicecenterSettlementStoretransferAuditAPIResponse struct {
    model.CommonResponse
    TmallServicecenterSettlementStoretransferAuditResponse
}

type TmallServicecenterSettlementStoretransferAuditResponse struct {
    XMLName xml.Name `xml:"tmall_servicecenter_settlement_storetransfer_audit_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 分账审批通知结果
    
    Result   *FulfilplatformResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
