package tmallgenie

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallgenie"
)

// AlibabaAilabsTmallgenieSdkDeviceIssupportsdk 是否支持云回看新SDK
// alibaba.ailabs.tmallgenie.sdk.device.issupportsdk
//
// 是否支持云回看新SDK
func AlibabaAilabsTmallgenieSdkDeviceIssupportsdk(clt *core.SDKClient, req *tmallgenie.AlibabaAilabsTmallgenieSdkDeviceIssupportsdkAPIRequest, session string) (*tmallgenie.AlibabaAilabsTmallgenieSdkDeviceIssupportsdkAPIResponse, error) {
	var resp tmallgenie.AlibabaAilabsTmallgenieSdkDeviceIssupportsdkAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
