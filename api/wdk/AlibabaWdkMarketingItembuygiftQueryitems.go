package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkMarketingItembuygiftQueryitems 查询买赠活动下的商品
// alibaba.wdk.marketing.itembuygift.queryitems
//
// 查询买赠活动下的商品
func AlibabaWdkMarketingItembuygiftQueryitems(clt *core.SDKClient, req *wdk.AlibabaWdkMarketingItembuygiftQueryitemsAPIRequest, resp *wdk.AlibabaWdkMarketingItembuygiftQueryitemsAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
