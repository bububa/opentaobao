package damaiticklet

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaDamaiTickletQrcodeDecodeAPIRequest
票夹-动态二维码-解码 API请求
alibaba.damai.ticklet.qrcode.decode

对于票夹的动态二维码进行解码 */
type AlibabaDamaiTickletQrcodeDecodeAPIRequest struct {
	model.Params
	// 生产系统
	_productSystemId string
	// 加密二维码
	_encryptedQrCode string
}

// New
