package crm

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/crm"
)

// TaobaoCrmMembersSearch 获取卖家会员（高级查询）
// taobao.crm.members.search
//
// 会员列表的高级查询，接口返回符合条件的会员列表.&lt;br&gt;&lt;br/&gt;注：建议获取09年以后的数据，09年之前的数据不是很完整
func TaobaoCrmMembersSearch(clt *core.SDKClient, req *crm.TaobaoCrmMembersSearchAPIRequest, resp *crm.TaobaoCrmMembersSearchAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
