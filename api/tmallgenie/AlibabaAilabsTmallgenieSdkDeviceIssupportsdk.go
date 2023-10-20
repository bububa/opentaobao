package tmallgenie

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallgenie"
)

// Alibabaailabstmallgeniesdkdeviceissupportsdk 是否支持云回看新SDK
// alibaba.ailabs.tmallgenie.sdk.device.issupportsdk
//
// 是否支持云回看新SDK
func Alibabaailabstmallgeniesdkdeviceissupportsdk(clt *core.SDKClient, req *tmallgenie.AlibabaailabstmallgeniesdkdeviceissupportsdkAPIRequest, session string) (*tmallgenie.AlibabaailabstmallgeniesdkdeviceissupportsdkAPIResponse, error) {
	var resp tmallgenie.AlibabaailabstmallgeniesdkdeviceissupportsdkAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
