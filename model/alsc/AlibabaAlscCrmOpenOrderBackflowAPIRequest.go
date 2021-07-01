package alsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlscCrmOpenOrderBackflowAPIRequest
订单回流接口 API请求
alibaba.alsc.crm.open.order.backflow

回流isv订单接口 */
type AlibabaAlscCrmOpenOrderBackflowAPIRequest struct {
	model.Params
	// 入参
	_paramOrderBackflowOpenReq *OrderBackflowOpenReq
}

// New
