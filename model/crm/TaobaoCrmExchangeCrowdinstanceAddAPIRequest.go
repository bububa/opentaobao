package crm

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoCrmExchangeCrowdinstanceAddAPIRequest
向活动人群实例中增加买家 API请求
taobao.crm.exchange.crowdinstance.add

向活动人群实例中增加买家 */
type TaobaoCrmExchangeCrowdinstanceAddAPIRequest struct {
	model.Params
	// 操作原因
	_reason string
	// 人群实例ID
	_crowdInstanceId int64
	// 买家昵称
	_buyerNick string
}

// New
