package tmallservice

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// TmallServicecenterSettlementStoretransferAudit 新康众审批门店分账
// tmall.servicecenter.settlement.storetransfer.audit
//
// 新康众审批门店分账
func TmallServicecenterSettlementStoretransferAudit(ctx context.Context, clt *core.SDKClient, req *tmallservice.TmallServicecenterSettlementStoretransferAuditAPIRequest, resp *tmallservice.TmallServicecenterSettlementStoretransferAuditAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
