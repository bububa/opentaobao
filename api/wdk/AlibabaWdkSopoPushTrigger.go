package wdk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkSopoPushTrigger 猫超共享库存寄售sopo推送触发
// alibaba.wdk.sopo.push.trigger
//
// 猫超共享库存寄售sopo触发推送给商家
func AlibabaWdkSopoPushTrigger(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaWdkSopoPushTriggerAPIRequest, resp *wdk.AlibabaWdkSopoPushTriggerAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
