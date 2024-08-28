package tmallgenie

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallgenie"
)

// AlibabaAilabsTmallgenieSdkDeviceIssupportsdk 是否支持云回看新SDK
// alibaba.ailabs.tmallgenie.sdk.device.issupportsdk
//
// 是否支持云回看新SDK
func AlibabaAilabsTmallgenieSdkDeviceIssupportsdk(ctx context.Context, clt *core.SDKClient, req *tmallgenie.AlibabaAilabsTmallgenieSdkDeviceIssupportsdkAPIRequest, resp *tmallgenie.AlibabaAilabsTmallgenieSdkDeviceIssupportsdkAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
