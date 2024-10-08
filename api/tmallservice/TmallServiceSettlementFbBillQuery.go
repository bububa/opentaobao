package tmallservice

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// TmallServiceSettlementFbBillQuery 服务商工单结算对账查询
// tmall.service.settlement.fb.bill.query
//
// 服务商工单结算对账查询，用于查询服务工单对应的结算费用情况，含工单对应的服务费、退款、增加费用、分成费用、提现流水
func TmallServiceSettlementFbBillQuery(ctx context.Context, clt *core.SDKClient, req *tmallservice.TmallServiceSettlementFbBillQueryAPIRequest, resp *tmallservice.TmallServiceSettlementFbBillQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
