package bus

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/bus"
)

// TaobaoBusAgentCityChange 城市变更
// taobao.bus.agent.city.change
//
// 代理商通知城市变更，比如可售变为不可售等
func TaobaoBusAgentCityChange(ctx context.Context, clt *core.SDKClient, req *bus.TaobaoBusAgentCityChangeAPIRequest, resp *bus.TaobaoBusAgentCityChangeAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
