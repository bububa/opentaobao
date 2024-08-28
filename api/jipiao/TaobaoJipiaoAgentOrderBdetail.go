package jipiao

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/jipiao"
)

// TaobaoJipiaoAgentOrderBdetail 【机票代理商订单】采购订单详情
// taobao.jipiao.agent.order.bdetail
//
// 根据淘宝系统订单号获取订单详情信息
func TaobaoJipiaoAgentOrderBdetail(ctx context.Context, clt *core.SDKClient, req *jipiao.TaobaoJipiaoAgentOrderBdetailAPIRequest, resp *jipiao.TaobaoJipiaoAgentOrderBdetailAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
