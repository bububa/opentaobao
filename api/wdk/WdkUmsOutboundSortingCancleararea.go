package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Wdkumsoutboundsortingcancleararea dps分货-是否能够清场
// wdk.ums.outbound.sorting.cancleararea
//
// dps分货-是否能够清场
func Wdkumsoutboundsortingcancleararea(clt *core.SDKClient, req *wdk.WdkumsoutboundsortingcanclearareaAPIRequest, session string) (*wdk.WdkumsoutboundsortingcanclearareaAPIResponse, error) {
	var resp wdk.WdkumsoutboundsortingcanclearareaAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
