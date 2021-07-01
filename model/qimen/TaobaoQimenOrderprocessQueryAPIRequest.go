package qimen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoQimenOrderprocessQueryAPIRequest
订单流水查询接口 API请求
taobao.qimen.orderprocess.query

ERP调用订单流水查询接口 */
type TaobaoQimenOrderprocessQueryAPIRequest struct {
	model.Params
	//
	_request *OrderProcessQueryRequest
}

// New
