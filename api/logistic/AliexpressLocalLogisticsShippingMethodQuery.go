package logistic

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/logistic"
)

// AliexpressLocalLogisticsShippingMethodQuery query shipping method
// aliexpress.local.logistics.shipping.method.query
//
// query shipping method
func AliexpressLocalLogisticsShippingMethodQuery(ctx context.Context, clt *core.SDKClient, req *logistic.AliexpressLocalLogisticsShippingMethodQueryAPIRequest, resp *logistic.AliexpressLocalLogisticsShippingMethodQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
