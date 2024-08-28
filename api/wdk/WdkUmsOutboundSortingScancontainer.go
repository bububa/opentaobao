package wdk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// WdkUmsOutboundSortingScancontainer dps分货-扫描分货容器判断是否可用
// wdk.ums.outbound.sorting.scancontainer
//
// dps分货-扫描分货容器判断是否可用
func WdkUmsOutboundSortingScancontainer(ctx context.Context, clt *core.SDKClient, req *wdk.WdkUmsOutboundSortingScancontainerAPIRequest, resp *wdk.WdkUmsOutboundSortingScancontainerAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
