package wdk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkMarketingPrice 促销价签服务
// alibaba.wdk.marketing.price
//
// 获取营销-促销商品中的实时价格
func AlibabaWdkMarketingPrice(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaWdkMarketingPriceAPIRequest, resp *wdk.AlibabaWdkMarketingPriceAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
