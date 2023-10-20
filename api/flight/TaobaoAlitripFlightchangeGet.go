package flight

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/flight"
)

// TaobaoAlitripFlightchangeGet 获取航变信息
// taobao.alitrip.flightchange.get
//
// 查询航变是为了两个目的，阿里旅行抓取到未确认航变的航变信息源时可以由代理确认航变，同时对于确认航变的航变信息也共享给代理人做本体业务使用。
func TaobaoAlitripFlightchangeGet(clt *core.SDKClient, req *flight.TaobaoAlitripFlightchangeGetAPIRequest, resp *flight.TaobaoAlitripFlightchangeGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
