package crm

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/crm"
)

// TaobaoCrmExchangeActivityCreate 创建积分兑换活动
// taobao.crm.exchange.activity.create
//
// 创建针对积分兑换类型的活动
func TaobaoCrmExchangeActivityCreate(ctx context.Context, clt *core.SDKClient, req *crm.TaobaoCrmExchangeActivityCreateAPIRequest, resp *crm.TaobaoCrmExchangeActivityCreateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
