package alilabs

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alilabs"
)

// Alibabaailabstmallgenieauthdeviceunbind 解绑设备
// alibaba.ailabs.tmallgenie.auth.device.unbind
//
// 通过此接口解绑天猫精灵设备
func Alibabaailabstmallgenieauthdeviceunbind(clt *core.SDKClient, req *alilabs.AlibabaailabstmallgenieauthdeviceunbindAPIRequest, session string) (*alilabs.AlibabaailabstmallgenieauthdeviceunbindAPIResponse, error) {
	var resp alilabs.AlibabaailabstmallgenieauthdeviceunbindAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
