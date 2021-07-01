package subuser

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoSellercenterSubusersGetAPIRequest
查询指定账户的子账号列表 API请求
taobao.sellercenter.subusers.get

根据主账号nick查询该账号下所有的子账号列表，只能查询属于自己的账号信息 (主账号以及所属子账号) */
type TaobaoSellercenterSubusersGetAPIRequest struct {
	model.Params
	// 表示卖家昵称
	_nick string
}

// New
