package tmallservice

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// TmallServiceSettlementBillinfoQuery 查询工单结算信息
// tmall.service.settlement.billinfo.query
//
// 提供给服务商查询工单结算信息，包含结算的分成金额以及结算的收款明细，平台抽佣比例。用于服务商进行对账
func TmallServiceSettlementBillinfoQuery(ctx context.Context, clt *core.SDKClient, req *tmallservice.TmallServiceSettlementBillinfoQueryAPIRequest, resp *tmallservice.TmallServiceSettlementBillinfoQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
