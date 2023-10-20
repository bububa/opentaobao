package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaHmMarketingItemdiscountQueryactivity 查找特价活动
// alibaba.hm.marketing.itemdiscount.queryactivity
//
// 查找特价活动
func AlibabaHmMarketingItemdiscountQueryactivity(clt *core.SDKClient, req *wdk.AlibabaHmMarketingItemdiscountQueryactivityAPIRequest, session string) (*wdk.AlibabaHmMarketingItemdiscountQueryactivityAPIResponse, error) {
	var resp wdk.AlibabaHmMarketingItemdiscountQueryactivityAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
