package lstlogistics

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaLstShiporderCancelAPIRequest
零售通发货单取消 API请求
alibaba.lst.shiporder.cancel

通过该接口可以取消零售通运保保发货单，并处理相关业务流程。 */
type AlibabaLstShiporderCancelAPIRequest struct {
	model.Params
	// 取消原因
	_reason string
	// 订单号
	_outOrderId string
	// 需要退款的明细ID
	_detailOrderIds []string
}

// New
