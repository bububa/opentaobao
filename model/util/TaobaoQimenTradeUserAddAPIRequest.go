package util

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoQimenTradeUserAddAPIRequest
添加奇门订单链路用户 API请求
taobao.qimen.trade.user.add

添加奇门订单链路用户 */
type TaobaoQimenTradeUserAddAPIRequest struct {
	model.Params
	// 商家备注
	_memo string
}

// New
