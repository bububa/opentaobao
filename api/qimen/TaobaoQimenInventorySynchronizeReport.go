package qimen

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/qimen"
)

// TaobaoQimenInventorySynchronizeReport 库存状态同步确认接口
// taobao.qimen.inventory.synchronize.report
//
// 库存状态同步确认接口
func TaobaoQimenInventorySynchronizeReport(ctx context.Context, clt *core.SDKClient, req *qimen.TaobaoQimenInventorySynchronizeReportAPIRequest, resp *qimen.TaobaoQimenInventorySynchronizeReportAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
