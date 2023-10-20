package tmallgenie

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallgenie"
)

// Alibabaailabstmallgenieauthdeviceqrcodestaticbind 静态二维码绑定
// alibaba.ailabs.tmallgenie.auth.device.qrcode.staticbind
//
// 静态二维码绑定
func Alibabaailabstmallgenieauthdeviceqrcodestaticbind(clt *core.SDKClient, req *tmallgenie.AlibabaailabstmallgenieauthdeviceqrcodestaticbindAPIRequest, session string) (*tmallgenie.AlibabaailabstmallgenieauthdeviceqrcodestaticbindAPIResponse, error) {
	var resp tmallgenie.AlibabaailabstmallgenieauthdeviceqrcodestaticbindAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
