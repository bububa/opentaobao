package qimen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoQimenStockchangeReportAPIRequest
库存异动通知接口 API请求
taobao.qimen.stockchange.report

WMS调用奇门的接口,将库存异动信息信息回传给ERP */
type TaobaoQimenStockchangeReportAPIRequest struct {
	model.Params
	//
	_request *StockChangeReportRequest
}

// New
