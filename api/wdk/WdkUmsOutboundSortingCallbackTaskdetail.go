package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// WdkUmsOutboundSortingCallbackTaskdetail dps分货，明细回传
// wdk.ums.outbound.sorting.callback.taskdetail
//
// dps分货-分货明细回传
func WdkUmsOutboundSortingCallbackTaskdetail(clt *core.SDKClient, req *wdk.WdkUmsOutboundSortingCallbackTaskdetailAPIRequest, session string) (*wdk.WdkUmsOutboundSortingCallbackTaskdetailAPIResponse, error) {
	var resp wdk.WdkUmsOutboundSortingCallbackTaskdetailAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
