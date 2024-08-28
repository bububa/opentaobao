package wdk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// TaobaoWdkEquipmentConveyorSystemeventGet 获取悬挂链系统事件
// taobao.wdk.equipment.conveyor.systemevent.get
//
// 五道口悬挂链系统事件查询
func TaobaoWdkEquipmentConveyorSystemeventGet(ctx context.Context, clt *core.SDKClient, req *wdk.TaobaoWdkEquipmentConveyorSystemeventGetAPIRequest, resp *wdk.TaobaoWdkEquipmentConveyorSystemeventGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
