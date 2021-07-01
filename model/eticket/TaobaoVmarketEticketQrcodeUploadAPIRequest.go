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

// NewTaobaoVmarketEticketQrcodeUploadRequest 初始化TaobaoVmarketEticketQrcodeUploadAPIRequest对象
func NewTaobaoVmarketEticketQrcodeUploadRequest() *TaobaoVmarketEticketQrcodeUploadAPIRequest {
	return &TaobaoVmarketEticketQrcodeUploadAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoVmarketEticketQrcodeUploadAPIRequest) GetApiMethodName() string {
	return "taobao.vmarket.eticket.qrcode.upload"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoVmarketEticketQrcodeUploadAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is CodeMerchantId Setter
// 码商ID
func (r *TaobaoVmarketEticketQrcodeUploadAPIRequest) SetCodeMerchantId(_codeMerchantId int64) error {
	r._codeMerchantId = _codeMerchantId
	r.Set("code_merchant_id", _codeMerchantId)
	return nil
}

// Get CodeMerchantId Getter
func (r TaobaoVmarketEticketQrcodeUploadAPIRequest) GetCodeMerchantId() int64 {
	return r._codeMerchantId
}

// Set is ImgBytes Setter
// 上传的图片byte[]  小于300K，图片尺寸400*400以内
func (r *TaobaoVmarketEticketQrcodeUploadAPIRequest) SetImgBytes(_imgBytes *model.File) error {
	r._imgBytes = _imgBytes
	r.Set("img_bytes", _imgBytes)
	return nil
}

// Get ImgBytes Getter
func (r TaobaoVmarketEticketQrcodeUploadAPIRequest) GetImgBytes() *model.File {
	return r._imgBytes
}
