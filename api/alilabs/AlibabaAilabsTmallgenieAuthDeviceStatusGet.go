package alilabs

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alilabs"
)

// Alibabaailabstmallgenieauthdevicestatusget 设备状态查询
// alibaba.ailabs.tmallgenie.auth.device.status.get
//
// 提供给天猫精灵定制机厂商 查询设备在线状态值
func Alibabaailabstmallgenieauthdevicestatusget(clt *core.SDKClient, req *alilabs.AlibabaailabstmallgenieauthdevicestatusgetAPIRequest, session string) (*alilabs.AlibabaailabstmallgenieauthdevicestatusgetAPIResponse, error) {
	var resp alilabs.AlibabaailabstmallgenieauthdevicestatusgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
