package flight

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/flight"
)

// TaobaoFliggyFlightAgentAuxproductPush 飞猪机票辅营商品投放
// taobao.fliggy.flight.agent.auxproduct.push
//
// 廉航辅营产品投放接口
func TaobaoFliggyFlightAgentAuxproductPush(ctx context.Context, clt *core.SDKClient, req *flight.TaobaoFliggyFlightAgentAuxproductPushAPIRequest, resp *flight.TaobaoFliggyFlightAgentAuxproductPushAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
