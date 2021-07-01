package eticket

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoVmarketEticketTimeExpandAPIRequest
订单延时接口 API请求
taobao.vmarket.eticket.time.expand

提供码商操作订单延期接口 */
type TaobaoVmarketEticketTimeExpandAPIRequest struct {
	model.Params
	// 订单ID
	_orderId int64
	// 延长天数，延长时间=当前过期时间+延长天数
	_expandDays int64
}

// New
