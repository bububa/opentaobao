package tblogistics

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tblogistics"
)

// TaobaoLogisticsOrdersDetailGet 批量查询物流订单,返回详细信息
// taobao.logistics.orders.detail.get
//
// 查询物流订单的详细信息，涉及用户隐私字段。
func TaobaoLogisticsOrdersDetailGet(ctx context.Context, clt *core.SDKClient, req *tblogistics.TaobaoLogisticsOrdersDetailGetAPIRequest, resp *tblogistics.TaobaoLogisticsOrdersDetailGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
