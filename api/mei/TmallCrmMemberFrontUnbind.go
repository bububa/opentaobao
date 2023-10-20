package mei

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/mei"
)

// TmallCrmMemberFrontUnbind 品牌会员解绑
// tmall.crm.member.front.unbind
//
// 品牌会员解绑功能
func TmallCrmMemberFrontUnbind(clt *core.SDKClient, req *mei.TmallCrmMemberFrontUnbindAPIRequest, resp *mei.TmallCrmMemberFrontUnbindAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
