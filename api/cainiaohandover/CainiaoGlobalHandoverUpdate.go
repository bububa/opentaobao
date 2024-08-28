package cainiaohandover

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/cainiaohandover"
)

// CainiaoGlobalHandoverUpdate 修改交接单
// cainiao.global.handover.update
//
// 提供给ISV通过该接口修改交接单
func CainiaoGlobalHandoverUpdate(ctx context.Context, clt *core.SDKClient, req *cainiaohandover.CainiaoGlobalHandoverUpdateAPIRequest, resp *cainiaohandover.CainiaoGlobalHandoverUpdateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
