package car

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/car"
)

// TaobaoAlitripCarDriverStatusUpdate 司机服务状态更新接口
// taobao.alitrip.car.driver.status.update
//
// 飞猪用车业务回调接口，用于服务商实时回传更新司机当前服务状态
func TaobaoAlitripCarDriverStatusUpdate(ctx context.Context, clt *core.SDKClient, req *car.TaobaoAlitripCarDriverStatusUpdateAPIRequest, resp *car.TaobaoAlitripCarDriverStatusUpdateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
