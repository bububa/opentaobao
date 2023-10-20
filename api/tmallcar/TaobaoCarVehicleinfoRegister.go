package tmallcar

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallcar"
)

// Taobaocarvehicleinforegister 全量车型导入
// taobao.car.vehicleinfo.register
//
// 全量车型导入
func Taobaocarvehicleinforegister(clt *core.SDKClient, req *tmallcar.TaobaocarvehicleinforegisterAPIRequest, session string) (*tmallcar.TaobaocarvehicleinforegisterAPIResponse, error) {
	var resp tmallcar.TaobaocarvehicleinforegisterAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
