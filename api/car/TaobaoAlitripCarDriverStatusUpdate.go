package car

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/car"
)

// Taobaoalitripcardriverstatusupdate 司机服务状态更新接口
// taobao.alitrip.car.driver.status.update
//
// 飞猪用车业务回调接口，用于服务商实时回传更新司机当前服务状态
func Taobaoalitripcardriverstatusupdate(clt *core.SDKClient, req *car.TaobaoalitripcardriverstatusupdateAPIRequest, session string) (*car.TaobaoalitripcardriverstatusupdateAPIResponse, error) {
	var resp car.TaobaoalitripcardriverstatusupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
