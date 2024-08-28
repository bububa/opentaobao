package wdk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// TaobaoWdkEquipmentConveyorContainerinfoGet 获取批次或波次中容器的信息
// taobao.wdk.equipment.conveyor.containerinfo.get
//
// 获取批次或波次中容器的信息
func TaobaoWdkEquipmentConveyorContainerinfoGet(ctx context.Context, clt *core.SDKClient, req *wdk.TaobaoWdkEquipmentConveyorContainerinfoGetAPIRequest, resp *wdk.TaobaoWdkEquipmentConveyorContainerinfoGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
