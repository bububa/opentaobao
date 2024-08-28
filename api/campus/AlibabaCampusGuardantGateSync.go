package campus

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/campus"
)

// AlibabaCampusGuardantGateSync 网点数据同步
// alibaba.campus.guardant.gate.sync
//
// 门禁供应商创建网点同步
func AlibabaCampusGuardantGateSync(ctx context.Context, clt *core.SDKClient, req *campus.AlibabaCampusGuardantGateSyncAPIRequest, resp *campus.AlibabaCampusGuardantGateSyncAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
