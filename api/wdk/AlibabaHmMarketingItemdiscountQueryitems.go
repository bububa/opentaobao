package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaHmMarketingItemdiscountQueryitems 查询特价商品
// alibaba.hm.marketing.itemdiscount.queryitems
//
// 查询参加特价活动的商品优惠详情
func AlibabaHmMarketingItemdiscountQueryitems(clt *core.SDKClient, req *wdk.AlibabaHmMarketingItemdiscountQueryitemsAPIRequest, session string) (*wdk.AlibabaHmMarketingItemdiscountQueryitemsAPIResponse, error) {
	var resp wdk.AlibabaHmMarketingItemdiscountQueryitemsAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
