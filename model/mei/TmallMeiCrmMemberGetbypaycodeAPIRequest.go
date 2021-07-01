package mei

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallMeiCrmMemberGetbypaycodeAPIRequest
支付码获取会员信息 API请求
tmall.mei.crm.member.getbypaycode

通过支付码获取会员信息 */
type TmallMeiCrmMemberGetbypaycodeAPIRequest struct {
	model.Params
	// 会员码
	_payCode string
}

// New
