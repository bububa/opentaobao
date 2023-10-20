package tmallgenie

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallgenie"
)

// AlibabaAilabsTmallgenieAuthDeviceGetcode 获取authcode
// alibaba.ailabs.tmallgenie.auth.device.getcode
//
// 获取绑定的authcode
func AlibabaAilabsTmallgenieAuthDeviceGetcode(clt *core.SDKClient, req *tmallgenie.AlibabaAilabsTmallgenieAuthDeviceGetcodeAPIRequest, resp *tmallgenie.AlibabaAilabsTmallgenieAuthDeviceGetcodeAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
