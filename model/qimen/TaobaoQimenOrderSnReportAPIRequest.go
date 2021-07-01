package qimen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoQimenOrderSnReportAPIRequest
订单SN通知接口 API请求
taobao.qimen.order.sn.report

WMS调用奇门的接口,在出库、发货、入库等场景下，ERP和WMS之间同步操作的SN列表 */
type TaobaoQimenOrderSnReportAPIRequest struct {
	model.Params
	//
	_request *TaobaoQimenOrderSnReportRequest
}

// New
