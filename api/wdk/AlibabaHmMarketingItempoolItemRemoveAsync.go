package wdk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaHmMarketingItempoolItemRemoveAsync 商品池删除商品
// alibaba.hm.marketing.itempool.item.remove.async
//
// 新模型下删除商品
func AlibabaHmMarketingItempoolItemRemoveAsync(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaHmMarketingItempoolItemRemoveAsyncAPIRequest, resp *wdk.AlibabaHmMarketingItempoolItemRemoveAsyncAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
