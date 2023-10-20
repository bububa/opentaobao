package tmallgenie

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallgenie"
)

// AlibabaAilabsTmallgenieSdkDeviceIssupportsdk 是否支持云回看新SDK
// alibaba.ailabs.tmallgenie.sdk.device.issupportsdk
//
// 是否支持云回看新SDK
func AlibabaAilabsTmallgenieSdkDeviceIssupportsdk(clt *core.SDKClient, req *tmallgenie.AlibabaAilabsTmallgenieSdkDeviceIssupportsdkAPIRequest, resp *tmallgenie.AlibabaAilabsTmallgenieSdkDeviceIssupportsdkAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
