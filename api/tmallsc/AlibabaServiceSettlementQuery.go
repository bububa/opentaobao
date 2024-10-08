package tmallsc

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallsc"
)

// AlibabaServiceSettlementQuery 服务平台结算单明细查询服务
// alibaba.service.settlement.query
//
// 给服务商提供结算单明细查询功能
func AlibabaServiceSettlementQuery(ctx context.Context, clt *core.SDKClient, req *tmallsc.AlibabaServiceSettlementQueryAPIRequest, resp *tmallsc.AlibabaServiceSettlementQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
