package alilabs

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alilabs"
)

// Alibabaailabstmallgenieauthtaobaoauth 天猫精灵淘宝登录授权绑定接口
// alibaba.ailabs.tmallgenie.auth.taobaoauth
//
// 厂商获取用户淘宝授权之后，通过此接口获取天猫精灵授权，并绑定一台设备
func Alibabaailabstmallgenieauthtaobaoauth(clt *core.SDKClient, req *alilabs.AlibabaailabstmallgenieauthtaobaoauthAPIRequest, session string) (*alilabs.AlibabaailabstmallgenieauthtaobaoauthAPIResponse, error) {
	var resp alilabs.AlibabaailabstmallgenieauthtaobaoauthAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
