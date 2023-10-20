package eticket

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoEticketMerchantImgUploadAPIRequest 码商上传二维码图片 API请求
// taobao.eticket.merchant.img.upload
//
// 电子凭证的码商可以通过这个接口，上传二维码图片
type TaobaoEticketMerchantImgUploadAPIRequest struct {
	model.Params
	// 二维码图片
	_imgBytes *model.File
}

// NewTaobaoEticketMerchantImgUploadRequest 初始化TaobaoEticketMerchantImgUploadAPIRequest对象
func NewTaobaoEticketMerchantImgUploadRequest() *TaobaoEticketMerchantImgUploadAPIRequest {
	return &TaobaoEticketMerchantImgUploadAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoEticketMerchantImgUploadAPIRequest) GetApiMethodName() string {
	return "taobao.eticket.merchant.img.upload"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoEticketMerchantImgUploadAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoEticketMerchantImgUploadAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetImgBytes is ImgBytes Setter
// 二维码图片
func (r *TaobaoEticketMerchantImgUploadAPIRequest) SetImgBytes(_imgBytes *model.File) error {
	r._imgBytes = _imgBytes
	r.Set("img_bytes", _imgBytes)
	return nil
}

// GetImgBytes ImgBytes Getter
func (r TaobaoEticketMerchantImgUploadAPIRequest) GetImgBytes() *model.File {
	return r._imgBytes
}
