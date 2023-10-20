package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Wdkumsoutboundsortingcallbackforpulltask dps分货-任务拉取确定接口
// wdk.ums.outbound.sorting.callbackforpulltask
//
// dps分货-任务拉取确定接口
func Wdkumsoutboundsortingcallbackforpulltask(clt *core.SDKClient, req *wdk.WdkumsoutboundsortingcallbackforpulltaskAPIRequest, session string) (*wdk.WdkumsoutboundsortingcallbackforpulltaskAPIResponse, error) {
	var resp wdk.WdkumsoutboundsortingcallbackforpulltaskAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
