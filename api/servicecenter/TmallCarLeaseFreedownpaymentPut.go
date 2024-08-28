package servicecenter

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/servicecenter"
)

// TmallCarLeaseFreedownpaymentPut 同步直租车免首付商品活动信息
// tmall.car.lease.freedownpayment.put
//
// 汽车行业直租车免首付需求中，用与对商品打标，活动范围设置，在消费者端商品详情页、订单等环节透出，表示该商品为直租免首付商品。
func TmallCarLeaseFreedownpaymentPut(ctx context.Context, clt *core.SDKClient, req *servicecenter.TmallCarLeaseFreedownpaymentPutAPIRequest, resp *servicecenter.TmallCarLeaseFreedownpaymentPutAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
