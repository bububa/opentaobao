package alicom

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alicom"
)

// Alibabaaliqinflowwalletconsume 流量扣减
// alibaba.aliqin.flow.wallet.consume
//
// 流量钱包流量扣减接口
func Alibabaaliqinflowwalletconsume(clt *core.SDKClient, req *alicom.AlibabaaliqinflowwalletconsumeAPIRequest, session string) (*alicom.AlibabaaliqinflowwalletconsumeAPIResponse, error) {
	var resp alicom.AlibabaaliqinflowwalletconsumeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
