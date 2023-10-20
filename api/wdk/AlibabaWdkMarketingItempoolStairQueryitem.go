package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkMarketingItempoolStairQueryitem 换购商品查询
// alibaba.wdk.marketing.itempool.stair.queryitem
//
// 换购商品查询
func AlibabaWdkMarketingItempoolStairQueryitem(clt *core.SDKClient, req *wdk.AlibabaWdkMarketingItempoolStairQueryitemAPIRequest, resp *wdk.AlibabaWdkMarketingItempoolStairQueryitemAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
