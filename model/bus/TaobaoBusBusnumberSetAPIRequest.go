package bus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoBusBusnumberSetAPIRequest
商家汽车票车次更新通知接口 API请求
taobao.bus.busnumber.set

商家汽车票车次更新后，调用该接口通知平台。 */
type TaobaoBusBusnumberSetAPIRequest struct {
	model.Params
	// 车次更新通知参数
	_pushParam *TopBusNumerPushRq
}

// New
