package bus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoBusBusnumberGetAPIRequest
汽车票车次查询 API请求
taobao.bus.busnumber.get

提供汽车票车次查询服务 */
type TaobaoBusBusnumberGetAPIRequest struct {
	model.Params
	// 车次查询入参
	_paramBusNumberSearchRQ *BusNumberSearchRq
}

// New
