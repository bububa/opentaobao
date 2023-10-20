package fundplatform

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fundplatform"
)

// Alibabafundplatformaccountcharge 资金平台余额账户充值
// alibaba.fundplatform.account.charge
//
// 资金平台余额账户充值【创建账户&amp;返回付款URL】
func Alibabafundplatformaccountcharge(clt *core.SDKClient, req *fundplatform.AlibabafundplatformaccountchargeAPIRequest, session string) (*fundplatform.AlibabafundplatformaccountchargeAPIResponse, error) {
	var resp fundplatform.AlibabafundplatformaccountchargeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
