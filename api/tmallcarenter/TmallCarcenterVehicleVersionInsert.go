package tmallcarenter

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallcarenter"
)

// TmallCarcenterVehicleVersionInsert 汽车EPC版本压缩库新增接口
// tmall.carcenter.vehicle.version.insert
//
// 汽车EPC版本压缩库新增接口
func TmallCarcenterVehicleVersionInsert(ctx context.Context, clt *core.SDKClient, req *tmallcarenter.TmallCarcenterVehicleVersionInsertAPIRequest, resp *tmallcarenter.TmallCarcenterVehicleVersionInsertAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
