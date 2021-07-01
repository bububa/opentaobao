package eticket

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoVmarketEticketQrcodeUploadAPIRequest
码商二维码图片上传 API请求
taobao.vmarket.eticket.qrcode.upload

电子凭证的码商可以通过这个接口，上传他们发送的二维码图片 */
type TaobaoVmarketEticketQrcodeUploadAPIRequest struct {
	model.Params
	// 码商ID
	_codeMerchantId int64
	// 上传的图片byte[]  小于300K，图片尺寸400*400以内
	_imgBytes *model.File
}

// New
