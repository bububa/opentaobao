package nrt

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallNrtMemberSynchronizeAPIRequest
新零售会员同步接口 API请求
tmall.nrt.member.synchronize

新零售会员上翻接口，商家的会员信息同步至阿里侧 */
type TmallNrtMemberSynchronizeAPIRequest struct {
	model.Params
	// 会员DTO
	_nrtMemberDto *NrtMemberDto
}

// New
