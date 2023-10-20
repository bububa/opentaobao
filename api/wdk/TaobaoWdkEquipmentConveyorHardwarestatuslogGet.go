package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Taobaowdkequipmentconveyorhardwarestatuslogget 硬件状态变化日志查询
// taobao.wdk.equipment.conveyor.hardwarestatuslog.get
//
// 硬件状态变化日志查询
func Taobaowdkequipmentconveyorhardwarestatuslogget(clt *core.SDKClient, req *wdk.TaobaowdkequipmentconveyorhardwarestatusloggetAPIRequest, session string) (*wdk.TaobaowdkequipmentconveyorhardwarestatusloggetAPIResponse, error) {
	var resp wdk.TaobaowdkequipmentconveyorhardwarestatusloggetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
