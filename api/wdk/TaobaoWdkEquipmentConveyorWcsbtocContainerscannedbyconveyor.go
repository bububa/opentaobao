package wdk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// TaobaoWdkEquipmentConveyorWcsbtocContainerscannedbyconveyor 容器被悬挂链扫描
// taobao.wdk.equipment.conveyor.wcsbtoc.containerscannedbyconveyor
//
// 容器被悬挂链扫描
func TaobaoWdkEquipmentConveyorWcsbtocContainerscannedbyconveyor(ctx context.Context, clt *core.SDKClient, req *wdk.TaobaoWdkEquipmentConveyorWcsbtocContainerscannedbyconveyorAPIRequest, resp *wdk.TaobaoWdkEquipmentConveyorWcsbtocContainerscannedbyconveyorAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
