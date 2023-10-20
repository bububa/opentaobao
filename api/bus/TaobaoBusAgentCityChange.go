package bus

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/bus"
)

// Taobaobusagentcitychange 城市变更
// taobao.bus.agent.city.change
//
// 代理商通知城市变更，比如可售变为不可售等
func Taobaobusagentcitychange(clt *core.SDKClient, req *bus.TaobaobusagentcitychangeAPIRequest, session string) (*bus.TaobaobusagentcitychangeAPIResponse, error) {
	var resp bus.TaobaobusagentcitychangeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
