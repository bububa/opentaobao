package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// TaobaoWdkEquipmentConveyorSystemeventGet 获取悬挂链系统事件
// taobao.wdk.equipment.conveyor.systemevent.get
//
// 五道口悬挂链系统事件查询
func TaobaoWdkEquipmentConveyorSystemeventGet(clt *core.SDKClient, req *wdk.TaobaoWdkEquipmentConveyorSystemeventGetAPIRequest, resp *wdk.TaobaoWdkEquipmentConveyorSystemeventGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
