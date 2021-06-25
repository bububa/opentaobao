package tmallservice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
新康众审批门店分账 APIRequest
tmall.servicecenter.settlement.storetransfer.audit

新康众审批门店分账
*/
type TmallServicecenterSettlementStoretransferAuditRequest struct {
    model.Params

    // 审批通过
    auditPass   bool 

    // 工单id
    workcardId   int64 

}

func NewTmallServicecenterSettlementStoretransferAuditRequest() *TmallServicecenterSettlementStoretransferAuditRequest{
    return &TmallServicecenterSettlementStoretransferAuditRequest{
        Params: model.NewParams(),
    }
}

func (r TmallServicecenterSettlementStoretransferAuditRequest) GetApiMethodName() string {
    return "tmall.servicecenter.settlement.storetransfer.audit"
}

func (r TmallServicecenterSettlementStoretransferAuditRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallServicecenterSettlementStoretransferAuditRequest) SetAuditPass(auditPass bool) error {
    r.auditPass = auditPass
    r.Set("audit_pass", auditPass)
    return nil
}

func (r TmallServicecenterSettlementStoretransferAuditRequest) GetAuditPass() bool {
    return r.auditPass
}

func (r *TmallServicecenterSettlementStoretransferAuditRequest) SetWorkcardId(workcardId int64) error {
    r.workcardId = workcardId
    r.Set("workcard_id", workcardId)
    return nil
}

func (r TmallServicecenterSettlementStoretransferAuditRequest) GetWorkcardId() int64 {
    return r.workcardId
}

