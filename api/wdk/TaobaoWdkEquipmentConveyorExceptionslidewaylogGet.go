package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Taobaowdkequipmentconveyorexceptionslidewaylogget 异常通道日志查询
// taobao.wdk.equipment.conveyor.exceptionslidewaylog.get
//
// 五道口悬挂链异常通道事件查询
func Taobaowdkequipmentconveyorexceptionslidewaylogget(clt *core.SDKClient, req *wdk.TaobaowdkequipmentconveyorexceptionslidewayloggetAPIRequest, session string) (*wdk.TaobaowdkequipmentconveyorexceptionslidewayloggetAPIResponse, error) {
	var resp wdk.TaobaowdkequipmentconveyorexceptionslidewayloggetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
