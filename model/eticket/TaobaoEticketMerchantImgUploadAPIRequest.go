package eticket

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoEticketMerchantImgUploadAPIRequest
码商上传二维码图片 API请求
taobao.eticket.merchant.img.upload

电子凭证的码商可以通过这个接口，上传二维码图片 */
type TaobaoEticketMerchantImgUploadAPIRequest struct {
	model.Params
	// 二维码图片
	_imgBytes *model.File
}

// New
