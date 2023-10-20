package alilabs

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alilabs"
)

// Alibabaailabstmallgenieauthdevicewithdeviceidget 根据三方ID查询设备注册激活信息
// alibaba.ailabs.tmallgenie.auth.device.withdeviceid.get
//
// 根据三方ID查询设备注册激活信息
func Alibabaailabstmallgenieauthdevicewithdeviceidget(clt *core.SDKClient, req *alilabs.AlibabaailabstmallgenieauthdevicewithdeviceidgetAPIRequest, session string) (*alilabs.AlibabaailabstmallgenieauthdevicewithdeviceidgetAPIResponse, error) {
	var resp alilabs.AlibabaailabstmallgenieauthdevicewithdeviceidgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
