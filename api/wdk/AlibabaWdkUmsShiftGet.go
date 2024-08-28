package wdk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkUmsShiftGet 移库单获取
// alibaba.wdk.ums.shift.get
//
// 移库单获取
func AlibabaWdkUmsShiftGet(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaWdkUmsShiftGetAPIRequest, resp *wdk.AlibabaWdkUmsShiftGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
