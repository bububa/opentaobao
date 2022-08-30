package crm

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/crm"
)

// TaobaoCrmMembersSearchPrivy 获取卖家会员(高级查询)
// taobao.crm.members.search.privy
//
// 会员列表的高级查询，接口返回符合条件的会员列表.<br><br/>注：建议获取09年以后的数据，09年之前的数据不是很完整
func TaobaoCrmMembersSearchPrivy(clt *core.SDKClient, req *crm.TaobaoCrmMembersSearchPrivyAPIRequest, session string) (*crm.TaobaoCrmMembersSearchPrivyAPIResponse, error) {
	var resp crm.TaobaoCrmMembersSearchPrivyAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
