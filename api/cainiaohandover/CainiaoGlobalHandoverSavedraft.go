package cainiaohandover

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/cainiaohandover"
)

// CainiaoGlobalHandoverSavedraft 创建交接单草稿
// cainiao.global.handover.savedraft
//
// 提供给ISV通过该接口创建交接单草稿
func CainiaoGlobalHandoverSavedraft(ctx context.Context, clt *core.SDKClient, req *cainiaohandover.CainiaoGlobalHandoverSavedraftAPIRequest, resp *cainiaohandover.CainiaoGlobalHandoverSavedraftAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
