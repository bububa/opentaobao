package mos

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/mos"
)

// AlibabaMjOcBigposBanksaleAdjustmentApply 大pos银行卡调账申请
// alibaba.mj.oc.bigpos.banksale.adjustment.apply
//
// 大pos银行卡调账申请
func AlibabaMjOcBigposBanksaleAdjustmentApply(ctx context.Context, clt *core.SDKClient, req *mos.AlibabaMjOcBigposBanksaleAdjustmentApplyAPIRequest, resp *mos.AlibabaMjOcBigposBanksaleAdjustmentApplyAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
