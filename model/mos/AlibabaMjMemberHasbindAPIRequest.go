package mos

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaMjMemberHasbindAPIRequest
喵街会员是否绑定 API请求
alibaba.mj.member.hasbind

喵街检测用户是否为数字化会员 */
type AlibabaMjMemberHasbindAPIRequest struct {
	model.Params
	// user_id
	_userId int64
	// open_id
	_openId string
}

// New
