package wdk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkReverseTimeslice 逆向取货时间片查询
// alibaba.wdk.reverse.timeslice
//
// 逆向取货时间片查询
func AlibabaWdkReverseTimeslice(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaWdkReverseTimesliceAPIRequest, resp *wdk.AlibabaWdkReverseTimesliceAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
