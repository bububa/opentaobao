package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

/* AlibabaRetailMarketingItemdiscountActivityCreate
创建单品特价活动【同城零售】
alibaba.retail.marketing.itemdiscount.activity.create

同城零售单品特价活动创建 */
func AlibabaRetailMarketingItemdiscountActivityCreate(clt *core.SDKClient, req *wdk.AlibabaRetailMarketingItemdiscountActivityCreateAPIRequest, session string) (*wdk.AlibabaRetailMarketingItemdiscountActivityCreateAPIResponse, error) {
	var resp wdk.AlibabaRetailMarketingItemdiscountActivityCreateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
