package logistic

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/logistic"
)

// AliexpressLocalLogisticsOrderCreate create logistics order
// aliexpress.local.logistics.order.create
//
// create logistics order
func AliexpressLocalLogisticsOrderCreate(ctx context.Context, clt *core.SDKClient, req *logistic.AliexpressLocalLogisticsOrderCreateAPIRequest, resp *logistic.AliexpressLocalLogisticsOrderCreateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
