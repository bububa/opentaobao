package alilabs

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alilabs"
)

// Alibabaailabstmallgenieauthgetcode 获取token
// alibaba.ailabs.tmallgenie.auth.getcode
//
// 获取天猫精灵authCode
func Alibabaailabstmallgenieauthgetcode(clt *core.SDKClient, req *alilabs.AlibabaailabstmallgenieauthgetcodeAPIRequest, session string) (*alilabs.AlibabaailabstmallgenieauthgetcodeAPIResponse, error) {
	var resp alilabs.AlibabaailabstmallgenieauthgetcodeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
