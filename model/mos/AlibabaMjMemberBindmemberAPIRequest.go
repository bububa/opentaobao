package mos

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaMjMemberBindmemberAPIRequest
绑定会员 API请求
alibaba.mj.member.bindmember

用于绑定喵街数字化会员 */
type AlibabaMjMemberBindmemberAPIRequest struct {
	model.Params
	// 用户号
	_userId int64
	// 商城Id
	_mallId int64
	// open_id
	_openId string
	// 渠道
	_channel string
}

// New
