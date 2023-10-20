package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// WdkUmsOutboundSortingCallbackTaskdetail dps分货，明细回传
// wdk.ums.outbound.sorting.callback.taskdetail
//
// dps分货-分货明细回传
func WdkUmsOutboundSortingCallbackTaskdetail(clt *core.SDKClient, req *wdk.WdkUmsOutboundSortingCallbackTaskdetailAPIRequest, resp *wdk.WdkUmsOutboundSortingCallbackTaskdetailAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
