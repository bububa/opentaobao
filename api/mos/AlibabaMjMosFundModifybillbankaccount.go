package mos

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/mos"
)

// Alibabamjmosfundmodifybillbankaccount 修改付款单的银行账户信息
// alibaba.mj.mos.fund.modifybillbankaccount
//
// 修改付款单的银行账户信息
func Alibabamjmosfundmodifybillbankaccount(clt *core.SDKClient, req *mos.AlibabamjmosfundmodifybillbankaccountAPIRequest, session string) (*mos.AlibabamjmosfundmodifybillbankaccountAPIResponse, error) {
	var resp mos.AlibabamjmosfundmodifybillbankaccountAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
