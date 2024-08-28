package wdk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// WdkUmsOutboundSortingUserquery dps-查询分货作业人员信息
// wdk.ums.outbound.sorting.userquery
//
// dps-查询分货作业人员信息
func WdkUmsOutboundSortingUserquery(ctx context.Context, clt *core.SDKClient, req *wdk.WdkUmsOutboundSortingUserqueryAPIRequest, resp *wdk.WdkUmsOutboundSortingUserqueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
