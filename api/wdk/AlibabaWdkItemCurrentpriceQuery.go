package wdk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkItemCurrentpriceQuery 查询商品当前价格
// alibaba.wdk.item.currentprice.query
//
// 通过渠道店id/sku编码/渠道查询商品当前价格，一次请求商品数量&lt;=20,返回结果key为skuCode
func AlibabaWdkItemCurrentpriceQuery(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaWdkItemCurrentpriceQueryAPIRequest, resp *wdk.AlibabaWdkItemCurrentpriceQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
