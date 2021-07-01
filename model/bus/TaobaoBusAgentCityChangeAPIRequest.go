package bus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoBusAgentCityChangeAPIRequest
城市变更 API请求
taobao.bus.agent.city.change

代理商通知城市变更，比如可售变为不可售等 */
type TaobaoBusAgentCityChangeAPIRequest struct {
	model.Params
	// 城市变更请求对象
	_paramCityChangeRQ *CityChangeRq
}

// New
