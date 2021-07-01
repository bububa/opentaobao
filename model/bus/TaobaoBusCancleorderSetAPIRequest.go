package bus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoBusCancleorderSetAPIRequest
取消订单 API请求
taobao.bus.cancleorder.set

取消订单 */
type TaobaoBusCancleorderSetAPIRequest struct {
	model.Params
	// 阿里订单号
	_aliOrderId string
}

// New
