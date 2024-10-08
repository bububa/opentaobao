package crm

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/crm"
)

// TaobaoCrmGradeSet 卖家设置等级规则
// taobao.crm.grade.set
//
// 设置等级信息，可以设置层级等级，也可以单独设置一个等级。出于安全原因，折扣现最低只能设置到700即7折。
func TaobaoCrmGradeSet(ctx context.Context, clt *core.SDKClient, req *crm.TaobaoCrmGradeSetAPIRequest, resp *crm.TaobaoCrmGradeSetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
