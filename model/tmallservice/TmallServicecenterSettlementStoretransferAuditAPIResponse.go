package tmallservice

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
新康众审批门店分账 API返回值 
tmall.servicecenter.settlement.storetransfer.audit

新康众审批门店分账
*/
type TmallServicecenterSettlementStoretransferAuditAPIResponse struct {
    model.CommonResponse
    TmallServicecenterSettlementStoretransferAuditAPIResponseModel
}

// 新康众审批门店分账 成功返回结果
type TmallServicecenterSettlementStoretransferAuditAPIResponseModel struct {
    XMLName xml.Name `xml:"tmall_servicecenter_settlement_storetransfer_audit_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 分账审批通知结果
    Result   *FulfilplatformResult `json:"result,omitempty" xml:"result,omitempty"`
}
