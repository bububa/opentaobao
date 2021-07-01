package crm

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoCrmExchangeCrowdinstanceDeleteAPIRequest
删除人群实例中的指定买家 API请求
taobao.crm.exchange.crowdinstance.delete

删除人群实例中的指定买家 */
type TaobaoCrmExchangeCrowdinstanceDeleteAPIRequest struct {
	model.Params
	// 操作原因
	_reason string
	// 人群实例ID
	_crowdInstanceId int64
	// 买家昵称
	_buyerNick string
}

// New
