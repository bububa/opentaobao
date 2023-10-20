package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkMarketingItembuygiftRemoveitem 移除买赠活动下的商品。【注意，此接口暂不支持并发！】
// alibaba.wdk.marketing.itembuygift.removeitem
//
// 移除买赠活动下的商品。【注意，此接口暂不支持并发！】
func AlibabaWdkMarketingItembuygiftRemoveitem(clt *core.SDKClient, req *wdk.AlibabaWdkMarketingItembuygiftRemoveitemAPIRequest, resp *wdk.AlibabaWdkMarketingItembuygiftRemoveitemAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
