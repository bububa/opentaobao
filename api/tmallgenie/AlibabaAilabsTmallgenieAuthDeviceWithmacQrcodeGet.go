package tmallgenie

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallgenie"
)

// Alibabaailabstmallgenieauthdevicewithmacqrcodeget 根据mac查询设备的安全二维码
// alibaba.ailabs.tmallgenie.auth.device.withmac.qrcode.get
//
// 根据mac查询二维码详细信息
func Alibabaailabstmallgenieauthdevicewithmacqrcodeget(clt *core.SDKClient, req *tmallgenie.AlibabaailabstmallgenieauthdevicewithmacqrcodegetAPIRequest, session string) (*tmallgenie.AlibabaailabstmallgenieauthdevicewithmacqrcodegetAPIResponse, error) {
	var resp tmallgenie.AlibabaailabstmallgenieauthdevicewithmacqrcodegetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
