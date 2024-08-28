package aedropshiper

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aedropshiper"
)

// AliexpressLogisticsBuyerFreightCalculate 提供给买家使用的运费计算接口
// aliexpress.logistics.buyer.freight.calculate
//
// 提供给买家使用的运费计算接口
func AliexpressLogisticsBuyerFreightCalculate(ctx context.Context, clt *core.SDKClient, req *aedropshiper.AliexpressLogisticsBuyerFreightCalculateAPIRequest, resp *aedropshiper.AliexpressLogisticsBuyerFreightCalculateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
