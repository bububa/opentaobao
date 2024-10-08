package wdk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// TaobaoWdkEquipmentConveyorExceptionslidewaylogGet 异常通道日志查询
// taobao.wdk.equipment.conveyor.exceptionslidewaylog.get
//
// 五道口悬挂链异常通道事件查询
func TaobaoWdkEquipmentConveyorExceptionslidewaylogGet(ctx context.Context, clt *core.SDKClient, req *wdk.TaobaoWdkEquipmentConveyorExceptionslidewaylogGetAPIRequest, resp *wdk.TaobaoWdkEquipmentConveyorExceptionslidewaylogGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
