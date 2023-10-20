package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// WdkUmsSortingFullContainer dps分货-满箱
// wdk.ums.sorting.full.container
//
// dps分货-满箱
func WdkUmsSortingFullContainer(clt *core.SDKClient, req *wdk.WdkUmsSortingFullContainerAPIRequest, resp *wdk.WdkUmsSortingFullContainerAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
