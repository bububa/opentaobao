package mos

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/mos"
)

// AlibabaMjMosFundModifybillbankaccount 修改付款单的银行账户信息
// alibaba.mj.mos.fund.modifybillbankaccount
//
// 修改付款单的银行账户信息
func AlibabaMjMosFundModifybillbankaccount(clt *core.SDKClient, req *mos.AlibabaMjMosFundModifybillbankaccountAPIRequest, session string) (*mos.AlibabaMjMosFundModifybillbankaccountAPIResponse, error) {
	var resp mos.AlibabaMjMosFundModifybillbankaccountAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
