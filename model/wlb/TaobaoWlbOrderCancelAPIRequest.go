package wlb

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoWlbOrderCancelAPIRequest
取消物流宝订单 API请求
taobao.wlb.order.cancel

取消物流宝订单 */
type TaobaoWlbOrderCancelAPIRequest struct {
	model.Params
	// 物流宝订单编号
	_wlbOrderCode string
}

// New
