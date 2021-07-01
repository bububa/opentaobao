package qimen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoQimenOrderCancelAPIRequest
单据取消接口 API请求
taobao.qimen.order.cancel

ERP调用奇门的接口,取消创建单据操作。场景介绍：ERP主动发起取消某些创建的单据。如入库单、出库单、退货单等；所有的场景 */
type TaobaoQimenOrderCancelAPIRequest struct {
	model.Params
	//
	_request *OrderCancelRequest
}

// New
