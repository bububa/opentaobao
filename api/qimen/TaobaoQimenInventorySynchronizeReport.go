package qimen

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/qimen"
)

// TaobaoQimenInventorySynchronizeReport 库存状态同步确认接口
// taobao.qimen.inventory.synchronize.report
//
// 库存状态同步确认接口
func TaobaoQimenInventorySynchronizeReport(clt *core.SDKClient, req *qimen.TaobaoQimenInventorySynchronizeReportAPIRequest, resp *qimen.TaobaoQimenInventorySynchronizeReportAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
