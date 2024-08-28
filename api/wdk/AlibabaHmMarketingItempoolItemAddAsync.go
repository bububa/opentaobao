package wdk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaHmMarketingItempoolItemAddAsync 商品池新增商品
// alibaba.hm.marketing.itempool.item.add.async
//
// 新分组模型下新增商品
func AlibabaHmMarketingItempoolItemAddAsync(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaHmMarketingItempoolItemAddAsyncAPIRequest, resp *wdk.AlibabaHmMarketingItempoolItemAddAsyncAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
