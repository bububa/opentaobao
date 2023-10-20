package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaHmMarketingItembuygiftQueryitems 查询买赠活动下的商品
// alibaba.hm.marketing.itembuygift.queryitems
//
// 查询买赠活动下的商品
func AlibabaHmMarketingItembuygiftQueryitems(clt *core.SDKClient, req *wdk.AlibabaHmMarketingItembuygiftQueryitemsAPIRequest, resp *wdk.AlibabaHmMarketingItembuygiftQueryitemsAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
