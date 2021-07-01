package alilabs

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alilabs"
)

/* AlibabaAilabsTmallgenieAuthTaobaoauth
天猫精灵淘宝登录授权绑定接口
alibaba.ailabs.tmallgenie.auth.taobaoauth

厂商获取用户淘宝授权之后，通过此接口获取天猫精灵授权，并绑定一台设备 */
func AlibabaAilabsTmallgenieAuthTaobaoauth(clt *core.SDKClient, req *alilabs.AlibabaAilabsTmallgenieAuthTaobaoauthAPIRequest, session string) (*alilabs.AlibabaAilabsTmallgenieAuthTaobaoauthAPIResponse, error) {
	var resp alilabs.AlibabaAilabsTmallgenieAuthTaobaoauthAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
