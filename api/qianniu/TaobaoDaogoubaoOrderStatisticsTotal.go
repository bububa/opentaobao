package qianniu

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/qianniu"
)

// TaobaoDaogoubaoOrderStatisticsTotal 销售订单总额统计
// taobao.daogoubao.order.statistics.total
//
// 对接千牛端数字中心
func TaobaoDaogoubaoOrderStatisticsTotal(ctx context.Context, clt *core.SDKClient, req *qianniu.TaobaoDaogoubaoOrderStatisticsTotalAPIRequest, resp *qianniu.TaobaoDaogoubaoOrderStatisticsTotalAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
