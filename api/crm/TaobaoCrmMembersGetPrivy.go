package crm

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/crm"
)

// TaobaoCrmMembersGetPrivy 获取卖家会员(基本查询)
// taobao.crm.members.get.privy
//
// 查询卖家的会员，进行基本的查询，返回符合条件的会员列表
func TaobaoCrmMembersGetPrivy(clt *core.SDKClient, req *crm.TaobaoCrmMembersGetPrivyAPIRequest, resp *crm.TaobaoCrmMembersGetPrivyAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
