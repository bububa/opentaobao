package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkMarketingItempoolStairRemoveitem 删除换购活动商品
// alibaba.wdk.marketing.itempool.stair.removeitem
//
// 删除换购商品
func AlibabaWdkMarketingItempoolStairRemoveitem(clt *core.SDKClient, req *wdk.AlibabaWdkMarketingItempoolStairRemoveitemAPIRequest, resp *wdk.AlibabaWdkMarketingItempoolStairRemoveitemAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
