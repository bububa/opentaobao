package qimen

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/qimen"
)

// TaobaoQimenSnReport 发货单SN通知接口
// taobao.qimen.sn.report
//
// WMS调用奇门的接口,在仓库出库单后, 把SN信息回传给ERP
func TaobaoQimenSnReport(ctx context.Context, clt *core.SDKClient, req *qimen.TaobaoQimenSnReportAPIRequest, resp *qimen.TaobaoQimenSnReportAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
