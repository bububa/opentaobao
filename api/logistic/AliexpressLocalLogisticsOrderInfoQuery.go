package logistic

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/logistic"
)

// AliexpressLocalLogisticsOrderInfoQuery query order details
// aliexpress.local.logistics.order.info.query
//
// query order details
func AliexpressLocalLogisticsOrderInfoQuery(ctx context.Context, clt *core.SDKClient, req *logistic.AliexpressLocalLogisticsOrderInfoQueryAPIRequest, resp *logistic.AliexpressLocalLogisticsOrderInfoQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
