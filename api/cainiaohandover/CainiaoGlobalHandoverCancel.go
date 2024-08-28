package cainiaohandover

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/cainiaohandover"
)

// CainiaoGlobalHandoverCancel 取消交接单
// cainiao.global.handover.cancel
//
// 提供给ISV通过该接口取消交接单
func CainiaoGlobalHandoverCancel(ctx context.Context, clt *core.SDKClient, req *cainiaohandover.CainiaoGlobalHandoverCancelAPIRequest, resp *cainiaohandover.CainiaoGlobalHandoverCancelAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
