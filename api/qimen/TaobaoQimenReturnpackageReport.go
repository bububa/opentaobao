package qimen

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/qimen"
)

// TaobaoQimenReturnpackageReport 退货包裹状态通知接口
// taobao.qimen.returnpackage.report
//
// 退货包裹状态通知接口
func TaobaoQimenReturnpackageReport(ctx context.Context, clt *core.SDKClient, req *qimen.TaobaoQimenReturnpackageReportAPIRequest, resp *qimen.TaobaoQimenReturnpackageReportAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
