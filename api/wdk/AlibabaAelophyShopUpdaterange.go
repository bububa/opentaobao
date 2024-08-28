package wdk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaAelophyShopUpdaterange 更新渠道店销售范围
// alibaba.aelophy.shop.updaterange
//
// 更新渠道店销售范围
func AlibabaAelophyShopUpdaterange(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaAelophyShopUpdaterangeAPIRequest, resp *wdk.AlibabaAelophyShopUpdaterangeAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
