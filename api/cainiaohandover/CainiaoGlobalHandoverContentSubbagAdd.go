package cainiaohandover

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/cainiaohandover"
)

// CainiaoGlobalHandoverContentSubbagAdd 预约单下追加大包
// cainiao.global.handover.content.subbag.add
//
// 预约单下追加大包
func CainiaoGlobalHandoverContentSubbagAdd(ctx context.Context, clt *core.SDKClient, req *cainiaohandover.CainiaoGlobalHandoverContentSubbagAddAPIRequest, resp *cainiaohandover.CainiaoGlobalHandoverContentSubbagAddAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
