package tmallcarenter

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallcarenter"
)

// TmallCarcenterVehicleChasisInsert EPC车型底盘压缩库新增接口
// tmall.carcenter.vehicle.chasis.insert
//
// EPC车型底盘压缩库新增接口
func TmallCarcenterVehicleChasisInsert(ctx context.Context, clt *core.SDKClient, req *tmallcarenter.TmallCarcenterVehicleChasisInsertAPIRequest, resp *tmallcarenter.TmallCarcenterVehicleChasisInsertAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
