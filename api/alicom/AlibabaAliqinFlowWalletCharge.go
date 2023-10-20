package alicom

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alicom"
)

// Alibabaaliqinflowwalletcharge 流量直充
// alibaba.aliqin.flow.wallet.charge
//
// 流量直充
func Alibabaaliqinflowwalletcharge(clt *core.SDKClient, req *alicom.AlibabaaliqinflowwalletchargeAPIRequest, session string) (*alicom.AlibabaaliqinflowwalletchargeAPIResponse, error) {
	var resp alicom.AlibabaaliqinflowwalletchargeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
