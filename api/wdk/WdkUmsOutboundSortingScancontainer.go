package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// WdkUmsOutboundSortingScancontainer dps分货-扫描分货容器判断是否可用
// wdk.ums.outbound.sorting.scancontainer
//
// dps分货-扫描分货容器判断是否可用
func WdkUmsOutboundSortingScancontainer(clt *core.SDKClient, req *wdk.WdkUmsOutboundSortingScancontainerAPIRequest, session string) (*wdk.WdkUmsOutboundSortingScancontainerAPIResponse, error) {
	var resp wdk.WdkUmsOutboundSortingScancontainerAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
