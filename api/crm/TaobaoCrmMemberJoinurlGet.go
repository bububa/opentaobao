package crm

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/crm"
)

// TaobaoCrmMemberJoinurlGet 会员入会地址获取
// taobao.crm.member.joinurl.get
//
// 会员入会地址获取
func TaobaoCrmMemberJoinurlGet(clt *core.SDKClient, req *crm.TaobaoCrmMemberJoinurlGetAPIRequest, resp *crm.TaobaoCrmMemberJoinurlGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
