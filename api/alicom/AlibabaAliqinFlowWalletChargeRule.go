package alicom

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alicom"
)

// Alibabaaliqinflowwalletchargerule 流量钱包直充（根据号码归属地省份路由）
// alibaba.aliqin.flow.wallet.charge.rule
//
// 流量钱包直充（根据号码归属地省份路由）
func Alibabaaliqinflowwalletchargerule(clt *core.SDKClient, req *alicom.AlibabaaliqinflowwalletchargeruleAPIRequest, session string) (*alicom.AlibabaaliqinflowwalletchargeruleAPIResponse, error) {
	var resp alicom.AlibabaaliqinflowwalletchargeruleAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
