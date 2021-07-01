package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

/* AlibabaRetailMarketingBuygiftActivityCreate
创建买赠活动
alibaba.retail.marketing.buygift.activity.create

同城供给买赠活动创建 */
func AlibabaRetailMarketingBuygiftActivityCreate(clt *core.SDKClient, req *wdk.AlibabaRetailMarketingBuygiftActivityCreateAPIRequest, session string) (*wdk.AlibabaRetailMarketingBuygiftActivityCreateAPIResponse, error) {
	var resp wdk.AlibabaRetailMarketingBuygiftActivityCreateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
