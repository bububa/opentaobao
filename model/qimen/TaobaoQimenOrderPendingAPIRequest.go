package qimen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoQimenOrderPendingAPIRequest
单据挂起（恢复）接口 API请求
taobao.qimen.order.pending

ERP调用奇门的接口,挂起某些创建的单据;场景介绍：ERP主动发起挂起（恢复）某些创建的单据，如入库单、出库单、退货单等 */
type TaobaoQimenOrderPendingAPIRequest struct {
	model.Params
	//
	_request *OrderPendingRequest
}

// New
