package mei

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/mei"
)

// TmallMeiCrmMemberGetbypaycode 支付码获取会员信息
// tmall.mei.crm.member.getbypaycode
//
// 通过支付码获取会员信息
func TmallMeiCrmMemberGetbypaycode(clt *core.SDKClient, req *mei.TmallMeiCrmMemberGetbypaycodeAPIRequest, resp *mei.TmallMeiCrmMemberGetbypaycodeAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
