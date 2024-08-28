package tblogistics

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tblogistics"
)

// TaobaoLogisticsOnlineCancel 取消物流订单接口
// taobao.logistics.online.cancel
//
// 调此接口取消发货的订单，重新选择物流公司发货。前提是物流公司未揽收货物。对未发货和已经被物流公司揽收的物流订单，是不能取消的。
func TaobaoLogisticsOnlineCancel(ctx context.Context, clt *core.SDKClient, req *tblogistics.TaobaoLogisticsOnlineCancelAPIRequest, resp *tblogistics.TaobaoLogisticsOnlineCancelAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
