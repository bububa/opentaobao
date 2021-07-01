package crm

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoCrmExchangeActivityCreateAPIRequest
创建积分兑换活动 API请求
taobao.crm.exchange.activity.create

创建针对积分兑换类型的活动 */
type TaobaoCrmExchangeActivityCreateAPIRequest struct {
	model.Params
	// 创建积分兑换活动
	_exchangeActivityCreateDto *ExchangeActivityCreateDto
}

// New
