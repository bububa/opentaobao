package shop

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/shop"
)

// AlibabaKoubeishopsPropertyGet 口碑店铺列表推荐
// alibaba.koubeishops.property.get
//
// 推荐用户附近的美食门店
func AlibabaKoubeishopsPropertyGet(ctx context.Context, clt *core.SDKClient, req *shop.AlibabaKoubeishopsPropertyGetAPIRequest, resp *shop.AlibabaKoubeishopsPropertyGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
