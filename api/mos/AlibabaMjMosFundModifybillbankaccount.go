package mos

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/mos"
)

// AlibabaMjMosFundModifybillbankaccount 修改付款单的银行账户信息
// alibaba.mj.mos.fund.modifybillbankaccount
//
// 修改付款单的银行账户信息
func AlibabaMjMosFundModifybillbankaccount(clt *core.SDKClient, req *mos.AlibabaMjMosFundModifybillbankaccountAPIRequest, resp *mos.AlibabaMjMosFundModifybillbankaccountAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
