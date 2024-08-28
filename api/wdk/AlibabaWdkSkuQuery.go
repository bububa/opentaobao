package wdk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkSkuQuery 查询商品
// alibaba.wdk.sku.query
//
// 查询商品
func AlibabaWdkSkuQuery(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaWdkSkuQueryAPIRequest, resp *wdk.AlibabaWdkSkuQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
