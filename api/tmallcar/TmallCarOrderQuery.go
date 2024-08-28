package tmallcar

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallcar"
)

// TmallCarOrderQuery 天猫汽车整车订单查询
// tmall.car.order.query
//
// 天猫汽车商家通过该接口查看整车订单信息
func TmallCarOrderQuery(ctx context.Context, clt *core.SDKClient, req *tmallcar.TmallCarOrderQueryAPIRequest, resp *tmallcar.TmallCarOrderQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
