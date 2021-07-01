package flight

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoAlitripFlightchangeGetAPIRequest
获取航变信息 API请求
taobao.alitrip.flightchange.get

查询航变是为了两个目的，阿里旅行抓取到未确认航变的航变信息源时可以由代理确认航变，同时对于确认航变的航变信息也共享给代理人做本体业务使用。 */
type TaobaoAlitripFlightchangeGetAPIRequest struct {
	model.Params
	// 查询信息封装
	_searchOption *FlightChangeDataQueryOption
}

// New
