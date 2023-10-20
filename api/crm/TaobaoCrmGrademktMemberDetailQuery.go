package crm

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/crm"
)

// TaobaoCrmGrademktMemberDetailQuery 会员等级营销-等级营销活动查询
// taobao.crm.grademkt.member.detail.query
//
// 商家通过该接口查询等级营销活动
func TaobaoCrmGrademktMemberDetailQuery(clt *core.SDKClient, req *crm.TaobaoCrmGrademktMemberDetailQueryAPIRequest, resp *crm.TaobaoCrmGrademktMemberDetailQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
