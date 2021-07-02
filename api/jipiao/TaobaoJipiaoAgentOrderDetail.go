package jipiao

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/jipiao"
)

// TaobaoJipiaoAgentOrderDetail 【机票代理商订单】订单详情
// taobao.jipiao.agent.order.detail
//
// 根据淘宝系统订单号获取订单详情信息
func TaobaoJipiaoAgentOrderDetail(clt *core.SDKClient, req *jipiao.TaobaoJipiaoAgentOrderDetailAPIRequest, session string) (*jipiao.TaobaoJipiaoAgentOrderDetailAPIResponse, error) {
	var resp jipiao.TaobaoJipiaoAgentOrderDetailAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
