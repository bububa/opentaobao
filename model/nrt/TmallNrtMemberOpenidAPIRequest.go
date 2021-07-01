package nrt

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallNrtMemberOpenidAPIRequest
根据会员手机查询openId API请求
tmall.nrt.member.openid

根据会员手机查询openId */
type TmallNrtMemberOpenidAPIRequest struct {
	model.Params
	// 会员DTO
	_nrtMemberDto *NrtMemberDto
}

// New
