package wdk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkMarketingItempoolItemAddAsync 商品池新增商品
// alibaba.wdk.marketing.itempool.item.add.async
//
// 新分组模型下新增商品
func AlibabaWdkMarketingItempoolItemAddAsync(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaWdkMarketingItempoolItemAddAsyncAPIRequest, resp *wdk.AlibabaWdkMarketingItempoolItemAddAsyncAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
