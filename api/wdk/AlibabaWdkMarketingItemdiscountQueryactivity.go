package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

/* AlibabaWdkMarketingItemdiscountQueryactivity
查找特价活动
alibaba.wdk.marketing.itemdiscount.queryactivity

查找特价活动 */
func AlibabaWdkMarketingItemdiscountQueryactivity(clt *core.SDKClient, req *wdk.AlibabaWdkMarketingItemdiscountQueryactivityAPIRequest, session string) (*wdk.AlibabaWdkMarketingItemdiscountQueryactivityAPIResponse, error) {
	var resp wdk.AlibabaWdkMarketingItemdiscountQueryactivityAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
