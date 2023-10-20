package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkMarketingItempoolItemRemoveAsync 商品池删除商品
// alibaba.wdk.marketing.itempool.item.remove.async
//
// 新模型下删除商品
func AlibabaWdkMarketingItempoolItemRemoveAsync(clt *core.SDKClient, req *wdk.AlibabaWdkMarketingItempoolItemRemoveAsyncAPIRequest, resp *wdk.AlibabaWdkMarketingItempoolItemRemoveAsyncAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
