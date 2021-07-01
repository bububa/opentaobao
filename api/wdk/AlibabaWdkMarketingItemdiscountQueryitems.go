package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

/* AlibabaWdkMarketingItemdiscountQueryitems
查询特价商品
alibaba.wdk.marketing.itemdiscount.queryitems

查询参加特价活动的商品优惠详情 */
func AlibabaWdkMarketingItemdiscountQueryitems(clt *core.SDKClient, req *wdk.AlibabaWdkMarketingItemdiscountQueryitemsAPIRequest, session string) (*wdk.AlibabaWdkMarketingItemdiscountQueryitemsAPIResponse, error) {
	var resp wdk.AlibabaWdkMarketingItemdiscountQueryitemsAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
