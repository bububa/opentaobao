package mos

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/mos"
)

// AlibabaMosOcTradeSyncbanksale 云闪付、银行卡销售数据回传接口
// alibaba.mos.oc.trade.syncbanksale
//
// 云闪付、银行卡销售数据回传
func AlibabaMosOcTradeSyncbanksale(ctx context.Context, clt *core.SDKClient, req *mos.AlibabaMosOcTradeSyncbanksaleAPIRequest, resp *mos.AlibabaMosOcTradeSyncbanksaleAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
