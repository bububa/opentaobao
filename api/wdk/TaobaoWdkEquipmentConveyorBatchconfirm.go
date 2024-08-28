package wdk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// TaobaoWdkEquipmentConveyorBatchconfirm 五道口悬挂链信息批量确认
// taobao.wdk.equipment.conveyor.batchconfirm
//
// 批量消息确认
func TaobaoWdkEquipmentConveyorBatchconfirm(ctx context.Context, clt *core.SDKClient, req *wdk.TaobaoWdkEquipmentConveyorBatchconfirmAPIRequest, resp *wdk.TaobaoWdkEquipmentConveyorBatchconfirmAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
