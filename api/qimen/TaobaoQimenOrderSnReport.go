package qimen

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/qimen"
)

// TaobaoQimenOrderSnReport 订单SN通知接口
// taobao.qimen.order.sn.report
//
// WMS调用奇门的接口,在出库、发货、入库等场景下，ERP和WMS之间同步操作的SN列表
func TaobaoQimenOrderSnReport(ctx context.Context, clt *core.SDKClient, req *qimen.TaobaoQimenOrderSnReportAPIRequest, resp *qimen.TaobaoQimenOrderSnReportAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
