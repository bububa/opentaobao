package trade

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/trade"
)

// TaobaoXhotelDistributionOrderDetailSearch 分销渠道订单详情查询
// taobao.xhotel.distribution.order.detail.search
//
// 该接口用于提供酒店分销渠道的订单详情查询
func TaobaoXhotelDistributionOrderDetailSearch(ctx context.Context, clt *core.SDKClient, req *trade.TaobaoXhotelDistributionOrderDetailSearchAPIRequest, resp *trade.TaobaoXhotelDistributionOrderDetailSearchAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
