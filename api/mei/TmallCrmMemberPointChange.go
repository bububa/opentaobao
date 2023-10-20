package mei

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/mei"
)

// TmallCrmMemberPointChange 会员积分变更
// tmall.crm.member.point.change
//
// 品牌CRM项目中，会员积分变更接口。
func TmallCrmMemberPointChange(clt *core.SDKClient, req *mei.TmallCrmMemberPointChangeAPIRequest, resp *mei.TmallCrmMemberPointChangeAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
