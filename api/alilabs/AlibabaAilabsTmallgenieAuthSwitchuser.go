package alilabs

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alilabs"
)

// Alibabaailabstmallgenieauthswitchuser 切换用户
// alibaba.ailabs.tmallgenie.auth.switchuser
//
// 设备切换授权用户
func Alibabaailabstmallgenieauthswitchuser(clt *core.SDKClient, req *alilabs.AlibabaailabstmallgenieauthswitchuserAPIRequest, session string) (*alilabs.AlibabaailabstmallgenieauthswitchuserAPIResponse, error) {
	var resp alilabs.AlibabaailabstmallgenieauthswitchuserAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
