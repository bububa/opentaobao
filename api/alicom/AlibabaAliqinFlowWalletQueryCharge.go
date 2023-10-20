package alicom

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alicom"
)

// Alibabaaliqinflowwalletquerycharge 查询流量充值状态
// alibaba.aliqin.flow.wallet.query.charge
//
// 查询流量充值状态
func Alibabaaliqinflowwalletquerycharge(clt *core.SDKClient, req *alicom.AlibabaaliqinflowwalletquerychargeAPIRequest, session string) (*alicom.AlibabaaliqinflowwalletquerychargeAPIResponse, error) {
	var resp alicom.AlibabaaliqinflowwalletquerychargeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
