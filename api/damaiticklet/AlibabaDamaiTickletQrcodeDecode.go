package damaiticklet

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/damaiticklet"
)

// AlibabaDamaiTickletQrcodeDecode 票夹-动态二维码-解码
// alibaba.damai.ticklet.qrcode.decode
//
// 对于票夹的动态二维码进行解码
func AlibabaDamaiTickletQrcodeDecode(clt *core.SDKClient, req *damaiticklet.AlibabaDamaiTickletQrcodeDecodeAPIRequest, resp *damaiticklet.AlibabaDamaiTickletQrcodeDecodeAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
