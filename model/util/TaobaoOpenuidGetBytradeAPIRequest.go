package util

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoOpenuidGetBytradeAPIRequest
通过订单获取对应买家的openUID API请求
taobao.openuid.get.bytrade

通过订单获取对应买家的openUID,需要卖家授权 */
type TaobaoOpenuidGetBytradeAPIRequest struct {
	model.Params
	// 订单ID
	_tid int64
}

// New
