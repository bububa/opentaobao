package qimen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoQimenSnReportAPIRequest
发货单SN通知接口 API请求
taobao.qimen.sn.report

WMS调用奇门的接口,在仓库出库单后, 把SN信息回传给ERP */
type TaobaoQimenSnReportAPIRequest struct {
	model.Params
	//
	_request *SnReportRequest
}

// New
