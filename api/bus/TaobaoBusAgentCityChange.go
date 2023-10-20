package bus

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/bus"
)

// TaobaoBusAgentCityChange 城市变更
// taobao.bus.agent.city.change
//
// 代理商通知城市变更，比如可售变为不可售等
func TaobaoBusAgentCityChange(clt *core.SDKClient, req *bus.TaobaoBusAgentCityChangeAPIRequest, resp *bus.TaobaoBusAgentCityChangeAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
