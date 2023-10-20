package mei

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/mei"
)

// TmallCrmMemberFrontUnbindPrivy 品牌会员解绑(隐私号版）
// tmall.crm.member.front.unbind.privy
//
// 品牌会员解绑功能
func TmallCrmMemberFrontUnbindPrivy(clt *core.SDKClient, req *mei.TmallCrmMemberFrontUnbindPrivyAPIRequest, resp *mei.TmallCrmMemberFrontUnbindPrivyAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
