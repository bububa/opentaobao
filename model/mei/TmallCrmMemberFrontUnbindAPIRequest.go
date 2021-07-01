package mei

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallCrmMemberFrontUnbindAPIRequest
品牌会员解绑 API请求
tmall.crm.member.front.unbind

品牌会员解绑功能 */
type TmallCrmMemberFrontUnbindAPIRequest struct {
	model.Params
	// 会员昵称
	_userNick string
}

// New
