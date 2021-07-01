package jipiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoJipiaoAgentOrderDetailAPIRequest
【机票代理商订单】订单详情 API请求
taobao.jipiao.agent.order.detail

根据淘宝系统订单号获取订单详情信息 */
type TaobaoJipiaoAgentOrderDetailAPIRequest struct {
	model.Params
	// 淘宝订单id列表，当前支持列表长度为1，即当前只支持单个订单详情查询
	_orderIds []int64
}

// New
