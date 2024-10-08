package wdk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// TaobaoWdkEquipmentConveyorStatusconfirm 悬挂链状态回传确认
// taobao.wdk.equipment.conveyor.statusconfirm
//
// 悬挂链状态回传确认
func TaobaoWdkEquipmentConveyorStatusconfirm(ctx context.Context, clt *core.SDKClient, req *wdk.TaobaoWdkEquipmentConveyorStatusconfirmAPIRequest, resp *wdk.TaobaoWdkEquipmentConveyorStatusconfirmAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
