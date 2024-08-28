package alihealth2

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealth2"
)

// AlibabaHealthNrCepOrderQuery 订单详情查询接口
// alibaba.health.nr.cep.order.query
//
// 订单详情查询接口
func AlibabaHealthNrCepOrderQuery(ctx context.Context, clt *core.SDKClient, req *alihealth2.AlibabaHealthNrCepOrderQueryAPIRequest, resp *alihealth2.AlibabaHealthNrCepOrderQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
