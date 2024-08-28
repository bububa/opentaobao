package tmallcar

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallcar"
)

// TaobaoCarVehicleinfoRegister 全量车型导入
// taobao.car.vehicleinfo.register
//
// 全量车型导入
func TaobaoCarVehicleinfoRegister(ctx context.Context, clt *core.SDKClient, req *tmallcar.TaobaoCarVehicleinfoRegisterAPIRequest, resp *tmallcar.TaobaoCarVehicleinfoRegisterAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
