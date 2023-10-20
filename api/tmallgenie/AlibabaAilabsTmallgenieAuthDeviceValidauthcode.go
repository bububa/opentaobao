package tmallgenie

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallgenie"
)

// Alibabaailabstmallgenieauthdevicevalidauthcode 根据authcode查询绑定结果
// alibaba.ailabs.tmallgenie.auth.device.validauthcode
//
// 根据authcode查询绑定结果
func Alibabaailabstmallgenieauthdevicevalidauthcode(clt *core.SDKClient, req *tmallgenie.AlibabaailabstmallgenieauthdevicevalidauthcodeAPIRequest, session string) (*tmallgenie.AlibabaailabstmallgenieauthdevicevalidauthcodeAPIResponse, error) {
	var resp tmallgenie.AlibabaailabstmallgenieauthdevicevalidauthcodeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
