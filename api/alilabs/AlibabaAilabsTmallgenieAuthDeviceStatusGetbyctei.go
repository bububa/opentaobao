package alilabs

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alilabs"
)

// Alibabaailabstmallgenieauthdevicestatusgetbyctei 根据ctei查询状态
// alibaba.ailabs.tmallgenie.auth.device.status.getbyctei
//
// 提供给电信查询设备在线状态值
func Alibabaailabstmallgenieauthdevicestatusgetbyctei(clt *core.SDKClient, req *alilabs.AlibabaailabstmallgenieauthdevicestatusgetbycteiAPIRequest, session string) (*alilabs.AlibabaailabstmallgenieauthdevicestatusgetbycteiAPIResponse, error) {
	var resp alilabs.AlibabaailabstmallgenieauthdevicestatusgetbycteiAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
