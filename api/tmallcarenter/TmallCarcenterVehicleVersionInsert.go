package tmallcarenter

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallcarenter"
)

/* TmallCarcenterVehicleVersionInsert
汽车EPC版本压缩库新增接口
tmall.carcenter.vehicle.version.insert

汽车EPC版本压缩库新增接口 */
func TmallCarcenterVehicleVersionInsert(clt *core.SDKClient, req *tmallcarenter.TmallCarcenterVehicleVersionInsertAPIRequest, session string) (*tmallcarenter.TmallCarcenterVehicleVersionInsertAPIResponse, error) {
	var resp tmallcarenter.TmallCarcenterVehicleVersionInsertAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
