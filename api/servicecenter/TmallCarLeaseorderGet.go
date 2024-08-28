package servicecenter

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/servicecenter"
)

// TmallCarLeaseorderGet 获取租赁订单信息
// tmall.car.leaseorder.get
//
// 卖家通过供销平台获取代销商的订单信息，但是部分情况下网商银行订单号获取不到，需要提供接口或者工具给卖家
func TmallCarLeaseorderGet(ctx context.Context, clt *core.SDKClient, req *servicecenter.TmallCarLeaseorderGetAPIRequest, resp *servicecenter.TmallCarLeaseorderGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
