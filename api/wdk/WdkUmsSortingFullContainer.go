package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Wdkumssortingfullcontainer dps分货-满箱
// wdk.ums.sorting.full.container
//
// dps分货-满箱
func Wdkumssortingfullcontainer(clt *core.SDKClient, req *wdk.WdkumssortingfullcontainerAPIRequest, session string) (*wdk.WdkumssortingfullcontainerAPIResponse, error) {
	var resp wdk.WdkumssortingfullcontainerAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
