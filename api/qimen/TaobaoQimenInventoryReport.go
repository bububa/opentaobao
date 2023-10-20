package qimen

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/qimen"
)

// TaobaoQimenInventoryReport 库存盘点通知接口
// taobao.qimen.inventory.report
//
// WMS调用奇门的接口,将库存盘点情况回传ERP
func TaobaoQimenInventoryReport(clt *core.SDKClient, req *qimen.TaobaoQimenInventoryReportAPIRequest, resp *qimen.TaobaoQimenInventoryReportAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
