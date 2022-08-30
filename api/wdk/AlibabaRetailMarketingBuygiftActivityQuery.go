package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaRetailMarketingBuygiftActivityQuery 查询单品买赠活动【同城零售】
// alibaba.retail.marketing.buygift.activity.query
//
// 查询单品买赠活动【同城零售】
func AlibabaRetailMarketingBuygiftActivityQuery(clt *core.SDKClient, req *wdk.AlibabaRetailMarketingBuygiftActivityQueryAPIRequest, session string) (*wdk.AlibabaRetailMarketingBuygiftActivityQueryAPIResponse, error) {
	var resp wdk.AlibabaRetailMarketingBuygiftActivityQueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
