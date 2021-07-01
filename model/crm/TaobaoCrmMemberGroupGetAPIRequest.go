package crm

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoCrmMemberGroupGetAPIRequest
获取买家身上的标签 API请求
taobao.crm.member.group.get

获取买家身上的标签，不返回标签的总人数 */
type TaobaoCrmMemberGroupGetAPIRequest struct {
	model.Params
	// 会员Nick
	_buyerNick string
}

// New
