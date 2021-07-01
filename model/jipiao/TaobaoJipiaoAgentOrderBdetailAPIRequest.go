package jipiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoJipiaoAgentOrderBdetailAPIRequest
【机票代理商订单】采购订单详情 API请求
taobao.jipiao.agent.order.bdetail

根据淘宝系统订单号获取订单详情信息 */
type TaobaoJipiaoAgentOrderBdetailAPIRequest struct {
	model.Params
	// 订单号
	_orderId int64
}

// New
