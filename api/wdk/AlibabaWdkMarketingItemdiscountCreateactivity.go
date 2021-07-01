package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

/* AlibabaWdkMarketingItemdiscountCreateactivity
创建商品特价活动
alibaba.wdk.marketing.itemdiscount.createactivity

创建商品特价活动 */
func AlibabaWdkMarketingItemdiscountCreateactivity(clt *core.SDKClient, req *wdk.AlibabaWdkMarketingItemdiscountCreateactivityAPIRequest, session string) (*wdk.AlibabaWdkMarketingItemdiscountCreateactivityAPIResponse, error) {
	var resp wdk.AlibabaWdkMarketingItemdiscountCreateactivityAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
