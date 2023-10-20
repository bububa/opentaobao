package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Wdkumsoutboundsortingcallbacktaskdetail dps分货，明细回传
// wdk.ums.outbound.sorting.callback.taskdetail
//
// dps分货-分货明细回传
func Wdkumsoutboundsortingcallbacktaskdetail(clt *core.SDKClient, req *wdk.WdkumsoutboundsortingcallbacktaskdetailAPIRequest, session string) (*wdk.WdkumsoutboundsortingcallbacktaskdetailAPIResponse, error) {
	var resp wdk.WdkumsoutboundsortingcallbacktaskdetailAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
