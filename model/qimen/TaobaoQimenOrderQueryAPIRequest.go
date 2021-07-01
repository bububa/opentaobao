package qimen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoQimenOrderQueryAPIRequest
根据收件人信息查询交易单号接口 API请求
taobao.qimen.order.query

WMS 调用该接口，根据收件人信息查询平台交易订单号。 */
type TaobaoQimenOrderQueryAPIRequest struct {
	model.Params
	// request
	_request *TaobaoQimenOrderQueryRequest
}

// New
