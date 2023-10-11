package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaHmMarketingItemdiscountAdditem 报名特价商品
// alibaba.hm.marketing.itemdiscount.additem
//
// 在商品特价活动中报名特价商品
func AlibabaHmMarketingItemdiscountAdditem(clt *core.SDKClient, req *wdk.AlibabaHmMarketingItemdiscountAdditemAPIRequest, session string) (*wdk.AlibabaHmMarketingItemdiscountAdditemAPIResponse, error) {
	var resp wdk.AlibabaHmMarketingItemdiscountAdditemAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
