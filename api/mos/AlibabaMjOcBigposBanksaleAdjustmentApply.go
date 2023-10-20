package mos

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/mos"
)

// AlibabaMjOcBigposBanksaleAdjustmentApply 大pos银行卡调账申请
// alibaba.mj.oc.bigpos.banksale.adjustment.apply
//
// 大pos银行卡调账申请
func AlibabaMjOcBigposBanksaleAdjustmentApply(clt *core.SDKClient, req *mos.AlibabaMjOcBigposBanksaleAdjustmentApplyAPIRequest, resp *mos.AlibabaMjOcBigposBanksaleAdjustmentApplyAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
