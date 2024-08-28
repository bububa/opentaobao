package aedropshiper

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aedropshiper"
)

// AliexpressDsCommissionorderListbyindex 联盟订单分页查询
// aliexpress.ds.commissionorder.listbyindex
//
// 联盟订单分页查询
func AliexpressDsCommissionorderListbyindex(ctx context.Context, clt *core.SDKClient, req *aedropshiper.AliexpressDsCommissionorderListbyindexAPIRequest, resp *aedropshiper.AliexpressDsCommissionorderListbyindexAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
