package alilabs

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alilabs"
)

// Alibabaailabstmallgenieauthdeviceget 获取设备详情
// alibaba.ailabs.tmallgenie.auth.device.get
//
// 通过此接口获取设备详情
func Alibabaailabstmallgenieauthdeviceget(clt *core.SDKClient, req *alilabs.AlibabaailabstmallgenieauthdevicegetAPIRequest, session string) (*alilabs.AlibabaailabstmallgenieauthdevicegetAPIResponse, error) {
	var resp alilabs.AlibabaailabstmallgenieauthdevicegetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
