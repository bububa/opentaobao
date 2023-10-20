package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// TaobaoWdkEquipmentConveyorHardwarestatuslogGet 硬件状态变化日志查询
// taobao.wdk.equipment.conveyor.hardwarestatuslog.get
//
// 硬件状态变化日志查询
func TaobaoWdkEquipmentConveyorHardwarestatuslogGet(clt *core.SDKClient, req *wdk.TaobaoWdkEquipmentConveyorHardwarestatuslogGetAPIRequest, resp *wdk.TaobaoWdkEquipmentConveyorHardwarestatuslogGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
