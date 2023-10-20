package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// WdkUmsOutboundSortingCallbackforpulltask dps分货-任务拉取确定接口
// wdk.ums.outbound.sorting.callbackforpulltask
//
// dps分货-任务拉取确定接口
func WdkUmsOutboundSortingCallbackforpulltask(clt *core.SDKClient, req *wdk.WdkUmsOutboundSortingCallbackforpulltaskAPIRequest, resp *wdk.WdkUmsOutboundSortingCallbackforpulltaskAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
