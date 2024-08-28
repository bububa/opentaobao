package wdk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// TaobaoWdkEquipmentConveyorWcsbtocContainerassignedtoconveyor 容器被预分拣器分配到悬挂链
// taobao.wdk.equipment.conveyor.wcsbtoc.containerassignedtoconveyor
//
// 容器被预分拣器分配到悬挂链
func TaobaoWdkEquipmentConveyorWcsbtocContainerassignedtoconveyor(ctx context.Context, clt *core.SDKClient, req *wdk.TaobaoWdkEquipmentConveyorWcsbtocContainerassignedtoconveyorAPIRequest, resp *wdk.TaobaoWdkEquipmentConveyorWcsbtocContainerassignedtoconveyorAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
