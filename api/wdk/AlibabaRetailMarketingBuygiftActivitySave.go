package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaRetailMarketingBuygiftActivitySave 【同城零售】单品买赠活动保存
// alibaba.retail.marketing.buygift.activity.save
//
// 同城零售单品买赠活动保存
func AlibabaRetailMarketingBuygiftActivitySave(clt *core.SDKClient, req *wdk.AlibabaRetailMarketingBuygiftActivitySaveAPIRequest, session string) (*wdk.AlibabaRetailMarketingBuygiftActivitySaveAPIResponse, error) {
	var resp wdk.AlibabaRetailMarketingBuygiftActivitySaveAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
