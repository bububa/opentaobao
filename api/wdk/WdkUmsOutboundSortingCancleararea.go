package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// WdkUmsOutboundSortingCancleararea dps分货-是否能够清场
// wdk.ums.outbound.sorting.cancleararea
//
// dps分货-是否能够清场
func WdkUmsOutboundSortingCancleararea(clt *core.SDKClient, req *wdk.WdkUmsOutboundSortingCanclearareaAPIRequest, resp *wdk.WdkUmsOutboundSortingCanclearareaAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
