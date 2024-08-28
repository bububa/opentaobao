package logistic

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/logistic"
)

// AliexpressLogisticsCreatewarehouseorder 创建线上物流订单
// aliexpress.logistics.createwarehouseorder
//
// 创建线上发货物流订单
func AliexpressLogisticsCreatewarehouseorder(ctx context.Context, clt *core.SDKClient, req *logistic.AliexpressLogisticsCreatewarehouseorderAPIRequest, resp *logistic.AliexpressLogisticsCreatewarehouseorderAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
