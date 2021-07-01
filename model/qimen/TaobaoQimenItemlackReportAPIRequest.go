package qimen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoQimenItemlackReportAPIRequest
发货单缺货通知接口 API请求
taobao.qimen.itemlack.report

WMS调用奇门的接口,将商家在库某商品缺货的信息回传给ERP */
type TaobaoQimenItemlackReportAPIRequest struct {
	model.Params
	//
	_request *ItemLackReportRequest
}

// New
