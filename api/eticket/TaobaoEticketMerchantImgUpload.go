package eticket

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/eticket"
)

// TaobaoEticketMerchantImgUpload 码商上传二维码图片
// taobao.eticket.merchant.img.upload
//
// 电子凭证的码商可以通过这个接口，上传二维码图片
func TaobaoEticketMerchantImgUpload(clt *core.SDKClient, req *eticket.TaobaoEticketMerchantImgUploadAPIRequest, resp *eticket.TaobaoEticketMerchantImgUploadAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
