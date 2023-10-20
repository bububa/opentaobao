package alilabs

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alilabs"
)

// Alibabaailabstmallgenieauthdevicelist 获取用户设备列表
// alibaba.ailabs.tmallgenie.auth.device.list
//
// 通过此接口获取用户绑定的设备信息列表
func Alibabaailabstmallgenieauthdevicelist(clt *core.SDKClient, req *alilabs.AlibabaailabstmallgenieauthdevicelistAPIRequest, session string) (*alilabs.AlibabaailabstmallgenieauthdevicelistAPIResponse, error) {
	var resp alilabs.AlibabaailabstmallgenieauthdevicelistAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
