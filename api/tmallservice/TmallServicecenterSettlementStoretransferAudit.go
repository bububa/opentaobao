package tmallservice

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/tmallservice"
)

/* 
新康众审批门店分账 
tmall.servicecenter.settlement.storetransfer.audit

新康众审批门店分账
*/
func TmallServicecenterSettlementStoretransferAudit(clt *core.SDKClient, req *tmallservice.TmallServicecenterSettlementStoretransferAuditRequest, session string) (*tmallservice.TmallServicecenterSettlementStoretransferAuditAPIResponse, error) {
    var resp tmallservice.TmallServicecenterSettlementStoretransferAuditAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
