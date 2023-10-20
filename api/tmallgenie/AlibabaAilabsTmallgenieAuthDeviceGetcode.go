package tmallgenie

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallgenie"
)

// AlibabaAilabsTmallgenieAuthDeviceGetcode 获取authcode
// alibaba.ailabs.tmallgenie.auth.device.getcode
//
// 获取绑定的authcode
func AlibabaAilabsTmallgenieAuthDeviceGetcode(clt *core.SDKClient, req *tmallgenie.AlibabaAilabsTmallgenieAuthDeviceGetcodeAPIRequest, session string) (*tmallgenie.AlibabaAilabsTmallgenieAuthDeviceGetcodeAPIResponse, error) {
	var resp tmallgenie.AlibabaAilabsTmallgenieAuthDeviceGetcodeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
