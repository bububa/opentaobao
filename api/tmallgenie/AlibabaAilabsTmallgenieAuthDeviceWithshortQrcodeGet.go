package tmallgenie

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallgenie"
)

// Alibabaailabstmallgenieauthdevicewithshortqrcodeget 根据安全简码查询二维码详细信息
// alibaba.ailabs.tmallgenie.auth.device.withshort.qrcode.get
//
// 根据安全简码查询二维码详细信息
func Alibabaailabstmallgenieauthdevicewithshortqrcodeget(clt *core.SDKClient, req *tmallgenie.AlibabaailabstmallgenieauthdevicewithshortqrcodegetAPIRequest, session string) (*tmallgenie.AlibabaailabstmallgenieauthdevicewithshortqrcodegetAPIResponse, error) {
	var resp tmallgenie.AlibabaailabstmallgenieauthdevicewithshortqrcodegetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
