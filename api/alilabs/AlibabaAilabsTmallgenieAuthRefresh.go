package alilabs

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alilabs"
)

// Alibabaailabstmallgenieauthrefresh 刷新token
// alibaba.ailabs.tmallgenie.auth.refresh
//
// 通过此接口刷新天猫精灵授权token
func Alibabaailabstmallgenieauthrefresh(clt *core.SDKClient, req *alilabs.AlibabaailabstmallgenieauthrefreshAPIRequest, session string) (*alilabs.AlibabaailabstmallgenieauthrefreshAPIResponse, error) {
	var resp alilabs.AlibabaailabstmallgenieauthrefreshAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
