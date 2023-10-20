package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkMarketingOpenDarunfaActivitySync 活动数据同步
// alibaba.wdk.marketing.open.darunfa.activity.sync
//
// 大润发活动数据同步
func AlibabaWdkMarketingOpenDarunfaActivitySync(clt *core.SDKClient, req *wdk.AlibabaWdkMarketingOpenDarunfaActivitySyncAPIRequest, resp *wdk.AlibabaWdkMarketingOpenDarunfaActivitySyncAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
