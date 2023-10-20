package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaHmMarketingItemdiscountQueryitems 查询特价商品
// alibaba.hm.marketing.itemdiscount.queryitems
//
// 查询参加特价活动的商品优惠详情
func AlibabaHmMarketingItemdiscountQueryitems(clt *core.SDKClient, req *wdk.AlibabaHmMarketingItemdiscountQueryitemsAPIRequest, resp *wdk.AlibabaHmMarketingItemdiscountQueryitemsAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
