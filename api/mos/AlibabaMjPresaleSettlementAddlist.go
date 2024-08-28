package mos

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/mos"
)

// AlibabaMjPresaleSettlementAddlist 预售结算数据回传
// alibaba.mj.presale.settlement.addlist
//
// 用于预售活动结算数据的回传。
func AlibabaMjPresaleSettlementAddlist(ctx context.Context, clt *core.SDKClient, req *mos.AlibabaMjPresaleSettlementAddlistAPIRequest, resp *mos.AlibabaMjPresaleSettlementAddlistAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
