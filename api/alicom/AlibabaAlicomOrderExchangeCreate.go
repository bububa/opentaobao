package alicom

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alicom"
)

// AlibabaAlicomOrderExchangeCreate 代理商积分兑换接口
// alibaba.alicom.order.exchange.create
//
// 代理商调用该接口来进行积分兑换
func AlibabaAlicomOrderExchangeCreate(clt *core.SDKClient, req *alicom.AlibabaAlicomOrderExchangeCreateAPIRequest, session string) (*alicom.AlibabaAlicomOrderExchangeCreateAPIResponse, error) {
	var resp alicom.AlibabaAlicomOrderExchangeCreateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
