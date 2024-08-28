package qimen

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/qimen"
)

// TaobaoQimenOrderprocessReport 订单流水通知接口
// taobao.qimen.orderprocess.report
//
// taobao.qimen.orderprocess.report
func TaobaoQimenOrderprocessReport(ctx context.Context, clt *core.SDKClient, req *qimen.TaobaoQimenOrderprocessReportAPIRequest, resp *qimen.TaobaoQimenOrderprocessReportAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
