package mos

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/mos"
)

// AlibabaMjPresaleSettlementStatistics 预购结算数据统计
// alibaba.mj.presale.settlement.statistics
//
// 预购结算数据统计
func AlibabaMjPresaleSettlementStatistics(ctx context.Context, clt *core.SDKClient, req *mos.AlibabaMjPresaleSettlementStatisticsAPIRequest, resp *mos.AlibabaMjPresaleSettlementStatisticsAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
