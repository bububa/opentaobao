package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaHmMarketingItemdiscountRemoveitem 移除报名的商品
// alibaba.hm.marketing.itemdiscount.removeitem
//
// 将报名特价活动的商品从特价活动中移除
func AlibabaHmMarketingItemdiscountRemoveitem(clt *core.SDKClient, req *wdk.AlibabaHmMarketingItemdiscountRemoveitemAPIRequest, session string) (*wdk.AlibabaHmMarketingItemdiscountRemoveitemAPIResponse, error) {
	var resp wdk.AlibabaHmMarketingItemdiscountRemoveitemAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
