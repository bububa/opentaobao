package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoSimbaRtrptCustGetAPIRequest
获取账户实时报表数据 API请求
taobao.simba.rtrpt.cust.get

获取账户实时报表数据 */
type TaobaoSimbaRtrptCustGetAPIRequest struct {
	model.Params
	// 昵称
	_nick string
	// 日期，格式yyyy-mm-dd
	_theDate string
}

// New
