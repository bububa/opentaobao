package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkMarketingItempoolRemoveitem 移除商品池里面的商品
// alibaba.wdk.marketing.itempool.removeitem
//
// 移除商品池里面的商品
func AlibabaWdkMarketingItempoolRemoveitem(clt *core.SDKClient, req *wdk.AlibabaWdkMarketingItempoolRemoveitemAPIRequest, resp *wdk.AlibabaWdkMarketingItempoolRemoveitemAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
