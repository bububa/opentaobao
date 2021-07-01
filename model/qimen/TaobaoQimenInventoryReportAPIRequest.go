package qimen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoQimenInventoryReportAPIRequest
库存盘点通知接口 API请求
taobao.qimen.inventory.report

WMS调用奇门的接口,将库存盘点情况回传ERP */
type TaobaoQimenInventoryReportAPIRequest struct {
	model.Params
	//
	_request *InventoryReportRequest
}

// New
