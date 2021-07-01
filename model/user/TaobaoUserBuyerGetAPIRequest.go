package user

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoUserBuyerGetAPIRequest
查询买家信息API API请求
taobao.user.buyer.get

查询买家信息API，只能买家类应用调用。 */
type TaobaoUserBuyerGetAPIRequest struct {
	model.Params
	// 只返回nick, avatar参数
	_fields string
}

// New
