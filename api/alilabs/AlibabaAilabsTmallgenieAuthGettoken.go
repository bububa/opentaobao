package alilabs

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alilabs"
)

// Alibabaailabstmallgenieauthgettoken 设备授权
// alibaba.ailabs.tmallgenie.auth.gettoken
//
// 获取设备授权码
func Alibabaailabstmallgenieauthgettoken(clt *core.SDKClient, req *alilabs.AlibabaailabstmallgenieauthgettokenAPIRequest, session string) (*alilabs.AlibabaailabstmallgenieauthgettokenAPIResponse, error) {
	var resp alilabs.AlibabaailabstmallgenieauthgettokenAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
