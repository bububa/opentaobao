package wdk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// TaobaoWdkEquipmentConveyorConveyorinfoGet 获取五道口悬挂链信息
// taobao.wdk.equipment.conveyor.conveyorinfo.get
//
// 获取五道口悬挂链信息
func TaobaoWdkEquipmentConveyorConveyorinfoGet(ctx context.Context, clt *core.SDKClient, req *wdk.TaobaoWdkEquipmentConveyorConveyorinfoGetAPIRequest, resp *wdk.TaobaoWdkEquipmentConveyorConveyorinfoGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
