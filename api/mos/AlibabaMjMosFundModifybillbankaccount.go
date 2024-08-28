package mos

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/mos"
)

// AlibabaMjMosFundModifybillbankaccount 修改付款单的银行账户信息
// alibaba.mj.mos.fund.modifybillbankaccount
//
// 修改付款单的银行账户信息
func AlibabaMjMosFundModifybillbankaccount(ctx context.Context, clt *core.SDKClient, req *mos.AlibabaMjMosFundModifybillbankaccountAPIRequest, resp *mos.AlibabaMjMosFundModifybillbankaccountAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
