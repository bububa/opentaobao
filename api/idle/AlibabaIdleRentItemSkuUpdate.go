package idle

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/idle"
)

// AlibabaIdleRentItemSkuUpdate 更新/增加sku信息
// alibaba.idle.rent.item.sku.update
//
// 更新/增加sku信息
func AlibabaIdleRentItemSkuUpdate(ctx context.Context, clt *core.SDKClient, req *idle.AlibabaIdleRentItemSkuUpdateAPIRequest, resp *idle.AlibabaIdleRentItemSkuUpdateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
