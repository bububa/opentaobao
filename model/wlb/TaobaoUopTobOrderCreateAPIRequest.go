package wlb

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoUopTobOrderCreateAPIRequest
ToB仓储发货 API请求
taobao.uop.tob.order.create

ToB仓储发货 */
type TaobaoUopTobOrderCreateAPIRequest struct {
	model.Params
	// ERP出库对象
	_deliveryOrder *DeliveryOrder
}

// New
