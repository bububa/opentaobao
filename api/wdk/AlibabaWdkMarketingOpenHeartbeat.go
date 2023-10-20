package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkMarketingOpenHeartbeat 心跳服务【10s一次】
// alibaba.wdk.marketing.open.heartbeat
//
// 商家数据同步心跳服务
func AlibabaWdkMarketingOpenHeartbeat(clt *core.SDKClient, req *wdk.AlibabaWdkMarketingOpenHeartbeatAPIRequest, resp *wdk.AlibabaWdkMarketingOpenHeartbeatAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
