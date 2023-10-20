package tmallgenie

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallgenie"
)

// Alibabaailabstmallgenieauthdevicegetcode 获取authcode
// alibaba.ailabs.tmallgenie.auth.device.getcode
//
// 获取绑定的authcode
func Alibabaailabstmallgenieauthdevicegetcode(clt *core.SDKClient, req *tmallgenie.AlibabaailabstmallgenieauthdevicegetcodeAPIRequest, session string) (*tmallgenie.AlibabaailabstmallgenieauthdevicegetcodeAPIResponse, error) {
	var resp tmallgenie.AlibabaailabstmallgenieauthdevicegetcodeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
