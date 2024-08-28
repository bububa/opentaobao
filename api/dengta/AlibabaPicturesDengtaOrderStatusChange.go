package dengta

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/dengta"
)

// AlibabaPicturesDengtaOrderStatusChange 天下秀订单状态变更通知
// alibaba.pictures.dengta.order.status.change
//
// 天下秀订单状态变更通知
func AlibabaPicturesDengtaOrderStatusChange(ctx context.Context, clt *core.SDKClient, req *dengta.AlibabaPicturesDengtaOrderStatusChangeAPIRequest, resp *dengta.AlibabaPicturesDengtaOrderStatusChangeAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
