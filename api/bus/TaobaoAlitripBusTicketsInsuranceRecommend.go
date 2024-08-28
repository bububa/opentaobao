package bus

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/bus"
)

// TaobaoAlitripBusTicketsInsuranceRecommend 汽车票保险推荐
// taobao.alitrip.bus.tickets.insurance.recommend
//
// 获取推荐保险内容
func TaobaoAlitripBusTicketsInsuranceRecommend(ctx context.Context, clt *core.SDKClient, req *bus.TaobaoAlitripBusTicketsInsuranceRecommendAPIRequest, resp *bus.TaobaoAlitripBusTicketsInsuranceRecommendAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
