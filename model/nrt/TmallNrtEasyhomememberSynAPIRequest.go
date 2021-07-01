package nrt

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallNrtEasyhomememberSynAPIRequest
会员信息同 API请求
tmall.nrt.easyhomemember.syn

居然之家将会员信息同步到零售中台 包含基本的会员信息 */
type TmallNrtEasyhomememberSynAPIRequest struct {
	model.Params
	// 入参
	_param *ExternalMemberDto
}

// New
