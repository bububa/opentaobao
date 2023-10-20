package crm

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/crm"
)

// TaobaoCrmGrademktMemberDetailCreate 会员等级营销-创建商品等级营销明细
// taobao.crm.grademkt.member.detail.create
//
// 创建商品等级营销明细
func TaobaoCrmGrademktMemberDetailCreate(clt *core.SDKClient, req *crm.TaobaoCrmGrademktMemberDetailCreateAPIRequest, resp *crm.TaobaoCrmGrademktMemberDetailCreateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
