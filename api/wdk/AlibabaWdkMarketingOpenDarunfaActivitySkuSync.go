package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkMarketingOpenDarunfaActivitySkuSync 营销商品数据同步
// alibaba.wdk.marketing.open.darunfa.activity.sku.sync
//
// 大润发营销商品数据同步
func AlibabaWdkMarketingOpenDarunfaActivitySkuSync(clt *core.SDKClient, req *wdk.AlibabaWdkMarketingOpenDarunfaActivitySkuSyncAPIRequest, resp *wdk.AlibabaWdkMarketingOpenDarunfaActivitySkuSyncAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
