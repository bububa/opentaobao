package bus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoBusRefundfeeGetAPIRequest
查询退票费用明细 API请求
taobao.bus.refundfee.get

查询退票的费用信息 */
type TaobaoBusRefundfeeGetAPIRequest struct {
	model.Params
	// 飞猪订单号
	_aliTripOrderId string
	// 飞猪子订单号
	_subOrderIds []int64
}

// New
