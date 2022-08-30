package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaRetailMarketingItemdiscountActivitySave 【同城零售】单品活动保存
// alibaba.retail.marketing.itemdiscount.activity.save
//
// 【同城零售】单品活动保存
func AlibabaRetailMarketingItemdiscountActivitySave(clt *core.SDKClient, req *wdk.AlibabaRetailMarketingItemdiscountActivitySaveAPIRequest, session string) (*wdk.AlibabaRetailMarketingItemdiscountActivitySaveAPIResponse, error) {
	var resp wdk.AlibabaRetailMarketingItemdiscountActivitySaveAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
