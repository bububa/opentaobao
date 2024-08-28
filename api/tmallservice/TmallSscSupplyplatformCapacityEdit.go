package tmallservice

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// TmallSscSupplyplatformCapacityEdit 容量编辑
// tmall.ssc.supplyplatform.capacity.edit
//
// 容量编辑
func TmallSscSupplyplatformCapacityEdit(ctx context.Context, clt *core.SDKClient, req *tmallservice.TmallSscSupplyplatformCapacityEditAPIRequest, resp *tmallservice.TmallSscSupplyplatformCapacityEditAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
