package alihealthcrm

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealthcrm"
)

// AlibabaAlihealthBabyBaseinfoOrderSync alibaba.alihealth.baby.baseinfo.order.sync
// alibaba.alihealth.baby.baseinfo.order.sync
//
// 育学园将订单信息回传给我们
func AlibabaAlihealthBabyBaseinfoOrderSync(ctx context.Context, clt *core.SDKClient, req *alihealthcrm.AlibabaAlihealthBabyBaseinfoOrderSyncAPIRequest, resp *alihealthcrm.AlibabaAlihealthBabyBaseinfoOrderSyncAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
