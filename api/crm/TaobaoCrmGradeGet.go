package crm

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/crm"
)

// TaobaoCrmGradeGet 卖家查询等级规则
// taobao.crm.grade.get
//
// 卖家查询等级规则，包括店铺客户、普通会员、高级会员、VIP会员、至尊VIP会员四个等级的信息
func TaobaoCrmGradeGet(ctx context.Context, clt *core.SDKClient, req *crm.TaobaoCrmGradeGetAPIRequest, resp *crm.TaobaoCrmGradeGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
