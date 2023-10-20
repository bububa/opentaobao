package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaRetailMarketingBuygiftActivityQuery 查询单品买赠活动【同城零售】
// alibaba.retail.marketing.buygift.activity.query
//
// 查询单品买赠活动【同城零售】
func AlibabaRetailMarketingBuygiftActivityQuery(clt *core.SDKClient, req *wdk.AlibabaRetailMarketingBuygiftActivityQueryAPIRequest, resp *wdk.AlibabaRetailMarketingBuygiftActivityQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
