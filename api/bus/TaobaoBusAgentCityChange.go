package bus

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/bus"
)

/* TaobaoBusAgentCityChange
城市变更
taobao.bus.agent.city.change

代理商通知城市变更，比如可售变为不可售等 */
func TaobaoBusAgentCityChange(clt *core.SDKClient, req *bus.TaobaoBusAgentCityChangeAPIRequest, session string) (*bus.TaobaoBusAgentCityChangeAPIResponse, error) {
	var resp bus.TaobaoBusAgentCityChangeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
