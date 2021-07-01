package lstlogistics

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaLstShiporderQueryAPIRequest
零售通发货单查询 API请求
alibaba.lst.shiporder.query

通过该接口可以查询零售通运保保发货单，并处理相关业务流程。 */
type AlibabaLstShiporderQueryAPIRequest struct {
	model.Params
	// 零售通
	_source string
	// 订单
	_outOrderId string
}

// New
