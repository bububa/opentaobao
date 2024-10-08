package wdk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkBmStockPublish 品牌营销涉及到的商品的库存同步接口
// alibaba.wdk.bm.stock.publish
//
// 用于操作sku的库存
func AlibabaWdkBmStockPublish(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaWdkBmStockPublishAPIRequest, resp *wdk.AlibabaWdkBmStockPublishAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
