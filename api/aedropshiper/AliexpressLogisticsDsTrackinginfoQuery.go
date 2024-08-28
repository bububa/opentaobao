package aedropshiper

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aedropshiper"
)

// AliexpressLogisticsDsTrackinginfoQuery 查询物流追踪信息
// aliexpress.logistics.ds.trackinginfo.query
//
// Dropshipper查询物流追踪信息
func AliexpressLogisticsDsTrackinginfoQuery(ctx context.Context, clt *core.SDKClient, req *aedropshiper.AliexpressLogisticsDsTrackinginfoQueryAPIRequest, resp *aedropshiper.AliexpressLogisticsDsTrackinginfoQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
