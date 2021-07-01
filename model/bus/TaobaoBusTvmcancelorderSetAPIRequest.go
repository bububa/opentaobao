package bus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoBusTvmcancelorderSetAPIRequest
线下自助机未付款取消订单 API请求
taobao.bus.tvmcancelorder.set

自助机汽车票未付款取消订单 */
type TaobaoBusTvmcancelorderSetAPIRequest struct {
	model.Params
	// 飞猪订单号
	_alitripOrderId string
}

// New
