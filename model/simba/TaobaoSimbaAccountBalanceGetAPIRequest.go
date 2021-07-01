package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoSimbaAccountBalanceGetAPIRequest
获取实时余额，”元”为单位 API请求
taobao.simba.account.balance.get

获取实时余额，”元”为单位 */
type TaobaoSimbaAccountBalanceGetAPIRequest struct {
	model.Params
	// 主人昵称
	_nick string
}

// New
