package crm

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/crm"
)

// TaobaoCrmMembersGet 获取卖家的会员（基本查询）
// taobao.crm.members.get
//
// 查询卖家的会员，进行基本的查询，返回符合条件的会员列表
func TaobaoCrmMembersGet(clt *core.SDKClient, req *crm.TaobaoCrmMembersGetAPIRequest, resp *crm.TaobaoCrmMembersGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
